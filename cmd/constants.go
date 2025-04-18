package main

const (
	EnvBeaconProject string = "BEACON_PROJECT"
)

const (
	FlagProjectLong    string = "project"
	FlagProjectDesc    string = "Names the PROJECT as a source for the event"
	FlagProjectDefault string = "unknown"

	FlagStatusLong    string = "status"
	FlagStatusShort   string = "s"
	FlagStatusDesc    string = "Emits an event with a custom status"
	FlagStatusDefault string = ""

	FlagFailLong    string = "fail"
	FlagFailShort   string = "f"
	FlagFailDesc    string = "Emits a failure event"
	FlagFailDefault bool   = false

	FlagInfoLong    string = "info"
	FlagInfoShort   string = "i"
	FlagInfoDesc    string = "Emits an informational event"
	FlagInfoDefault bool   = false

	FlagPassLong    string = "pass"
	FlagPassShort   string = "p"
	FlagPassDesc    string = "Emits a successful event" // #nosec G101
	FlagPassDefault bool   = false

	FlagPermissionsLong    string = "permissions"
	FlagPermissionsDesc    string = "Displays IAM permissions required by the application"
	FlagPermissionsDefault bool   = false

	FlagConfigLong    string = "config"
	FlagConfigShort   string = "c"
	FlagConfigDesc    string = "Specifies a path to load a config file"
	FlagConfigDefault string = ""

	FlagGenerateConfigLong    string = "generate-config"
	FlagGenerateConfigDesc    string = "Displays a template for the config file"
	FlagGenerateConfigDefault bool   = false
)
