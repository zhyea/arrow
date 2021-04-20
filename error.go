package vibe

import "fmt"

// ConfigParseError denotes failing to parse configuration file.
type ConfigParseError struct {
	err error
}

// Error returns the formatted configuration error.
func (e ConfigParseError) Error() string {
	return fmt.Sprintf("While parsing config: %s", e.err.Error())
}

// UnsupportedConfigError denotes encountering an unsupported configuration filetype.
type UnsupportedConfigError string

// Error returns the formatted configuration error.
func (err UnsupportedConfigError) Error() string {
	return fmt.Sprintf("Unsupported Config Type %q", string(err))
}

// ConfigNotFoundError denotes encountering an unsupported
type ConfigNotFoundError string

func (err ConfigNotFoundError) Error() string {
	return fmt.Sprintf("Config Not found %q", string(err))
}
