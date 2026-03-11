package processor

import (
	"fmt"
	"strconv"
	"strings"
)

func FilterData(data any, condition string) (any, error) {
	slice, ok := data.([]any)
	if !ok {
		return nil, fmt.Errorf("Filter only for array")
	}
	//[task > 5 ] ->["task",">","5"]
	parts := strings.Fields(condition)
	if len(parts) > 3 {
		return nil, fmt.Errorf("Not suitable ( example : age >18 )")
	}
	key, op, valStr := parts[0], parts[1], parts[2]
	var filtered []any
	for _, item := range slice {
		obj, ok := item.(map[string]any)
		if !ok {
			continue
		}

		targetVal, exist := obj[key]
		if !exist {
			continue
		}

		if matchCondition(targetVal, op, valStr) {
			filtered = append(filtered, item)
		}

	}
	return filtered, nil
}

func matchCondition(val any, op string, condValStr string) bool {
	switch v := val.(type) {
	case float64: // JSON numbers is auto  float64 in Go
		cv, _ := strconv.ParseFloat(condValStr, 64)
		switch op {
		case ">":
			return v > cv
		case "<":
			return v < cv
		case "==":
			return v == cv
		}
	case string:
		switch op {
		case "==":
			return v == condValStr
		}
	}
	return false
}
