package mappers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	nvdschema "github.com/facebookincubator/nvdtools/cveapi/nvd/schema"
	"github.com/facebookincubator/nvdtools/cvss2"
	"github.com/facebookincubator/nvdtools/cvss3"
	"github.com/gogo/protobuf/types"
	"github.com/quay/claircore"
	"github.com/quay/claircore/pkg/cpe"
	"github.com/quay/zlog"
	v4 "github.com/stackrox/rox/generated/internalapi/scanner/v4"
	"github.com/stackrox/rox/scanner/enricher/fixedby"
	"github.com/stackrox/rox/scanner/enricher/nvd"
)

var (
	severityMapping = map[claircore.Severity]v4.VulnerabilityReport_Vulnerability_Severity{
		claircore.Unknown:    v4.VulnerabilityReport_Vulnerability_SEVERITY_UNSPECIFIED,
		claircore.Negligible: v4.VulnerabilityReport_Vulnerability_SEVERITY_LOW,
		claircore.Low:        v4.VulnerabilityReport_Vulnerability_SEVERITY_LOW,
		claircore.Medium:     v4.VulnerabilityReport_Vulnerability_SEVERITY_MODERATE,
		claircore.High:       v4.VulnerabilityReport_Vulnerability_SEVERITY_IMPORTANT,
		claircore.Critical:   v4.VulnerabilityReport_Vulnerability_SEVERITY_CRITICAL,
	}
	// vulnNamePatterns is a prioritized list of regexes to match against
	// vulnerability information to extract their ID.
	vulnNamePatterns = []*regexp.Regexp{
		regexp.MustCompile(`((RHSA|RHBA|RHEA)-\d{4}:\d+)|(ALAS\d*-\d{4}-\d+)`),
		regexp.MustCompile(`CVE-\d{4}-\d+`),
		// GHSA, see: https://github.com/github/advisory-database#ghsa-ids
		regexp.MustCompile(`GHSA(-[2-9cfghjmpqrvwx]{4}){3}`),
		// Catchall.
		regexp.MustCompile(`[A-Z]+-\d{4}[-:]\d+`),
	}
)

// ToProtoV4IndexReport maps claircore.IndexReport to v4.IndexReport.
func ToProtoV4IndexReport(r *claircore.IndexReport) (*v4.IndexReport, error) {
	if r == nil {
		return nil, nil
	}
	contents, err := toProtoV4Contents(r.Packages, r.Distributions, r.Repositories, r.Environments, nil)
	if err != nil {
		return nil, err
	}
	return &v4.IndexReport{
		State:    r.State,
		Success:  r.Success,
		Err:      r.Err,
		Contents: contents,
	}, nil
}

// ToProtoV4VulnerabilityReport maps claircore.VulnerabilityReport to v4.VulnerabilityReport.
func ToProtoV4VulnerabilityReport(ctx context.Context, r *claircore.VulnerabilityReport) (*v4.VulnerabilityReport, error) {
	if r == nil {
		return nil, nil
	}
	nvdVulns, err := nvdVulnerabilities(r.Enrichments)
	if err != nil {
		return nil, fmt.Errorf("internal error: parsing nvd vulns: %w", err)
	}
	vulnerabilities, err := toProtoV4VulnerabilitiesMap(ctx, r.Vulnerabilities, nvdVulns)
	if err != nil {
		return nil, fmt.Errorf("internal error: %w", err)
	}
	pkgFixedBy, err := pkgFixedBy(r.Enrichments)
	if err != nil {
		return nil, fmt.Errorf("internal error: parsing package-level fixedbys: %w", err)
	}
	contents, err := toProtoV4Contents(r.Packages, r.Distributions, r.Repositories, r.Environments, pkgFixedBy)
	if err != nil {
		return nil, err
	}
	return &v4.VulnerabilityReport{
		Vulnerabilities:        vulnerabilities,
		PackageVulnerabilities: toProtoV4PackageVulnerabilitiesMap(r.PackageVulnerabilities),
		Contents:               contents,
	}, nil
}

