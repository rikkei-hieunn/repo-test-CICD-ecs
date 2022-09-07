package string_utils

import (
	"repo-test-CICD-S3/model"
	"strconv"
)

//ToString convert interface to string
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		if v == model.EmptyString {
			return model.NullValue
		}

		return v
	case []uint8:

		return string(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case int:
		return strconv.Itoa(v)
	case bool:
		if v {
			return model.TrueValue
		}

		return model.FalseValue
	default:
		return model.NullValue
	}
}
