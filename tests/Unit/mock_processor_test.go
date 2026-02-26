package Unit

import (
	"context"

	"github.com/AFelipeTrujillo/vidgo/internal/Domain/Entity"
	"github.com/stretchr/testify/mock"
)

type VideoProcessorMock struct {
	mock.Mock
}

func (m *VideoProcessorMock) Trim(ctx context.Context, clip *Entity.Clip) error {
	args := m.Called(ctx, clip)
	return args.Error(0)
}

func (m *VideoProcessorMock) GetMetadata(ctx context.Context, filePath string) (*Entity.Video, error) {
	args := m.Called(ctx, filePath)
	return args.Get(0).(*Entity.Video), args.Error(1)
}

func (m *VideoProcessorMock) Concatenate(ctx context.Context, clips []*Entity.Clip, output string) error {
	args := m.Called(ctx, clips, output)
	return args.Error(0)
}
