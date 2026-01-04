package service

import (
	"log"
	"os"

	"github.com/go-audio/wav"
)

func Sample() []int {
	f, err := os.Open("output.wav")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	decoder := wav.NewDecoder(f)
	if !decoder.IsValidFile() {
		log.Fatal("Invalid wav file!")
	}

	buff, err := decoder.FullPCMBuffer()
	if err != nil {
		log.Fatal(err)
	}

	smaples := buff.Data
	return smaples
}
