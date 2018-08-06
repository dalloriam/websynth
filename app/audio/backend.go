package audio

import (
	"github.com/gordonklaus/portaudio"
)

type PortaudioBackend struct {
	params portaudio.StreamParameters
}

func NewBackend() (*PortaudioBackend, error) {
	// Quick-and-dirty way to initialize portaudio
	if err := portaudio.Initialize(); err != nil {
		panic(err)
	}

	devices, err := portaudio.Devices()
	if err != nil {
		return nil, err
	}
	inDevice, outDevice := devices[1], devices[2]

	params := portaudio.LowLatencyParameters(inDevice, outDevice)

	params.Input.Channels = 1
	params.Output.Channels = audioChannelCount

	params.SampleRate = float64(sampleRate)
	params.FramesPerBuffer = bufferSize

	return &PortaudioBackend{params}, err
}

func (b *PortaudioBackend) Start(callback func(in []float32, out [][]float32)) error {
	stream, err := portaudio.OpenStream(b.params, callback)
	if err != nil {
		return err
	}

	return stream.Start()
}

func (b *PortaudioBackend) FrameSize() int {
	return b.params.FramesPerBuffer
}
