package Entity

import "github.com/google/uuid"

type Video struct {
	ID                string
	FilePath          string
	Format            string
	DurationInSeconds float64
}

func NewVideo(file_path, format string, duration_in_seconds float64) *Video {
	return &Video{
		ID:                uuid.NewString(),
		FilePath:          file_path,
		Format:            format,
		DurationInSeconds: duration_in_seconds,
	}
}
