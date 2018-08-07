package resolver

import (
	"github.com/dalloriam/synthia/modular"
	"github.com/dalloriam/websynth/app/audio"
)

type ClockResolver struct {
	sys   *audio.System
	clock *modular.Clock
}

func (c *ClockResolver) Tempo() *KnobResolver {
	return &KnobResolver{c.sys, c.clock.Tempo}
}
