package DTO

type CreateClipRequest struct {
	InputPath         string
	StartTime         string // e.g., "00:01:30"
	DurationInSeconds int    // e.g., 10 (seconds)
	OutputPath        string
}

type CreateClipResponse struct {
	Success    bool
	OutputPath string
	Message    string
}