// ToClairCoreIndexReport converts v4.Contents to a claircore.IndexReport.
func ToClairCoreIndexReport(contents *v4.Contents) (*claircore.IndexReport, error) {
	if contents == nil {
		return nil, errors.New("internal error: empty contents")
	}
	pkgs, err := convertSliceToMap(contents.GetPackages(), toClairCorePackage)
	if err != nil {
		return nil, fmt.Errorf("internal error: %w", err)
	}
	dists, err := convertSliceToMap(contents.GetDistributions(), toClairCoreDistribution)
	if err != nil {
		return nil, fmt.Errorf("internal error: %w", err)
	}
	repos, err := convertSliceToMap(contents.GetRepositories(), toClairCoreRepository)
	if err != nil {
		return nil, fmt.Errorf("internal error: %w", err)
	}
	var environments map[string][]*claircore.Environment
	if envs := contents.GetEnvironments(); envs != nil {
		environments = make(map[string][]*claircore.Environment, len(envs))
		for k, v := range envs {
			for _, env := range v.GetEnvironments() {
				ccEnv, err := toClairCoreEnvironment(env)
				if err != nil {
					return nil, err
				}
				environments[k] = append(environments[k], ccEnv)
			}
		}
	}
	return &claircore.IndexReport{
		Packages:      pkgs,
		Distributions: dists,
		Repositories:  repos,
		Environments:  environments,
	}, nil
}

func toProtoV4Contents(
	pkgs map[string]*claircore.Package,
	dists map[string]*claircore.Distribution,
	repos map[string]*claircore.Repository,
	envs map[string][]*claircore.Environment,
	pkgFixedBy map[string]string,
) (*v4.Contents, error) {
	var environments map[string]*v4.Environment_List
	if len(envs) > 0 {
		environments = make(map[string]*v4.Environment_List, len(envs))
	}
	for k, v := range envs {
		l, ok := environments[k]
		if !ok {
			l = &v4.Environment_List{}
			environments[k] = l
		}
		for _, e := range v {
			l.Environments = append(l.Environments, toProtoV4Environment(e))
		}
	}
	var packages []*v4.Package
	if pkgFixedBy == nil {
		pkgFixedBy = map[string]string{}
	}
	for _, ccP := range pkgs {
		pkg, err := toProtoV4Package(ccP)
		if err != nil {
			return nil, err
		}
		pkg.FixedInVersion = pkgFixedBy[pkg.GetId()]
		packages = append(packages, pkg)
	}
	return &v4.Contents{
		Packages:      packages,
		Distributions: convertMapToSlice(toProtoV4Distribution, dists),
		Repositories:  convertMapToSlice(toProtoV4Repository, repos),
		Environments:  environments,
	}, nil
}

func toProtoV4Package(p *claircore.Package) (*v4.Package, error) {
	if p == nil {
		return nil, nil
	}
	if p.Source != nil && p.Source.Source != nil {
		return nil, fmt.Errorf("package %q: invalid source package %q: source specifies source",
			p.ID, p.Source.ID)
	}
	// Conversion function.
	toNormalizedVersion := func(version claircore.Version) *v4.NormalizedVersion {
		return &v4.NormalizedVersion{
			Kind: version.Kind,
			V:    version.V[:],
		}
	}
	srcPkg, err := toProtoV4Package(p.Source)
	if err != nil {
		return nil, err
	}
	return &v4.Package{
		Id:                p.ID,
		Name:              p.Name,
		Version:           p.Version,
		NormalizedVersion: toNormalizedVersion(p.NormalizedVersion),
		Kind:              p.Kind,
		Source:            srcPkg,
		PackageDb:         p.PackageDB,
		RepositoryHint:    p.RepositoryHint,
		Module:            p.Module,
		Arch:              p.Arch,
		Cpe:               toCPEString(p.CPE),
	}, nil
}

// VersionID returns the distribution version ID.
func VersionID(d *claircore.Distribution) string {
	vID := d.VersionID
	if vID == "" {
		switch d.DID {
		// TODO(ROX-21678): `VersionId` is currently not populated for Alpine[1],
		//                  temporarily falling back to the version.
		//
		// [1]: https://github.com/quay/claircore/blob/88ccfbecee88d7b326b9a2fb3ab5b5f4cfa0b610/alpine/distributionscanner.go#L110-L113
		case "alpine":
			vID = d.Version
		}
	}
	return vID
}

func toProtoV4Distribution(d *claircore.Distribution) *v4.Distribution {
	if d == nil {
		return nil
	}
	return &v4.Distribution{
		Id:              d.ID,
		Did:             d.DID,
		Name:            d.Name,
		Version:         d.Version,
		VersionCodeName: d.VersionCodeName,
		VersionId:       VersionID(d),
		Arch:            d.Arch,
		Cpe:             toCPEString(d.CPE),
		PrettyName:      d.PrettyName,
	}
}

