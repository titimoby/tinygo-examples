package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

type note struct {
	tone     float64
	duration float64
}

// to shorten code and add the triplet duration
const (
	half    = buzzer.Half
	quarter = buzzer.Quarter
	eigth   = buzzer.Eighth
	triplet = quarter / 3
)

func main() {
	bzrPin := machine.BUZZER
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bzr := buzzer.New(bzrPin)

	// because a geek wil always be a geek
	notes := []note{
		{buzzer.D3, triplet},
		{buzzer.D3, triplet},
		{buzzer.D3, triplet},
		{buzzer.G3, half},
		{buzzer.D4, half},
		{buzzer.C4, triplet},
		{buzzer.B3, triplet},
		{buzzer.A3, triplet},
		{buzzer.G4, half},
		{buzzer.D4, quarter},
		{buzzer.C4, triplet},
		{buzzer.B3, triplet},
		{buzzer.A3, triplet},
		{buzzer.G4, half},
		{buzzer.D4, quarter},
		{buzzer.C4, triplet},
		{buzzer.B3, triplet},
		{buzzer.C4, triplet},
		{buzzer.A3, half},
		{buzzer.Rest, buzzer.Eighth},
	}

	for _, n := range notes {
		bzr.Tone(n.tone, n.duration)
		time.Sleep(10 * time.Millisecond)
	}

	for {
		time.Sleep(time.Hour)
	}
}
