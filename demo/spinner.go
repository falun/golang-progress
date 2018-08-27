package main

import "github.com/falun/golang-progress"
import "time"

func main() {
	t := time.NewTimer(4 * time.Second)
	p := progress.NewSpinnerWithFrames(progress.Animations["weather"])
	p.Start("Looking out the window ")
	<-t.C
	p.Stopln()

	t.Reset(4 * time.Second)
	p.ChangeFrames(progress.Animations["hearts"])
	p.Start("Feeling ")
	<-t.C
	p.Stopln()
}
