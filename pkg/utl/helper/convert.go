package helper

import (
	"fmt"
	"reflect"
	"strconv"
)

//AsUInt - convert Uint
func AsUInt(input interface{}) uint {
	var results uint
	if reflect.ValueOf(input).Kind() == reflect.String {
		results1, err := strconv.ParseUint(input.(string), 10, 32)
		if err != nil {
			panic(fmt.Sprintf("Error not convert type %v to %v", reflect.String, reflect.Uint))
		} else {
			results = uint(results1)
		}
	} else if reflect.ValueOf(input).Kind() == reflect.Uint {
		results = uint(input.(uint))
	} else {
		results = uint(input.(float64))
	}
	return results
}

//AsInt - convert Int
func AsInt(input interface{}) int {
	var results int
	if reflect.ValueOf(input).Kind() == reflect.String {
		results1, err := strconv.ParseInt(input.(string), 10, 32)
		if err != nil {
			panic(fmt.Sprintf("Error not convert type %v to %v", reflect.String, reflect.Int))
		} else {
			results = int(results1)
		}
	} else if reflect.ValueOf(input).Kind() == reflect.Uint {
		results = int(input.(int))
	} else {
		results = int(input.(float64))
	}
	return results
}
