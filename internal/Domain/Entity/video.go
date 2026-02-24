package Entity

type Video struct {
	ID                string
	FilePath          string
	Format            string
	DurationInSeconds float64
}

func NewVideo(id, file_path, format string, duration_in_seconds float64) *Video {
	return &Video{
		ID:                id,
		FilePath:          file_path,
		Format:            format,
		DurationInSeconds: duration_in_seconds,
	}
}
