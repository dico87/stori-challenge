package format

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const dateSeparator = "/"

var ErrNotValidFormat = errors.New("not valid format format")

func ParseDate(date string) (*time.Time, error) {
	fields := strings.Split(date, dateSeparator)
	if len(fields) < 2 {
		return nil, ErrNotValidFormat
	}

	if len(fields) > 3 {
		return nil, ErrNotValidFormat
	}

	dateAsString := fmt.Sprintf("%d-%s-%s", time.Now().Year(), fmt.Sprintf("%02s", fields[0]), fmt.Sprintf("%02s", fields[1]))
	if len(fields) == 3 {
		dateAsString = fmt.Sprintf("%s-%s-%s", fields[2], fmt.Sprintf("%02s", fields[0]), fmt.Sprintf("%02s", fields[1]))
	}

	dateAsTime, err := time.Parse(time.DateOnly, dateAsString)
	if err != nil {
		return nil, ErrNotValidFormat
	}

	return &dateAsTime, nil
}
