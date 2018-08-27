package progress

import (
	"time"
)

type Frameset interface {
	Name() string
	Frames() []string
	TickRate() time.Duration
}

func NewFrameset(name string, frames []string, rateMS int) Frameset {
	return frameset{
		name,
		frames,
		time.Duration(rateMS) * time.Millisecond,
	}
}

type frameset struct {
	name   string
	frames []string
	rate   time.Duration
}

func (fs frameset) Name() string            { return fs.name }
func (fs frameset) Frames() []string        { return fs.frames }
func (fs frameset) TickRate() time.Duration { return fs.rate }
