package resolver

import (
	"github.com/dalloriam/websynth/app/audio"

	"github.com/dalloriam/synthia/core"
)

type ChannelResolver struct {
	sys     *audio.System
	channel *core.MixerChannel
}

func (r *ChannelResolver) Volume() *KnobResolver {
	return &KnobResolver{r.sys, r.channel.Volume}
}

func (r *ChannelResolver) Pan() *KnobResolver {
	return &KnobResolver{r.sys, r.channel.Pan}
}
