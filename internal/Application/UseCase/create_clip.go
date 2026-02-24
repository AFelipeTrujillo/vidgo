package UseCase

import (
	"context"
	"fmt"

	"github.com/AFelipeTrujillo/vidgo/internal/Application/DTO"
	"github.com/AFelipeTrujillo/vidgo/internal/Domain/Entity"
	"github.com/AFelipeTrujillo/vidgo/internal/Domain/Repository"
	"github.com/AFelipeTrujillo/vidgo/internal/Domain/ValueObject"
)

type CreateClipUseCase struct {
	processor Repository.VideoProcessor
}

func NewCreateClipUseCase(p Repository.VideoProcessor) *CreateClipUseCase {
	return &CreateClipUseCase{processor: p}
}

func (us *CreateClipUseCase) Execute(ctx context.Context, req DTO.CreateClipRequest) (DTO.CreateClipResponse, error) {
	video, err := us.processor.GetMetadata(ctx, req.InputPath)

	if err != nil {
		return DTO.CreateClipResponse{Success: false}, err
	}

	start, err := ValueObject.NewTimestamp(req.StartTime)

	if err != nil {
		return DTO.CreateClipResponse{Success: false}, err
	}

	clip, err := Entity.NewCLip(video, start, req.DurationInSeconds, req.OutputPath)
	if err != nil {
		return DTO.CreateClipResponse{Success: false}, err
	}

	err = us.processor.Trim(ctx, clip)
	if err != nil {
		return DTO.CreateClipResponse{Success: false}, fmt.Errorf("ffmpeg error: %w", err)
	}

	return DTO.CreateClipResponse{
		Success:    true,
		OutputPath: clip.OutputFile,
		Message:    "Clip created successfully!",
	}, nil

}
