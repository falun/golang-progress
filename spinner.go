package progress

import (
	"fmt"
	"time"
)

type Spinner interface {
	Start(s string)
	Stop()
	Stopln()
	Text(s string)
	ChangeFrames(fs Frameset)
}

type spinner struct {
	tickRate    time.Duration
	lastText    string
	currentText string
	frames      []string
	frame       int
	active      bool
	timer       *time.Timer
}

var defaultFrameset = "line"

func NewSpinner() Spinner {
	return NewSpinnerWithFrames(
		Animations[defaultFrameset],
	)
}

func NewSpinnerWithFrames(frameset Frameset) Spinner {
	return &spinner{
		tickRate:    frameset.TickRate(),
		currentText: "",
		frames:      frameset.Frames(),
		frame:       0,
		active:      false,
	}
}

func (p *spinner) ChangeFrames(fs Frameset) {
	p.tickRate = fs.TickRate()
	p.frames = fs.Frames()
	p.frame = 0
}

func (p *spinner) draw() {
	reset := "\033[2K\r"
	fmt.Printf(reset + "%s %s", p.frames[p.frame], p.currentText)
	p.frame = (p.frame + 1) % len(p.frames)
}

func (p *spinner) Text(s string) {
	p.currentText = s
}

func (p *spinner) Start(s string) {
	p.Text(s)

	if p.active {
		return
	}

	if p.timer == nil {
		p.timer = time.NewTimer(p.tickRate)
	} else {
		p.timer.Reset(p.tickRate)
	}
	p.active = true

	go func() {
		for p.active {
			<-p.timer.C
			p.draw()
			p.timer.Reset(p.tickRate)
		}
	}()
}

func (p *spinner) Stop() {
	if p.active {
		p.active = false
		if !p.timer.Stop() {
			<-p.timer.C
		}
	}
}

func (p *spinner) Stopln() {
	p.Stop()
	fmt.Println()
}
