package gfun

import (
	"encoding/json"
	"fmt"
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
