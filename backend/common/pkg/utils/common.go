package utils

import "strconv"

func ConverStr2Uint(str string) (uint, error) {
	uit, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(uit), nil
}
