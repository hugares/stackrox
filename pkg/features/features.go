// Package features helps enable or disable features.
package features

import (
	"fmt"
	"strings"

	"github.com/stackrox/rox/pkg/buildinfo"
)

// A FeatureFlag is a product behavior that can be enabled or disabled using an environment variable.
type FeatureFlag interface {
	Name() string
	EnvVar() string
	Enabled() bool
	Default() bool
	Stage() string
}

var (
	// Flags contains all defined FeatureFlags by name.
	Flags = make(map[string]FeatureFlag)
)

// registerFeature global registers and returns a new feature flag that can be overridden from the default state regardless of build.
func registerFeature(name, envVar string, defaultValue bool) FeatureFlag {
	return saveFeature(name, envVar, defaultValue, true, devPreview)
}

// registerTechPreviewFeature is like registerFeature, but registers a tech-preview feature.
func registerTechPreviewFeature(name, envVar string, defaultValue bool) FeatureFlag {
	return saveFeature(name, envVar, defaultValue, true, techPreview)
}

// registerUnchangeableFeature global registers and returns a new feature flag that is locked to the default value in release builds.
func registerUnchangeableFeature(name, envVar string, defaultValue bool) FeatureFlag {
	return saveFeature(name, envVar, defaultValue, !buildinfo.ReleaseBuild, devPreview)
}

// registerUnchangeableTechPreviewFeature is like registerUnchangeableFeature, but registers a tech-preview feature.
func registerUnchangeableTechPreviewFeature(name, envVar string, defaultValue bool) FeatureFlag {
	return saveFeature(name, envVar, defaultValue, !buildinfo.ReleaseBuild, techPreview)
}

func saveFeature(name, envVar string, defaultValue, overridable bool, release stage) FeatureFlag {
	if !strings.HasPrefix(envVar, "ROX_") {
		panic(fmt.Sprintf("invalid env var: %s, must start with ROX_", envVar))
	}
	f := &feature{
		name:         name,
		envVar:       envVar,
		defaultValue: defaultValue,
		overridable:  overridable,
		techPreview:  release,
	}
	Flags[f.envVar] = f
	return f
}
