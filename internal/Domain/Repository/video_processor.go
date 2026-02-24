package Repository

import (
	"context"

	"github.com/AFelipeTrujillo/vidgo/internal/Domain/Entity"
)

type VideoProcessor interface {
	Trim(ctx context.Context, clip *Entity.Clip) error
	Concatenate(ctx context.Context, clips []*Entity.Clip, output_name string) error
	GetMetadata(ctx context.Context, file_path string) (*Entity.Video, error)
}
