package common

import (
	"strconv"
	"time"
)

func convertStrToInt (str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// error handling
	}
	return num

}

func ConvertStrToDay (str string) time.Time {
	startDay := time.Date(
		convertStrToInt(str[0:4]),
		time.Month(convertStrToInt(str[4:6])),
		convertStrToInt(str[6:8]),
		0, 0, 0, 0, time.Local)

	return startDay
}

