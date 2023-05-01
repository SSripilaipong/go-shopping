package gfun

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func UnmarshalFromMap[T any](raw map[string]any, target *T) error {
	data, err := json.Marshal(raw)
	if err != nil {
		return fmt.Errorf("UnmarshalFromMap failed to marshal to json string: %w", err)
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("UnmarshalFromMap failed to unmarshal to struct: %w", err)
	}
	return nil
}

func MarshalToMap[T any](target T) (result map[string]any, err error) {
	data, err := json.Marshal(target)
	if err != nil {
		return nil, fmt.Errorf("MarshalToMap failed to marshal to json string: %w", err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("MarshalToMap failed to unmarshal to struct: %w", err)
	}

	return result, nil
}

func CheckAllStructFieldsSet[T any](body T) (bool, error) {
	structType := reflect.TypeOf(body)
	if structType.Kind() != reflect.Struct {
		return false, errors.New("unexpected error")
	}

	structVal := reflect.ValueOf(body)
	fieldNum := structVal.NumField()

	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		isSet := field.IsValid() && !field.IsZero()

		if !isSet {
			return false, nil
		}
	}
	return true, nil
}

func AllStructFieldsSet[T any](body T) bool {
	allSet, err := CheckAllStructFieldsSet(body)
	if err != nil {
		panic("unexpected error")
	}
	return allSet
}
