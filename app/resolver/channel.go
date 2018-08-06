package resolver

import (
	"fmt"

	"github.com/dalloriam/synthia/core"
)

type ChannelResolver struct {
	channel *core.MixerChannel
}

func (r *ChannelResolver) Volume() *KnobResolver {
	return &KnobResolver{r.channel.Volume}
}

func (r *ChannelResolver) Pan() *KnobResolver {
	return &KnobResolver{r.channel.Pan}
}

func (r *ChannelResolver) Input() *string {
	if r.channel.Input == nil {
		return nil
	}

	outStr := fmt.Sprintf("%T", r.channel.Input)
	return &outStr
}
