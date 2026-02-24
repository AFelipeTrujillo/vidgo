package ValueObject

import (
	"errors"
	"regexp"
)

type Timestamp struct {
	value string
}

func NewTimestamp(time string) (Timestamp, error) {
	re := regexp.MustCompile(`^(\d{1,2}:)?([0-5]?\d):([0-5]?\d)$`)

	if !re.MatchString(time) {
		return Timestamp{}, errors.New("Invalid timestamp format: must be HH:MM:SS")
	}

	return Timestamp{value: time}, nil
}

func (t Timestamp) String() string {
	return t.value
}
