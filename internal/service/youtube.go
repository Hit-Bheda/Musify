package service

import "os/exec"

func DownloadVideo(url string) error {
	cmd := exec.Command(
		"yt-dlp",
		"--embed-metadata",
		"--embed-thumbnail",
		"--add-metadata",
		"-f", "bv*+ba/b",
		url,
	)

	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
