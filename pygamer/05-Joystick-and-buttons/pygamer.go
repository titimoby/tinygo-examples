package main

import (
	"machine"

	"tinygo.org/x/drivers/shifter"
)

// NewButtons returns a new shifter device for the buttons on an AdaFruit PyBadge
func NewButtons() *shifter.Device {
	var device = shifter.New(shifter.EIGHT_BITS, machine.BUTTON_LATCH, machine.BUTTON_CLK, machine.BUTTON_OUT)
	return &device
}

type joystick struct {
	x machine.ADC
	y machine.ADC
}

func NewJoystick() *joystick {
	return &joystick{x: machine.ADC{machine.JOYX}, y: machine.ADC{machine.JOYY}}
}

func (j *joystick) configure() {
	j.x.Configure(machine.ADCConfig{})
	j.y.Configure(machine.ADCConfig{})
}

func (j *joystick) getX() uint16 {
	return j.x.Get()
}

func (j *joystick) getY() uint16 {
	return j.y.Get()
}
