package vibe

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const (
	EXT_YAML = "yml"
)

var v *Vibe

func init() {
	v = New()
}

// StringReplacer applies a set of replacements to a string.
type StringReplacer interface {
	// Replace returns a copy of s with all replacements performed.
	Replace(s string) string
}

type Vibe struct {

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

//New returns an initialized Vibe instance.
//
func New() *Vibe {

	v := new(Vibe)
	v.configName = "config"
	v.configPermissions = os.FileMode(0644)
	v.config = make(map[string]interface{})
	v.override = make(map[string]interface{})
	v.defaults = make(map[string]interface{})
	v.kvStore = make(map[string]interface{})
	v.env = make(map[string][]string)
	v.aliases = make(map[string]string)
	v.typeByDefValue = false

	return v
}

// AddConfigFile explicitly defines the path, name and extension of the config file.
// Can be called multiple times to add multiple config files.
func AddConfigFile(in string) {
	v.AddConfigFile(in)
}

func (v *Vibe) AddConfigFile(in string) {
	if in != "" {
		v.configFiles = append(v.configFiles, in)
	}
}

func (v *Vibe) ReadInConfig() error {
	for _, file := range v.configFiles {
		config := make(map[string]interface{})
		err := v.readInConfig(file, &config)
		if nil != err {
			return err
		}
		unifyMap(config)
		//TODO Add Config
	}
	return nil
}

func (v *Vibe) readInConfig(fileName string, config *map[string]interface{}) error {

	ext := filepath.Ext(fileName)

	if EXT_YAML != ext {
		return UnsupportedConfigError(ext)
	}

	file, err := os.ReadFile(fileName)
	if nil != err {
		return ConfigParseError{err}
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return ConfigParseError{err}
	}

	return nil
}
