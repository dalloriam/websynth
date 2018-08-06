package resolver

import (
	"fmt"

	"github.com/dalloriam/websynth/app/audio"

	"github.com/dalloriam/synthia/core"
)

type MixerResolver struct {
	sys   *audio.System
	mixer *core.Mixer
}

func (m *MixerResolver) Channel(args struct{ Idx int32 }) (*ChannelResolver, error) {
	i := int(args.Idx)

	if i >= len(m.mixer.Channels) {
		return nil, fmt.Errorf("cannot return channel %d, mixer has only %d channels", args.Idx, len(m.mixer.Channels))
	}

	return &ChannelResolver{m.sys, m.mixer.Channels[i]}, nil
}

func (m *MixerResolver) Channels() []*ChannelResolver {
	resolvers := make([]*ChannelResolver, len(m.mixer.Channels))

	for i := 0; i < len(m.mixer.Channels); i++ {
		resolvers[i] = &ChannelResolver{m.sys, m.mixer.Channels[i]}
	}
	return resolvers
}
