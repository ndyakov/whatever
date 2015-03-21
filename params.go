package whatever

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Params is a type on top of map[string]interface{}
// that has a few useful getters and other methods.
// This type can help you to work with JSON data that
// can be unmarshaled into it and then extracted with
// the appropriate type (it even has a getter for time.Time).
// Also the Required and Empty methods may be helpful with
// validating data. It is possible to work with nested Params,
// getting an Interface that you can cast yourself to the
// wanted type and adding new values to the Params structure.
type Params map[string]interface{}

// NewFromJSON receives a slice of bytes that should be json
// body and returns Params structure for that json body and an
// error if there was one while decoding the json.
func NewFromJSON(jsonBody []byte) (Params, error) {
	var p Params
	err := json.Unmarshal(jsonBody, &p)
	return p, err
}

// Add adds a new pair(key, value) to the Params structure.
// Returns true if an existing value was overwritten.
func (p Params) Add(key string, value interface{}) bool {
	if fmt.Sprintf("%p", p) == fmt.Sprintf("%p", value) {
		return false
	}

	_, ok := p[key]
	p[key] = value
	return ok
}

// Empty checks if the Params structure is empty.
// Returns false if a there are elements in the structure.
func (p Params) Empty() bool {
	for _ = range p {
		return false
	}
	return true
}

// Remove deletes a key from the structure.
// It won't fail even if called on a missing key.
func (p Params) Remove(key string) {
	delete(p, key)
}

// GetP will return Params structure if the value with
// the specified key is of either map[string]interface{} or Params type.
// Returns empty Params if the key is missing or if the value
// was not one of the desired types.
func (p Params) GetP(key string) Params {
	if val, ok := p[key]; ok {
		if vmap, ok := val.(map[string]interface{}); ok {
			return Params(vmap)
		} else if vp, ok := val.(Params); ok {
			return vp
		}
	}

	return Params{}
}

// GetI will return the same as map[key] would have returned.
// Returns an interface that you can try to cast to other types.
// If there is no value with that key it will return nil.
func (p Params) GetI(key string) interface{} {
	if val, ok := p[key]; ok {
		return val
	}

	return nil
}

// Get returns a string representation of the value with
// the specified key.
// Returns empty string if there is no value with the provided key.
func (p Params) Get(key string) string {
	if val, ok := p[key]; ok {
		return stringify(val)
	}

	return ""
}

// GetString returns a string only if the value with the specified key
// can be casted to string. Will return an empty string otherwise.
func (p Params) GetString(key string) string {
	if val, ok := p[key]; ok {
		if vs, ok := val.(string); ok {
			return vs
		}
	}
	return ""
}

// GetInt parses the value with the provided key to an int.
// If there is an error with the parsing, returns 0.
func (p Params) GetInt(key string) int {
	result, err := strconv.ParseInt(p.Get(key), 0, 0)
	if err == nil {
		return int(result)
	}

	return 0
}

// GetInt8 parses the value with the provided key to an int8.
// If there is an error with the parsing, returns 0.
func (p Params) GetInt8(key string) int8 {
	result, err := strconv.ParseInt(p.Get(key), 0, 8)
	if err == nil {
		return int8(result)
	}

	return 0
}

// GetInt64 parses the value with the provided key to an int64.
// If there is an error with the parsing, returns 0.
func (p Params) GetInt64(key string) int64 {
	result, err := strconv.ParseInt(p.Get(key), 0, 64)
	if err == nil {
		return result
	}

	return 0
}

// GetFloat32 parses the value with the provided key to an float32.
// If there is an error with the parsing, returns 0.
func (p Params) GetFloat32(key string) float32 {
	result, err := strconv.ParseFloat(p.Get(key), 32)
	if err == nil {
		return float32(result)
	}

	return 0
}

// GetFloat64 parses the value with the provided key to an float64.
// If there is an error with the parsing, returns 0.
func (p Params) GetFloat64(key string) float64 {
	result, err := strconv.ParseFloat(p.Get(key), 64)
	if err == nil {
		return result
	}

	return 0
}

// GetFloat parses the value with the provided key to an float32.
// If there is an error with the parsing, returns 0.
// This is actually the same as calling GetFloat32.
func (p Params) GetFloat(key string) float32 {
	return p.GetFloat32(key)
}

// GetTime will return the value with the specified key
// parsed as time.Time structure. The layout is time.RFC3339
// or in other words the JSON format for the Date object in JavaScript.
// The value should look like this:
//     "2015-02-27T21:53:57.582Z"
// Otherwise returns time.Time{}
func (p Params) GetTime(key string) time.Time {
	if result, err := time.Parse(time.RFC3339, p.Get(key)); err == nil {
		return result
	}
	return time.Time{}
}

// GetSlice returns a slice of interface{} if the value
// with the provided key can be casted to a slice.
// Otherwise returns nil.
func (p Params) GetSlice(key string) []interface{} {
	if val, ok := p[key]; ok {
		if val, ok := val.([]interface{}); ok {
			return val
		}
	}

	return nil
}

// GetSliceStrings will return a slice of strings.
// This slice of strings will be a result of casting the value
// corresponding to the provided key to a slice of interface{}
// and then individually casting each element to a string.
// Only those elements that can be casted to a string will be
// present in the final result. Those that cannot be casted to string
// will be silently ignored. If the is not value with that
// key or the value is not a slice, nil will be returned.
func (p Params) GetSliceStrings(key string) []string {
	var result []string
	if val, ok := p[key]; ok {
		if val, ok := val.([]interface{}); ok {
			for _, v := range val {
				if vs, ok := v.(string); ok {
					result = append(result, vs)
				}
			}
		}
		return result
	}

	return nil
}

