package app

import (
	"github.com/dalloriam/synthia"
	"github.com/dalloriam/synthia/modular"
)

const (
	bufferSize        = 512   // Size of the audio buffer.
	sampleRate        = 44100 // Audio sample rate.
	audioChannelCount = 2     // Stereo.
	mixerChannelCount = 2     // Three oscillators.
)

func getSynth() (*synthia.Synthia, error) {

	backend, err := NewBackend()
	if err != nil {
		return nil, err
	}

	// Create an oscillator and attach the sequencer to it.
	osc1 := modular.NewOscillator()

	// Create the synthesizer with two mixer channels and set it to output to our audio backend.
	synth := synthia.New(mixerChannelCount, bufferSize, backend)

	// Map two different waves to the two outputs of our mixer.
	synth.Mixer.Channels[0].Input = osc1.Sine
	synth.Mixer.Channels[1].Input = osc1.Triangle

	// Pan the sine wave to the right channel and the triangle wave to the left channel
	synth.Mixer.Channels[0].Pan.SetValue(-1)
	synth.Mixer.Channels[1].Pan.SetValue(1)

	return synth, nil
}
