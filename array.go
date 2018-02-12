package main

import (
	"fmt"
	"reflect"
)

func InArray(tar interface{}, c interface{}) (bool, int) {
	tarValue := reflect.ValueOf(tar)

	switch reflect.TypeOf(tar).Kind() {
	case reflect.Slice, reflect.Array:
		for i = 0; i < tarValue.Len(); i++ {
			if tarValue.Index(i).Interface() == c {
				return true, i
			}
		}

	case reflect.Map:
		if tarValue.MapIndex(reflect.ValueOf(c)).IsValid() {
			return true, 0
		}
	}

	return false, -1
}

func main() {

}