// GetSliceInts will return a slice of strings.
// This slice of strings will be a result of casting the value
// corresponding to the provided key to a slice of interface{}
// and then individually casting each element to an int.
// Only those elements that can be casted to an int will be
// present in the final result. Those who cannot be casted to int
// will be silently ignored. If the is not value with that
// key or the value is not a slice, nil will be returned.
func (p Params) GetSliceInts(key string) []int {
	var result []int
	if val, ok := p[key]; ok {
		if slice, ok := val.([]interface{}); ok {
			for _, v := range slice {
				if intVal, err := strconv.ParseInt(stringify(v), 0, 0); err == nil {
					result = append(result, int(intVal))
				}
			}
		}
		return result
	}

	return nil
}

// URLValues return the values in the Params structure
// as url.Values that can be then used with packages as
// gorilla`s schema or goji`s params. The schema and params
// packages expects different keys for the nested values. For
// this to be able to generate the expected keys you may need to
// pass the prefix and the suffix of the nested key. If the prefix
// is empty string, then the prefix will be set to dot (".") and
// the suffix will be an empty string. This is the way that gorilla`s
// schema expect the nested keys to be represented. Example:
//     some_key.inner_key.last_key
// This somewhat resembles the mongo notation of nested objects as well.
// To be able to use this with params, you will need to pass an prefix "["
// and suffix "]". The key from the previous example will now be:
//     some_key[inner_key][last_key]
//
// If the Params structure is blank, then an empty url.Values will be returned.
func (p Params) URLValues(prefix, suffix string) url.Values {
	return toURLValues(p, prefix, suffix, false)
}

// Required will return an error if one of the passed keys is missing
// in the Params structure. As far as Required cares - an empty string
// is the same as missing value.
// To validate nested parameters please use the dotted notation:
//     some_key.nested_key.last_key
// Returns an error message of the following type:
//     "the parameter {key} is required"
// If all keys are present will return nil.
func (p Params) Required(keys ...string) error {
	for _, key := range keys {
		if ok := exists(p, key); !ok {
			return fmt.Errorf("the parameter %s is required", key)
		}
	}

	return nil
}

// Keys will return the top-level keys of the params.
// Keep in mind if you need the nested keys as well you can
// use NestedKeys.
func (p Params) Keys() []string {
	return keys(p)
}

// NestedKeys will return all keys in the params map.
// The nested keys will be prefixed with a dot.
// Example:
//     { "one": { "two": 3 } }
// Will result in the following key:
//     "one.two"
func (p Params) NestedKeys() []string {
	return nestedKeys(p, "", false)
}

func (p Params) Merge(set map[string]interface{}) {
	for k, v := range set {
		p[k] = v
	}
}

func (p Params) Defaults(set map[string]interface{}) {
	for k, v := range set {
		if _, ok := p[k]; !ok {
			p[k] = v
		}
	}
}

func exists(input map[string]interface{}, key string) (ok bool) {
	if input == nil {
		return false
	}

	if index := strings.Index(key, "."); index != -1 {
		var casted bool
		var params map[string]interface{}
		pair := strings.SplitN(key, ".", 2)

		if params, casted = input[pair[0]].(Params); !casted {
			params, casted = input[pair[0]].(map[string]interface{})
		}

		if !casted {
			return false
		}

		ok = exists(params, pair[1])
	} else {
		var v interface{}
		v, ok = input[key]
		if s, isS := v.(string); isS && ok {
			ok = (s != "")
		}
	}

	return ok
}

func stringify(v interface{}) string {
	if vs, ok := v.(string); ok {
		return vs
	}

	return fmt.Sprintf("%v", v)
}

func keys(set Params) (result []string) {
	for k := range set {
		result = append(result, k)
	}
	return
}

func nestedKeys(set map[string]interface{}, parent string, subParse bool) []string {
	var key string
	var result []string
	for k, v := range set {

		if subParse {
			key = fmt.Sprintf("%s.%s", parent, k)
		} else {
			key = k
		}

		result = append(result, key)

		if val, ok := v.(map[string]interface{}); ok {
			subKeys := nestedKeys(val, key, true)
			result = append(result, subKeys...)
		} else if val, ok := v.(Params); ok {
			subKeys := nestedKeys(val, key, true)
			result = append(result, subKeys...)
		}

	}
	return result
}

func toURLValues(set Params, prefix, suffix string, subParse bool) url.Values {
	if prefix == "" {
		prefix = "."
		suffix = ""
	}

	result := url.Values{}
	var subset url.Values
	for key, value := range set {
		var foundSubset bool
		if v, ok := value.(Params); ok {
			subset = toURLValues(v, prefix, suffix, true)
			foundSubset = true
		} else if v, ok := value.(map[string]interface{}); ok {
			subset = toURLValues(Params(v), prefix, suffix, true)
			foundSubset = true
		} else if v, ok := value.([]interface{}); ok {
			for _, el := range v {
				result[key] = append(result[key], stringify(el))
			}
			continue
		}
		if foundSubset {
			for k, v := range subset {
				nestedKey := fmt.Sprintf("%s%s", key, k)
				if subParse {
					nestedKey = fmt.Sprintf("%s%s%s%s", prefix, key, suffix, k)
				}
				result[nestedKey] = append(result[nestedKey], v...)
			}
		} else {
			valueKey := key
			if subParse {
				valueKey = fmt.Sprintf("%s%s%s", prefix, key, suffix)
			}
			result[valueKey] = append(result[valueKey], stringify(value))
		}
	}
	return result
}
