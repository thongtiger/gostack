package util

import (
	"encoding/json"
	"log"
	"reflect"
	"regexp"
	"strconv"
)

type Map map[string]interface{}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func (c *Map) ToString() string {
	json, err := json.Marshal(c)
	if err != nil {
		return "{}"
	}
	return string(json)
}

// IsZero reports whether v is zero struct
func IsZero(v interface{}) bool {
	value := reflect.ValueOf(v)
	if !value.IsValid() {
		return true
	}

	switch value.Kind() {
	case reflect.Bool:
		return value.Bool() == false

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return value.Float() == 0

	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == 0

	case reflect.Ptr, reflect.Interface:
		return IsZero(value.Elem())

	case reflect.Array:
		for i := 0; i < value.Len(); i++ {
			if !IsZero(value.Index(i)) {
				return false
			}
		}
		return true

	case reflect.Slice, reflect.String, reflect.Map:
		return value.Len() == 0

	case reflect.Struct:
		for i, n := 0, value.NumField(); i < n; i++ {
			if !IsZero(value.Field(i)) {
				return false
			}
		}
		return true
	// reflect.Chan, reflect.UnsafePointer, reflect.Func
	default:
		return value.IsNil()
	}
}

func UsernameValid(str string) bool {
	reUser := regexp.MustCompile(`^(?:[a-z])[a-z0-9]{5,15}$`)
	return reUser.MatchString(str)
}

func PasswordValid(str string) bool {
	rePass := regexp.MustCompile(`^(?:[a-zA-Z0-9])[a-zA-Z0-9]{6,30}$`)
	return rePass.MatchString(str)
}

func ToStr(input int) string {
	return strconv.Itoa(input)
}
