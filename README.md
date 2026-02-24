# vidgo

`vidgo` is a high-performance CLI tool written in Go for seamless video manipulation, powered by FFmpeg and built with Clean Architecture principles.

## Features (MVP)
- **Clip:** Extract segments from videos without quality loss.
- **Trailer:** Combine multiple clips into a single highlight reel.
- **Snapshot:** Generate quick previews or thumbnails.

## Architecture
The project follows **Clean Architecture** to ensure maintainability and testability:
- **Domain:** Core business logic and entities.
- **Application:** Use cases and orchestration.
- **Infrastructure:** FFmpeg integration and CLI delivery (Cobra).

## Prerequisites
- [Go](https://go.dev/) 1.21+
- [FFmpeg](https://ffmpeg.org/) installed and added to your PATH.

## Installation
```bash
go build -o vidgo ./cmd/app/main.go