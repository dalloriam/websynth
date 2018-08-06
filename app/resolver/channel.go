package resolver

import (
	"fmt"

	"github.com/dalloriam/synthia/core"
)

type ChannelResolver struct {
	channel *core.MixerChannel
}

func (r *ChannelResolver) Volume() float64 {
	return r.channel.Volume.GetValue()
}

func (r *ChannelResolver) Pan() float64 {
	return r.channel.Pan.GetValue()
}

func (r *ChannelResolver) Input() *string {
	if r.channel.Input == nil {
		return nil
	}

	outStr := fmt.Sprintf("%T", r.channel.Input)
	return &outStr
}

func (r *ChannelResolver) SetPan(args struct{ Pan float64 }) *ChannelResolver {
	r.channel.Pan.SetValue(args.Pan)
	return r
}

func (r *ChannelResolver) SetVolume(args struct{ Volume float64 }) *ChannelResolver {
	r.channel.Volume.SetValue(args.Volume)
	return r
}