func toProtoV4Repository(r *claircore.Repository) *v4.Repository {
	if r == nil {
		return nil
	}
	return &v4.Repository{
		Id:   r.ID,
		Name: r.Name,
		Key:  r.Key,
		Uri:  r.URI,
		Cpe:  toCPEString(r.CPE),
	}
}

func toProtoV4Environment(e *claircore.Environment) *v4.Environment {
	if e == nil {
		return nil
	}
	return &v4.Environment{
		PackageDb:      e.PackageDB,
		IntroducedIn:   toDigestString(e.IntroducedIn),
		DistributionId: e.DistributionID,
		RepositoryIds:  append([]string(nil), e.RepositoryIDs...),
	}
}

func toProtoV4PackageVulnerabilitiesMap(ccPkgVulnerabilities map[string][]string) map[string]*v4.StringList {
	if ccPkgVulnerabilities == nil {
		return nil
	}
	var pkgVulns map[string]*v4.StringList
	if len(ccPkgVulnerabilities) > 0 {
		pkgVulns = make(map[string]*v4.StringList, len(ccPkgVulnerabilities))
	}
	for k, v := range ccPkgVulnerabilities {
		if v == nil {
			continue
		}
		pkgVulns[k] = &v4.StringList{
			Values: append([]string(nil), v...),
		}
	}
	return pkgVulns
}

func toProtoV4VulnerabilitiesMap(
	ctx context.Context,
	vulns map[string]*claircore.Vulnerability,
	nvdVulns map[string]*nvdschema.CVEAPIJSON20CVEItem,
) (map[string]*v4.VulnerabilityReport_Vulnerability, error) {
	if vulns == nil {
		return nil, nil
	}
	var vulnerabilities map[string]*v4.VulnerabilityReport_Vulnerability
	for k, v := range vulns {
		if v == nil {
			continue
		}
		issued, err := types.TimestampProto(v.Issued)
		if err != nil {
			return nil, err
		}
		var pkgID string
		if v.Package != nil {
			pkgID = v.Package.ID
		}
		var distID string
		if v.Dist != nil {
			distID = v.Dist.ID
		}
		var repoID string
		if v.Repo != nil {
			repoID = v.Repo.ID
		}
		normalizedSeverity := toProtoV4VulnerabilitySeverity(ctx, v.NormalizedSeverity)
		sev, err := severityAndScores(ctx, v, nvdVulns)
		if err != nil {
			zlog.Warn(ctx).
				Err(err).
				Str("vuln_id", v.ID).
				Str("vuln_name", v.Name).
				Str("vuln_updater", v.Updater).
				Str("severity", v.Severity).
				Msg("missing severity and/or CVSS score(s): proceeding with partial values")
		}
		// Look for CVSS scores in the severity field, then set the
		// scores only if at least one is found.
		var cvss *v4.VulnerabilityReport_Vulnerability_CVSS
		hasV2, hasV3 := sev.v2Vector != "", sev.v3Vector != ""
		if hasV2 || hasV3 {
			cvss = &v4.VulnerabilityReport_Vulnerability_CVSS{}
		}
		if hasV2 {
			cvss.V2 = &v4.VulnerabilityReport_Vulnerability_CVSS_V2{
				BaseScore: sev.v2Score,
				Vector:    sev.v2Vector,
			}
		}
		if hasV3 {
			cvss.V3 = &v4.VulnerabilityReport_Vulnerability_CVSS_V3{
				BaseScore: sev.v3Score,
				Vector:    sev.v3Vector,
			}
		}
		description := v.Description
		if description == "" {
			// No description provided, so fall back to NVD.
			if v, ok := nvdVulns[v.ID]; ok {
				if len(v.Descriptions) > 0 {
					description = v.Descriptions[0].Value
				}
			}
		}
		if vulnerabilities == nil {
			vulnerabilities = make(map[string]*v4.VulnerabilityReport_Vulnerability, len(vulns))
		}
		vulnerabilities[k] = &v4.VulnerabilityReport_Vulnerability{
			Id:                 v.ID,
			Name:               vulnerabilityName(v),
			Description:        description,
			Issued:             issued,
			Link:               v.Links,
			Severity:           sev.severity,
			NormalizedSeverity: normalizedSeverity,
			PackageId:          pkgID,
			DistributionId:     distID,
			RepositoryId:       repoID,
			FixedInVersion:     fixedInVersion(v),
			Cvss:               cvss,
		}
	}
	return vulnerabilities, nil
}

