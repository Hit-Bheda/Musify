package main

import (
	"log"

	"github.com/Hit-Bheda/Musify/internal/service"
)

const (
	windowSize = 1024
	hopSize    = 512
)

func main() {
	service.ConvertToMono48k("song1.wav", "output.wav")
	data, bitDepth, err := service.Sample("output.wav")
	if err != nil {
		log.Fatalf("Failed to sample audio : %v", err)
	}

	newData := service.Normalize(data, bitDepth)

	spec := service.Spectrogram(newData, windowSize, hopSize)
	service.PlotSpectrogram(spec, "spectrogram.png")
}
