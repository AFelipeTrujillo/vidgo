package Unit

import (
	"context"
	"testing"

	"github.com/AFelipeTrujillo/vidgo/internal/Application/DTO"
	"github.com/AFelipeTrujillo/vidgo/internal/Application/UseCase"
	"github.com/AFelipeTrujillo/vidgo/internal/Domain/Entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClipUseCase_Execute_Success(t *testing.T) {
	mockProcessor := new(VideoProcessorMock)
	useCase := UseCase.NewCreateClipUseCase(mockProcessor)
	ctx := context.Background()

	dummyVideo := Entity.NewVideo("input.mp4", "mp4", 60.0)

	mockProcessor.On("GetMetadata", ctx, "input.mp4").Return(dummyVideo, nil)
	mockProcessor.On("Trim", ctx, mock.AnythingOfType("*Entity.Clip")).Return(nil)

	req := DTO.CreateClipRequest{
		InputPath:         "input.mp4",
		StartTime:         "00:00:10",
		DurationInSeconds: 5,
		OutputPath:        "out.mp4",
	}

	res, err := useCase.Execute(ctx, req)

	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "out.mp4", res.OutputPath)
	mockProcessor.AssertExpectations(t)
}

func TestCreateClipUseCase_InvalidTimestamp(t *testing.T) {
	mockProcessor := new(VideoProcessorMock)
	useCase := UseCase.NewCreateClipUseCase(mockProcessor)
	ctx := context.Background()

	dummyVideo := Entity.NewVideo("video.mp4", "mp4", 60.0)
	mockProcessor.On("GetMetadata", ctx, "video.mp4").Return(dummyVideo, nil)

	req := DTO.CreateClipRequest{
		InputPath:         "video.mp4",
		StartTime:         "invalid-time", // This should fail
		DurationInSeconds: 10,
	}

	res, err := useCase.Execute(ctx, req)

	assert.Error(t, err)
	assert.False(t, res.Success)
	mockProcessor.AssertExpectations(t)

}
