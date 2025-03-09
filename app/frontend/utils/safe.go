package utils

import "net/url"

var validHost = []string{
	"localhost:8080",
}

func ValidateNext(next string) bool {
	urlObj, err := url.Parse(next)
	if err != nil {
		return false
	}
	if InArray(urlObj.Host, validHost) {
		return true
	}
	return false
}

func InArray[T int | int32 | int64 | float32 | float64 | string](needle T, haystack []T) bool {
	for _, k := range haystack {
		if needle == k {
			return true
		}
	}
	return false
}
