package vibe

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

func desensitizeMap(m map[string]interface{}) {
	for key, val := range m {
		switch val.(type) {
		case map[interface{}]interface{}:
			// nested map: cast and recursively unify
			val, _ = ToStringMap(val)
			desensitizeMap(val.(map[string]interface{}))
		case map[string]interface{}:
			// nested map: recursively unify
			desensitizeMap(val.(map[string]interface{}))
		}

		lower := strings.ToLower(key)
		if key != lower {
			// remove old key (not lower-cased)
			delete(m, key)
		}
		// update map
		m[lower] = val
	}
}

func castToMapStringInterface(src map[interface{}]interface{}) map[string]interface{} {
	tgt := map[string]interface{}{}
	for k, v := range src {
		tgt[fmt.Sprintf("%v", k)] = v
	}
	return tgt
}

// mergeMaps merges two maps.
//The `itgt` parameter is for handling go-yaml's insistence on parsing nested structures as `map[interface{}]interface{}`
// instead of using a `string` as the key for nest structures beyond one level
// deep. Both map types are supported as there is a go-yaml fork that uses
// `map[string]interface{}` instead.
func mergeMaps(src, target map[string]interface{}) {
	for srcKey, srcVal := range src {
		tgtKey := keyExists(srcKey, target)
		if tgtKey == "" {
			target[srcKey] = srcVal
			continue
		}

		tgtVal, ok := target[tgtKey]
		if !ok {
			target[srcKey] = srcVal
			continue
		}

		svType := reflect.TypeOf(srcVal)
		tvType := reflect.TypeOf(tgtVal)
		if svType != tvType {
			log.Printf("svType != tvType; key=%s, st=%v, tt=%v, srcVal=%v, tgtVal=%v", srcKey, svType, tvType, srcVal, tgtVal)
			continue
		}

		switch tgtValType := tgtVal.(type) {
		case map[interface{}]interface{}:
			srcValType := srcVal.(map[interface{}]interface{})
			ssv := castToMapStringInterface(srcValType)
			stv := castToMapStringInterface(tgtValType)
			mergeMaps(ssv, stv)
		case map[string]interface{}:
			log.Printf("merging maps")
			mergeMaps(srcVal.(map[string]interface{}), tgtValType)
		default:
			log.Printf("setting value")
			target[tgtKey] = srcVal
		}
	}
}

func keyExists(key string, m map[string]interface{}) string {
	lowerLey := strings.ToLower(key)
	for mk := range m {
		lmk := strings.ToLower(mk)
		if lmk == lowerLey {
			return mk
		}
	}
	return ""
}
