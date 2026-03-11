package processor

import (
	"fmt"
	"strconv"
	"strings"
)

func GetValueByPath(data any, path string) (any, error) {
	if path == "" {
		return data, nil
	}
	parts := strings.Split(path, ".")
	current := data
	for _, part := range parts {
		switch v := current.(type) {
		case map[string]any:
			val, ok := v[part]
			if !ok {
				return nil, fmt.Errorf("Cannot find key \n")
			}
			current = val

		case []any:
			idx, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("Index not suitable: %s", part)
			}
			if idx < 0 || idx >= len(v) {
				return nil, fmt.Errorf("index not in range: %d", idx)
			}
			current = v[idx]
		default:
			return nil, fmt.Errorf("cannot access data type %T", v)
		}

	}
	return current, nil
}
