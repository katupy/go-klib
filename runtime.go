package klib

type DeployEnv int

const (
	// The default, non-ephemeral, deployment environment.
	DeployEnvProduction DeployEnv = 0

	// A remote environment that closely matches production.
	DeployEnvStaging DeployEnv = 10

	// A remote environment to test newly developed features.
	// Can be used to stress-test the infrastructure.
	DeployEnvTesting DeployEnv = 20

	// A remote environment that is used for automated tests,
	// usually when changes are pushed to code repositories.
	DeployEnvCI DeployEnv = 30

	// A local environment that is used for automated tests,
	// usually to ensure tests pass before pushing changes to code repositories.
	DeployEnvLocalCI DeployEnv = 40

	// A local environment that is used for local development or testing.
	DeployEnvLocal DeployEnv = 50
)

var buildInfo string

// BuildInfo returns the static string value that holds information set during build,
// like app version, build version, etc.
func BuildInfo() string {
	return buildInfo
}
