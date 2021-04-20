package vibe

import "github.com/mitchellh/mapstructure"

//decode
func decode(input interface{}, rawVal interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           rawVal,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
