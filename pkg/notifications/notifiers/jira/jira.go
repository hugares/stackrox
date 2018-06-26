package jira

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/errorhelpers"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/notifications/notifiers"
	"bitbucket.org/stack-rox/apollo/pkg/urlfmt"
	jiraLib "github.com/andygrunwald/go-jira"
)

const (
	timeout = 5 * time.Second
)

var (
	log = logging.LoggerForModule()
)

// Jira notifier plugin
type jira struct {
	client *jiraLib.Client

	conf *v1.Jira

	*v1.Notifier
}

func (j *jira) getAlertDescription(alert *v1.Alert) (string, error) {
	funcMap := template.FuncMap{
		"header": func(s string) string {
			return fmt.Sprintf("\r\n h4. %v\r\n", s)
		},
		"subheader": func(s string) string {
			return fmt.Sprintf("\r\n h5. %v\r\n", s)
		},
		"line": func(s string) string {
			return fmt.Sprintf("%v\r\n", s)
		},
		"list": func(s string) string {
			return fmt.Sprintf("* %v\r\n", s)
		},
		"nestedList": func(s string) string {
			return fmt.Sprintf("** %v\r\n", s)
		},
	}
	alertLink := notifiers.AlertLink(j.Notifier.UiEndpoint, alert.GetId())
	return notifiers.FormatPolicy(alert, alertLink, funcMap)
}

func (j *jira) getBenchmarkDescription(schedule *v1.BenchmarkSchedule) (string, error) {
	benchmarkLink := notifiers.BenchmarkLink(j.Notifier.UiEndpoint)
	return notifiers.FormatBenchmark(schedule, benchmarkLink)
}

// AlertNotify takes in an alert and generates the notification
func (j *jira) AlertNotify(alert *v1.Alert) error {
	description, err := j.getAlertDescription(alert)
	if err != nil {
		return err
	}

	project := notifiers.GetLabelValue(alert, j.GetLabelKey(), j.GetLabelDefault())
	i := &jiraLib.Issue{
		Fields: &jiraLib.IssueFields{
			Summary: fmt.Sprintf("Deployment %v (%v) violates '%v' Policy", alert.Deployment.Name, alert.Deployment.Id, alert.Policy.Name),
			Type: jiraLib.IssueType{
				Name: j.conf.GetIssueType(),
			},
			Project: jiraLib.Project{
				Key: project,
			},
			Description: description,
			Priority: &jiraLib.Priority{
				Name: severityToPriority(alert.GetPolicy().GetSeverity()),
			},
		},
	}
	return j.createIssue(i)
}

// BenchmarkNotify takes in a benchmark and generates the notification
func (j *jira) BenchmarkNotify(schedule *v1.BenchmarkSchedule) error {
	description, err := j.getBenchmarkDescription(schedule)
	if err != nil {
		return err
	}

	i := &jiraLib.Issue{
		Fields: &jiraLib.IssueFields{
			Summary: fmt.Sprintf("New Benchmark Results for %v", schedule.GetBenchmarkName()),
			Type: jiraLib.IssueType{
				Name: j.conf.GetIssueType(),
			},
			Project: jiraLib.Project{
				Key: j.GetLabelDefault(),
			},
			Description: description,
			Priority: &jiraLib.Priority{
				Name: "P3-Low",
			},
		},
	}
	return j.createIssue(i)
}

func validate(jira *v1.Jira) error {
	var errors []string
	if jira.GetIssueType() == "" {
		errors = append(errors, "Issue Type must be specified")
	}
	if jira.GetUrl() == "" {
		errors = append(errors, "URL must be specified")
	}
	if jira.GetUsername() == "" {
		errors = append(errors, "Username must be specified")
	}
	if jira.GetPassword() == "" {
		errors = append(errors, "Password must be specified")
	}
	return errorhelpers.FormatErrorStrings("Jira validation", errors)
}

func newJira(notifier *v1.Notifier) (*jira, error) {
	jiraConfig, ok := notifier.GetConfig().(*v1.Notifier_Jira)
	if !ok {
		return nil, fmt.Errorf("Jira configuration required")
	}
	conf := jiraConfig.Jira
	if err := validate(conf); err != nil {
		return nil, err
	}

	url, err := urlfmt.FormatURL(conf.GetUrl(), true, true)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Timeout: timeout,
	}
	client, err := jiraLib.NewClient(httpClient, url)
	if err != nil {
		return nil, err
	}
	res, err := client.Authentication.AcquireSessionCookie(conf.GetUsername(), conf.GetPassword())
	if err != nil {
		return nil, err
	}
	if !res {
		return nil, errors.New("Result of authentication is false")
	}
	// forces the auth to use basic auth per request
	client.Authentication.SetBasicAuth(conf.GetUsername(), conf.GetPassword())

	return &jira{
		client:   client,
		Notifier: notifier,
	}, nil
}

func (j *jira) ProtoNotifier() *v1.Notifier {
	return j.Notifier
}

func (j *jira) createIssue(i *jiraLib.Issue) error {
	_, resp, err := j.client.Issue.Create(i)
	if err != nil {
		bytes, readErr := ioutil.ReadAll(resp.Body)
		if readErr == nil {
			return fmt.Errorf("Error creating issue: %+v. Response: %v", err, string(bytes))
		}
	}
	return err
}

func (j *jira) Test() error {
	i := &jiraLib.Issue{
		Fields: &jiraLib.IssueFields{
			Description: "StackRox Test Issue",
			Type: jiraLib.IssueType{
				Name: j.conf.GetIssueType(),
			},
			Project: jiraLib.Project{
				Key: j.GetLabelDefault(),
			},
			Summary: "This is a test issue created to test integration with StackRox.",
			Priority: &jiraLib.Priority{
				Name: severityToPriority(v1.Severity_LOW_SEVERITY),
			},
		},
	}
	return j.createIssue(i)
}

func severityToPriority(sev v1.Severity) string {
	switch sev {
	case v1.Severity_CRITICAL_SEVERITY:
		return "P0-Highest"
	case v1.Severity_HIGH_SEVERITY:
		return "P1-High"
	case v1.Severity_MEDIUM_SEVERITY:
		return "P2-Medium"
	case v1.Severity_LOW_SEVERITY:
		return "P3-Low"
	default:
		return "P4-Lowest"
	}
}

func init() {
	notifiers.Add("jira", func(notifier *v1.Notifier) (notifiers.Notifier, error) {
		j, err := newJira(notifier)
		return j, err
	})
}
