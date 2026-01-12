package service

import (
	"os/exec"

	"github.com/google/uuid"
)

func DownloadVideo(url string) error {
	id := uuid.New().String()
	cmd := exec.Command(
		"yt-dlp",
		"-x",
		"--audio-format", "wav",
		"-o", id+".%(ext)s",
		url,
	)

	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func ConvertToMono48k(input, output string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", input,
		"-ac", "1",
		"-ar", "48000",
		"-f", "wav",
		output,
	)

	cmd.Stdout = nil
	cmd.Stderr = nil

	return cmd.Run()
}
