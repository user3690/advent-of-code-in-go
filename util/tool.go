package util

import (
	"strconv"
)

type Integer interface {
	int64 | int32 | int16 | int8 | int
}

func StrToInt64(str string) (number int64, err error) {
	number, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		return number, err
	}

	return number, err
}

func PowInt[T Integer](n T, m T) T {
	var i T

	if m == 0 {
		return 1
	}

	result := n
	for i = 2; i <= m; i++ {
		result *= n
	}

	return result
}
