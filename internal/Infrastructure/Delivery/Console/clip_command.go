package Console

import (
	"context"
	"fmt"

	"github.com/AFelipeTrujillo/vidgo/internal/Application/DTO"
	"github.com/AFelipeTrujillo/vidgo/internal/Application/UseCase"
	"github.com/AFelipeTrujillo/vidgo/internal/Infrastructure/Persistence"
	"github.com/spf13/cobra"
)

func NewClipCommand() *cobra.Command {
	var input, output, start string
	var duration int

	cmd := &cobra.Command{
		Use:   "clip",
		Short: "Cut a specific segment from a video",
		Run: func(cmd *cobra.Command, args []string) {
			processor := Persistence.NewFFmpegProcessor()
			usecase := UseCase.NewCreateClipUseCase(processor)
			req := DTO.CreateClipRequest{
				InputPath:         input,
				OutputPath:        output,
				StartTime:         start,
				DurationInSeconds: duration,
			}

			res, err := usecase.Execute(context.Background(), req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			fmt.Printf("%s Saved to: %s\n", res.Message, res.OutputPath)

		},
	}

	cmd.Flags().StringVarP(&input, "input", "i", "", "Input video file (required)")
	cmd.Flags().StringVarP(&output, "output", "o", "output.mp4", "Output video file")
	cmd.Flags().StringVarP(&start, "start", "s", "00:00:00", "Start time (HH:MM:SS)")
	cmd.Flags().IntVarP(&duration, "duration", "d", 10, "Duration in seconds")

	cmd.MarkFlagRequired("input")

	return cmd
}
