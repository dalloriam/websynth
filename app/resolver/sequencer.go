package resolver

import (
	"github.com/dalloriam/synthia/modular"
	"github.com/dalloriam/websynth/app/audio"
)

type SequencerResolver struct {
	sys       *audio.System
	sequencer *modular.Sequencer
}

func (r *SequencerResolver) Sequence() []float64 {
	return r.sequencer.Sequence
}

func (r *SequencerResolver) SetSequence(args struct{ Seq *[]float64 }) bool {
	if args.Seq == nil {
		return false
	}

	r.sequencer.Sequence = *args.Seq
	return true
}

func (r *SequencerResolver) BeatsPerStep() *KnobResolver {
	return &KnobResolver{r.sys, r.sequencer.BeatsPerStep}
}

func (r *SequencerResolver) Clock() *SignalResolver {
	return &SignalResolver{r.sys, &r.sequencer.Clock}
}