func toProtoV4VulnerabilitySeverity(ctx context.Context, ccSeverity claircore.Severity) v4.VulnerabilityReport_Vulnerability_Severity {
	if mappedSeverity, ok := severityMapping[ccSeverity]; ok {
		return mappedSeverity
	}
	zlog.Warn(ctx).
		Str("claircore_severity", ccSeverity.String()).
		Msgf("unknown ClairCore severity, mapping to %s", v4.VulnerabilityReport_Vulnerability_SEVERITY_UNSPECIFIED.String())
	return v4.VulnerabilityReport_Vulnerability_SEVERITY_UNSPECIFIED
}

func toCPEString(c cpe.WFN) string {
	return c.BindFS()
}

func toDigestString(digest claircore.Digest) string {
	return digest.String()
}

func toClairCoreCPE(s string) (cpe.WFN, error) {
	c, err := cpe.UnbindFS(s)
	if err != nil {
		return c, fmt.Errorf("%q: %s", s, strings.TrimPrefix(err.Error(), "cpe: "))
	}
	return c, nil
}

func toClairCorePackage(p *v4.Package) (string, *claircore.Package, error) {
	if p == nil {
		return "", nil, nil
	}
	// Conversion function.
	toNormalizedVersion := func(v *v4.NormalizedVersion) (ccV claircore.Version) {
		ccV.Kind = v.GetKind()
		copy(ccV.V[:], v.GetV())
		return
	}
	// Fields that might fail.
	ccCPE, err := toClairCoreCPE(p.GetCpe())
	if err != nil {
		return "", nil, fmt.Errorf("package %q: %w", p.GetId(), err)
	}
	if p.GetSource().GetSource() != nil {
		return "", nil, fmt.Errorf("package %q: invalid source package %q: source specifies source",
			p.GetId(), p.GetSource().GetId())
	}
	_, src, err := toClairCorePackage(p.GetSource())
	if err != nil {
		return "", nil, err
	}
	return p.GetId(), &claircore.Package{
		ID:                p.GetId(),
		Name:              p.GetName(),
		Version:           p.GetVersion(),
		Kind:              p.GetKind(),
		Source:            src,
		PackageDB:         p.GetPackageDb(),
		RepositoryHint:    p.GetRepositoryHint(),
		NormalizedVersion: toNormalizedVersion(p.GetNormalizedVersion()),
		Module:            p.GetModule(),
		Arch:              p.GetArch(),
		CPE:               ccCPE,
	}, nil
}

func toClairCoreDistribution(d *v4.Distribution) (string, *claircore.Distribution, error) {
	if d == nil {
		return "", nil, nil
	}
	ccCPE, err := toClairCoreCPE(d.GetCpe())
	if err != nil {
		return "", nil, fmt.Errorf("distribution %q: %w", d.GetId(), err)
	}
	return d.GetId(), &claircore.Distribution{
		ID:              d.GetId(),
		DID:             d.GetDid(),
		Name:            d.GetName(),
		Version:         d.GetVersion(),
		VersionCodeName: d.GetVersionCodeName(),
		VersionID:       d.GetVersionId(),
		Arch:            d.GetArch(),
		CPE:             ccCPE,
		PrettyName:      d.GetPrettyName(),
	}, nil
}

func toClairCoreRepository(r *v4.Repository) (string, *claircore.Repository, error) {
	if r == nil {
		return "", nil, nil
	}
	ccCPE, err := toClairCoreCPE(r.GetCpe())
	if err != nil {
		return "", nil, fmt.Errorf("repository %q: %w", r.GetId(), err)
	}
	return r.GetId(), &claircore.Repository{
		ID:   r.Id,
		Name: r.Name,
		Key:  r.Key,
		URI:  r.Uri,
		CPE:  ccCPE,
	}, nil
}

func toClairCoreEnvironment(env *v4.Environment) (*claircore.Environment, error) {
	introducedIn, err := claircore.ParseDigest(env.GetIntroducedIn())
	if err != nil {
		return nil, err
	}
	return &claircore.Environment{
		PackageDB:      env.GetPackageDb(),
		IntroducedIn:   introducedIn,
		DistributionID: env.GetDistributionId(),
		RepositoryIDs:  env.GetRepositoryIds(),
	}, nil
}

