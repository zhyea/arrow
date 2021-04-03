package vibe

import "os"

// StringReplacer applies a set of replacements to a string.
type StringReplacer interface {
	// Replace returns a copy of s with all replacements performed.
	Replace(s string) string
}

type Viper struct {

	// A set of paths to look for the config file in
	configFiles []string

	// Name of file to look for inside the path
	configName        string
	configType        string
	configPermissions os.FileMode
	envPrefix         string

	automaticEnvApplied bool
	envKeyReplacer      StringReplacer
	allowEmptyEnv       bool

	config         map[string]interface{}
	override       map[string]interface{}
	defaults       map[string]interface{}
	kvStore        map[string]interface{}
	env            map[string][]string
	aliases        map[string]string
	typeByDefValue bool

	onConfigChange func()
}
