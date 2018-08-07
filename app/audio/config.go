package audio

// Config wraps audio configuration
type Config struct {
	BufferSize     int
	SampleRate     int
	MixerChanCount int

	InputDevice  int
	OutputDevice int
}
