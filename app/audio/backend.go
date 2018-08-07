package audio

import (
	"github.com/gordonklaus/portaudio"
)

const (
	audioChannelCount = 2 // Stereo.
)

type portaudioBackend struct {
	params portaudio.StreamParameters
}

func newBackend(sampleRate, bufferSize, inputDevice, outputDevice int) (*portaudioBackend, error) {
	// Quick-and-dirty way to initialize portaudio
	if err := portaudio.Initialize(); err != nil {
		panic(err)
	}

	devices, err := portaudio.Devices()
	if err != nil {
		return nil, err
	}

	// TODO: Detect devices properly
	inDevice, outDevice := devices[inputDevice], devices[outputDevice]

	params := portaudio.LowLatencyParameters(inDevice, outDevice)

	params.Input.Channels = 1
	params.Output.Channels = audioChannelCount

	params.SampleRate = float64(sampleRate)
	params.FramesPerBuffer = bufferSize

	return &portaudioBackend{params}, err
}

func (b *portaudioBackend) Start(callback func(in []float32, out [][]float32)) error {
	stream, err := portaudio.OpenStream(b.params, callback)
	if err != nil {
		return err
	}

	return stream.Start()
}

func (b *portaudioBackend) FrameSize() int {
	return b.params.FramesPerBuffer
}
