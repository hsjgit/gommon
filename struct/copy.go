package _struct

import (
	"fmt"
	"reflect"
)

// 将src的值复制给target
func CopyFields(target interface{}, src interface{}, fields ...string) (err error) {
	targetType := reflect.TypeOf(target)
	targetValue := reflect.ValueOf(target)
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)

	// 简单判断下
	if targetType.Kind() != reflect.Ptr {
		return fmt.Errorf("target must be target struct pointer")
	}
	targetValue = reflect.ValueOf(targetValue.Interface())

	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < srcValue.NumField(); i++ {
			_fields = append(_fields, srcType.Field(i).Name)
		}
	}

	if len(_fields) == 0 {
		return fmt.Errorf("no fields to copy")

	}

	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := targetValue.Elem().FieldByName(name)
		bValue := srcValue.FieldByName(name)

		// src中有同名的字段并且类型一致才复制
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			continue
		}
	}
	return
}
