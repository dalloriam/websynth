package audio

import (
	"github.com/dalloriam/synthia"
)

const (
	bufferSize        = 512   // Size of the audio buffer.
	sampleRate        = 44100 // Audio sample rate.
	audioChannelCount = 2     // Stereo.
	mixerChannelCount = 2     // Two oscillators.
)

type System struct {
	Synth *synthia.Synthia

	Modules []interface{}

	ModuleChannelMap map[int]int
}

func NewSystem() (*System, error) {
	backend, err := NewBackend()
	if err != nil {
		return nil, err
	}

	synth := getSynth(backend)

	return &System{synth, []interface{}{}, make(map[int]int)}, nil
}

func getSynth(backend *PortaudioBackend) *synthia.Synthia {

	// Create the synthesizer with two mixer channels and set it to output to our audio backend.
	synth := synthia.New(mixerChannelCount, bufferSize, backend)

	return synth
}
