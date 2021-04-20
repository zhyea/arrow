package vibe

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

func desensitizeMap(m map[string]interface{}) {
	for key, val := range m {
		switch val.(type) {
		case map[interface{}]interface{}:
			// nested map: cast and recursively unify
			val = cast.ToStringMap(val)
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
	result := map[string]interface{}{}
	for k, v := range src {
		result[fmt.Sprintf("%v", k)] = v
	}
	return result
}

// mergeMaps merge two maps.
func mergeMaps(target, src map[string]interface{}) {
	for srcKey, srcVal := range src {
		tgtKey := keyExists(target, srcKey)
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
			continue
		}

		switch tgtValType := tgtVal.(type) {
		case map[interface{}]interface{}:
			srcValType := srcVal.(map[interface{}]interface{})
			ssv := castToMapStringInterface(srcValType)
			stv := castToMapStringInterface(tgtValType)
			mergeMaps(stv, ssv)
		case map[string]interface{}:
			mergeMaps(tgtValType, srcVal.(map[string]interface{}))
		default:
			target[tgtKey] = srcVal
		}
	}
}

func keyExists(m map[string]interface{}, key string) string {
	lowerKey := strings.ToLower(key)
	for k := range m {
		lk := strings.ToLower(k)
		if lk == lowerKey {
			return k
		}
	}
	return ""
}

//mergeKeys merge all the keys in config
func mergeKeys(shadow map[string]bool, m map[string]interface{}, prefix string, keyDelimiter string) map[string]bool {
	if shadow != nil && prefix != "" && shadow[prefix] {
		return shadow
	}

	if shadow == nil {
		shadow = make(map[string]bool)
	}

	var m2 map[string]interface{}
	if prefix != "" {
		prefix += keyDelimiter
	}
	for k, val := range m {
		fullKey := prefix + k
		switch val.(type) {
		case map[string]interface{}:
			m2 = val.(map[string]interface{})
		case map[interface{}]interface{}:
			m2 = cast.ToStringMap(val)
		default:
			shadow[fullKey] = true
			continue
		}
		shadow = mergeKeys(shadow, m2, fullKey, keyDelimiter)
	}
	return shadow
}

//searchMap recursively searches for a value for path in source map.
func searchMap(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	next, ok := source[path[0]]
	if ok {
		if len(path) == 1 {
			return next
		}
		switch next.(type) {
		case map[interface{}]interface{}:
			return searchMap(cast.ToStringMap(next), path[1:])
		case map[string]interface{}:
			return searchMap(next.(map[string]interface{}), path[1:])
		default:
			return nil
		}
	}

	return nil
}

//searchMapWithPathPrefixes recursively searches for a value for path in source map with path prefix.
func searchMapWithPathPrefixes(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	for i := len(path); i > 0; i-- {
		prefixKey := strings.ToLower(strings.Join(path[0:i], v.keyDelimiter))

		next, ok := source[prefixKey]
		if ok {
			// Fast path
			if i == len(path) {
				return next
			}
			// Nested case
			var val interface{}
			switch next.(type) {
			case map[interface{}]interface{}:
				val = searchMapWithPathPrefixes(cast.ToStringMap(next), path[i:])
			case map[string]interface{}:
				// Type assertion is safe here since it is only reached
				// if the type of `next` is the same as the type being asserted
				val = searchMapWithPathPrefixes(next.(map[string]interface{}), path[i:])
			default:
				// got a value but nested key expected, do nothing and look for next prefix
			}
			if val != nil {
				return val
			}
		}
	}

	// not found
	return nil
}

// deepSearch scans deep maps, following the key indexes listed in the
// sequence "path".
// The last value is expected to be another map, and is returned.
//
// In case intermediate keys do not exist, or map to a non-map value,
// a new map is created and inserted, and the search continues from there:
// the initial map "m" may be modified!
func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			// intermediate key does not exist
			// => create it and continue from there
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]interface{})
		if !ok {
			// intermediate key is a value
			// => replace with a new map
			m3 = make(map[string]interface{})
			m[k] = m3
		}
		// continue search from here
		m = m3
	}
	return m
}
