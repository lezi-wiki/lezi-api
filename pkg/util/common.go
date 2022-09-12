package util

import "strings"

func Replace(raw string, index map[string]string) string {
	for key, value := range index {
		raw = strings.ReplaceAll(raw, key, value)
	}

	return raw
}
