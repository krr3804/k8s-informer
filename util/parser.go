package util

import (
	"fmt"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Int32PointerToString(i *int32) string {
	if i == nil {
		return "None"
	}

	return strconv.Itoa(int(*i))
}

func MetaV1TimeToString(t *metav1.Time) string {
	if t == nil {
		return "None"
	}

	return t.String()
}

func StringPointerToString(s *string) string {
	if s == nil {
		return "None"
	}
	return *s
}

func MapToString(m map[string]string, sep string) string {
	var result string
	count := 0
	length := len(m)
	for key, value := range m {
		count++
		result += fmt.Sprintf("%s%s%s", key, sep, value)

		if count < length {
			result += ","
		}
	}

	return result
}

func BoolPointerToBool(b *bool, def bool) bool {
	if b == nil {
		return def
	}
	return *b
}
