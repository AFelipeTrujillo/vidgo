package Persistence

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AFelipeTrujillo/vidgo/internal/Domain/Entity"
	"github.com/tidwall/gjson"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type FFmpegProcessor struct{}

func NewFFmpegProcessor() *FFmpegProcessor {
	return &FFmpegProcessor{}
}

func (p *FFmpegProcessor) GetMetadata(ctx context.Context, file_path string) (*Entity.Video, error) {
	data, err := ffmpeg_go.Probe(file_path)

	if err != nil {
		return nil, fmt.Errorf("ffmpeg probe failed: %w", err)
	}

	total_duration_str := gjson.Get(data, "format.duration").String()
	total_duration, err := strconv.ParseFloat(total_duration_str, 64)
	if err != nil {
		return nil, fmt.Errorf("ffmpeg probe failed: %w", err)
	}

	format_name := gjson.Get(data, "format.format_name").String()

	return Entity.NewVideo(file_path, format_name, total_duration), nil
}

func (p *FFmpegProcessor) Trim(ctx context.Context, clip *Entity.Clip) error {
	stream_input := ffmpeg_go.Input(clip.SourceVideo.FilePath, ffmpeg_go.KwArgs{"ss": clip.StartTime.String()})

	stream_output := stream_input.Output(clip.OutputFile, ffmpeg_go.KwArgs{
		"t":   strconv.Itoa(clip.DurationInSeconds),
		"c:v": "copy",
		"c:a": "copy",
	})

	err := stream_output.Run()

	if err != nil {
		return fmt.Errorf("failed to trim video: %w", err)
	}

	return nil
}

func (p *FFmpegProcessor) Concatenate(ctx context.Context, clips []*Entity.Clip, outputName string) error {
	// For now, we leave it as a "Not Implemented" error so the compiler is happy
	// but the app doesn't crash if called.
	return fmt.Errorf("concatenate feature is coming soon in the next sprint")
}
