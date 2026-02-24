package Entity

import (
	"errors"

	"github.com/AFelipeTrujillo/vidgo/internal/Domain/ValueObject"
)

type Clip struct {
	SourceVideo       *Video
	StartTime         ValueObject.Timestamp
	DurationInSeconds int
	OutputFile        string
}

func NewCLip(video *Video, start ValueObject.Timestamp, duration_in_seconds int, output_file string) (*Clip, error) {

	if duration_in_seconds <= 0 {
		return nil, errors.New("Clip duration must be greater than zero")
	}

	return &Clip{
		SourceVideo:       video,
		StartTime:         start,
		DurationInSeconds: duration_in_seconds,
		OutputFile:        output_file,
	}, nil
}