// convertSliceToMap converts a slice of pointers of a generic type to a map
// based on the returned value of a conversion function that returns a string
// key, the pointer to the converted value, or error if the conversion failed.
// Nils in the slice are ignored.
func convertSliceToMap[IN any, OUT any](in []*IN, convF func(*IN) (string, *OUT, error)) (map[string]*OUT, error) {
	if len(in) == 0 {
		return nil, nil
	}
	m := make(map[string]*OUT, len(in))
	for _, v := range in {
		if v == nil {
			continue
		}
		k, ccV, err := convF(v)
		if err != nil {
			return nil, err
		}
		if ccV == nil {
			continue
		}
		m[k] = ccV
	}
	return m, nil
}

// convertMapToSlice converts generic maps keyed by strings to a slice using a
// provided conversion function.
func convertMapToSlice[IN any, OUT any](convF func(*IN) *OUT, in map[string]*IN) (out []*OUT) {
	for _, i := range in {
		out = append(out, convF(i))
	}
	return out
}

// fixedInVersion returns the fixed in string, typically provided the report's
// `FixedInVersion` as a plain string, but, in some OSV updaters, it can be an
// urlencoded string.
func fixedInVersion(v *claircore.Vulnerability) string {
	fixedIn := v.FixedInVersion
	// Try to parse url encoded params; if expected values are not found leave it.
	if q, err := url.ParseQuery(fixedIn); err == nil && q.Has("fixed") {
		fixedIn = q.Get("fixed")
	}
	return fixedIn
}

// nvdVulnerabilities look for NVD CVSS in the vulnerability report enrichments and
// returns a map of CVEs.
func nvdVulnerabilities(enrichments map[string][]json.RawMessage) (map[string]*nvdschema.CVEAPIJSON20CVEItem, error) {
	enrichmentsList := enrichments[nvd.Type]
	if len(enrichmentsList) == 0 {
		return nil, nil
	}
	var items map[string][]nvdschema.CVEAPIJSON20CVEItem
	// The CVSS enrichment always contains only one element.
	err := json.Unmarshal(enrichmentsList[0], &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	ret := make(map[string]*nvdschema.CVEAPIJSON20CVEItem)
	for id, list := range items {
		// There is no criteria for selecting more than one enrichment record, assume the
		// first is the right one.
		i := list[0]
		ret[id] = &i
	}
	return ret, nil
}

func pkgFixedBy(enrichments map[string][]json.RawMessage) (map[string]string, error) {
	enrichmentsList := enrichments[fixedby.Type]
	if len(enrichmentsList) == 0 {
		return nil, nil
	}
	var pkgFixedBys map[string]string
	// The fixedby enrichment always contains only one element.
	err := json.Unmarshal(enrichmentsList[0], &pkgFixedBys)
	if err != nil {
		return nil, err
	}
	if len(pkgFixedBys) == 0 {
		return nil, nil
	}
	return pkgFixedBys, nil
}

// severityValues contains severity information that can retrieved from a
// ClairCore vulnerability report.
type severityValues struct {
	severity string
	v2Vector string
	v2Score  float32
	v3Vector string
	v3Score  float32
}

var (
	osvUpdaterPattern  = regexp.MustCompile(`^osv/.*`)
	rhelUpdaterPattern = regexp.MustCompile(`^RHEL\d+-`)
)

// severityAndScores returns the severity and scores information out of a
// ClairCore vulnerability. The returned information is dependent on the
// underlying updater.
func severityAndScores(ctx context.Context, vuln *claircore.Vulnerability, nvdVulns map[string]*nvdschema.CVEAPIJSON20CVEItem) (severityValues, error) {
	ctx = zlog.ContextWithValues(ctx,
		"component", "mappers/severityAndScores",
		"updater", vuln.Updater,
		"vuln_id", vuln.ID,
		"vuln_name", vuln.Name,
	)

	switch {
	case rhelUpdaterPattern.MatchString(vuln.Updater):
		return rhelSeverityAndScores(vuln)
	case osvUpdaterPattern.MatchString(vuln.Updater):
		sev, err := osvVectors(vuln)
		if err != nil {
			zlog.Debug(ctx).
				Err(err).
				Str("severity", vuln.Severity).
				Msg("parsing severity, falling back to NVD")
			break
		}
		return sev, nil
	}

	// Default/fallback is NVD.
	return nvdSeverityAndScores(vuln, nvdVulns)
}

func rhelSeverityAndScores(vuln *claircore.Vulnerability) (severityValues, error) {
	if vuln.Severity == "" {
		return severityValues{}, errors.New("severity is empty")
	}

	q, err := url.ParseQuery(vuln.Severity)
	if err != nil {
		return severityValues{}, fmt.Errorf("parsing severity: %w", err)
	}

	values := severityValues{
		severity: q.Get("severity"),
	}

	if v := q.Get("cvss2_vector"); v != "" {
		if _, err := cvss2.VectorFromString(v); err != nil {
			return values, fmt.Errorf("parsing CVSS v2 vector: %w", err)
		}

		s := q.Get("cvss2_score")
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return values, fmt.Errorf("parsing CVSS v2 score: %w", err)
		}

		values.v2Vector = v
		values.v2Score = float32(f)
	}

	if v := q.Get("cvss3_vector"); v != "" {
		if _, err := cvss3.VectorFromString(v); err != nil {
			return values, fmt.Errorf("parsing CVSS v3 vector: %w", err)
		}

		s := q.Get("cvss3_score")
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return values, fmt.Errorf("parsing CVSS v3 score: %w", err)
		}
		values.v3Vector = v
		values.v3Score = float32(f)
	}

	return values, nil
}

