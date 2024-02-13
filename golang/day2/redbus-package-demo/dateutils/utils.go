package dateutils

import "errors"

func IsLeap(year int) (bool, error) {
	if year < 0 {
		return false, errors.New("Invalid year")
	}
	return year%400 == 0 || year%4 == 0 && year%100 != 0, nil
}
