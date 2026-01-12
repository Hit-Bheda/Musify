package service

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/go-audio/wav"
	"gonum.org/v1/gonum/dsp/fourier"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type SpecGrid struct {
	Data [][]float64
}

func (g SpecGrid) Dims() (c, r int) {
	if len(g.Data) == 0 {
		return 0, 0
	}
	return len(g.Data), len(g.Data[0])
}

func (g SpecGrid) Z(c, r int) float64 {
	return g.Data[c][r]
}

func (g SpecGrid) X(c int) float64 {
	return float64(c)
}

func (g SpecGrid) Y(r int) float64 {
	return float64(r)
}

func Sample(filename string) ([]int, int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	decoder := wav.NewDecoder(f)
	if !decoder.IsValidFile() {
		return nil, 0, fmt.Errorf("invalid wav file")
	}

	buff, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, 0, err
	}

	return buff.Data, int(decoder.BitDepth), nil
}

func Normalize(samples []int, bitDepth int) []float64 {
	// Dynamically calculate max value based on bit depth
	maxInt := float64(int64(1)<<(bitDepth-1) - 1)

	out := make([]float64, len(samples))
	for i, s := range samples {
		v := float64(s) / maxInt
		if v > 1.0 {
			v = 1.0
		} else if v < -1.0 {
			v = -1.0
		}
		out[i] = v
	}
	return out
}

func Int64ToFloat64PCM(samples []int, bitDepth int) []float64 {
	maxAmp := math.Pow(2, float64(bitDepth-1))

	out := make([]float64, len(samples))
	for i, s := range samples {
		out[i] = float64(s) / maxAmp
	}

	return out
}

func SaveFloat64SamplesToTXT(samples []float64, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, s := range samples {
		fmt.Fprintf(w, "%.10f\n", s)
	}

	return w.Flush()
}

func Spectrogram(samples []float64, winSize, hop int) [][]float64 {
	fft := fourier.NewFFT(winSize)
	window := Hann(winSize)

	var spec [][]float64

	// Iterate over samples with stride 'hop'
	for i := 0; i+winSize <= len(samples); i += hop {
		frame := make([]float64, winSize)
		for j := 0; j < winSize; j++ {
			frame[j] = samples[i+j] * window[j]
		}

		// fft.Coefficients returns complex coefficients.
		// For real input size N, output is N/2 + 1.
		coeffs := fft.Coefficients(nil, frame)

		bins := winSize / 2
		mag := make([]float64, bins)

		for k := 0; k < bins; k++ {
			re := real(coeffs[k])
			im := imag(coeffs[k])
			// Add small epsilon to avoid log(0)
			mag[k] = 20 * math.Log10(math.Sqrt(re*re+im*im)+1e-9)
		}

		spec = append(spec, mag)
	}

	return spec
}

func Hann(n int) []float64 {
	w := make([]float64, n)
	for i := range w {
		w[i] = 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/float64(n-1)))
	}
	return w
}

func PlotSpectrogram(spec [][]float64, out string) error {
	if len(spec) == 0 || len(spec[0]) == 0 {
		return fmt.Errorf("empty spectrogram")
	}

	p := plot.New()
	p.Title.Text = "Spectrogram"
	p.X.Label.Text = "Time (Frames)"
	p.Y.Label.Text = "Frequency (Bins)"

	grid := SpecGrid{Data: spec}

	pal := palette.Heat(12, 1)

	hm := plotter.NewHeatMap(grid, pal)
	p.Add(hm)

	// Removed InvertedScale: Frequency 0 should be at the bottom (Y=0)
	// p.Y.Scale = plot.InvertedScale{}

	return p.Save(10*vg.Inch, 4*vg.Inch, out)
}