// osvVectors parses the CVSS vectors from the give OSV vulnerability.
//
// Note: we do not populate severity nor the scores here. It is up to the client to determine those.
// See pkg/scanners/scannerv4/convert.go.
func osvVectors(vuln *claircore.Vulnerability) (severityValues, error) {
	if vuln.Severity == "" {
		return severityValues{}, errors.New("severity is empty")
	}

	values := severityValues{
		severity: vuln.Severity,
	}

	// ClairCore has no CVSS version indicator for OSV data, so guess via the prefix.
	if strings.HasPrefix(vuln.Severity, "CVSS:3") {
		if _, err := cvss3.VectorFromString(vuln.Severity); err != nil {
			return values, fmt.Errorf("parsing CVSS v3 vector: %w", err)
		}
		values.v3Vector = vuln.Severity
		return values, nil
	}

	// We do not support CVSS v4 yet, so if it's not 3, then it's 2.
	if _, err := cvss2.VectorFromString(vuln.Severity); err != nil {
		return values, fmt.Errorf("parsing CVSS v2 vector: %w", err)
	}
	values.v2Vector = vuln.Severity
	return values, nil
}

func nvdSeverityAndScores(vuln *claircore.Vulnerability, nvdVulns map[string]*nvdschema.CVEAPIJSON20CVEItem) (severityValues, error) {
	values := severityValues{
		severity: vuln.Severity,
	}

	v, ok := nvdVulns[vuln.ID]
	if !ok {
		return values, errors.New("cannot find NVD data")
	}

	// Sanity check the NVD data.
	if v.Metrics == nil || (v.Metrics.CvssMetricV31 == nil && v.Metrics.CvssMetricV30 == nil && v.Metrics.CvssMetricV2 == nil) {
		return values, errors.New("no metrics for vuln")
	}

	if len(v.Metrics.CvssMetricV30) > 0 {
		if cvssv30 := v.Metrics.CvssMetricV30[0]; cvssv30 != nil && cvssv30.CvssData != nil {
			values.v3Score = float32(cvssv30.CvssData.BaseScore)
			values.v3Vector = cvssv30.CvssData.VectorString
		}
	}
	// If there is both CVSS 3.0 and 3.1 data, use 3.1.
	if len(v.Metrics.CvssMetricV31) > 0 {
		if cvssv31 := v.Metrics.CvssMetricV31[0]; cvssv31 != nil && cvssv31.CvssData != nil {
			values.v3Score = float32(cvssv31.CvssData.BaseScore)
			values.v3Vector = cvssv31.CvssData.VectorString
		}
	}
	if len(v.Metrics.CvssMetricV2) > 0 {
		if cvssv2 := v.Metrics.CvssMetricV2[0]; cvssv2 != nil && cvssv2.CvssData != nil {
			values.v2Score = float32(cvssv2.CvssData.BaseScore)
			values.v2Vector = cvssv2.CvssData.VectorString
		}
	}

	return values, nil
}

// vulnerabilityName searches the best known candidate for the vulnerability name
// in the vulnerability details. It works by matching data against well-known
// name patterns, and defaults to the original name if nothing is found.
func vulnerabilityName(vuln *claircore.Vulnerability) string {
	for _, p := range vulnNamePatterns {
		v := p.FindString(vuln.Name)
		if v != "" {
			return v
		}
		v = p.FindString(vuln.Links)
		if v != "" {
			return v
		}
	}
	return vuln.Name
}
