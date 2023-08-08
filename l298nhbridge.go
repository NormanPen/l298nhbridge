package l298nhbridge

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

// DC_MAX ist die maximale Drehgeschwindigkeit der Motoren
const DC_MAX = 70

var (
	leftmotorIn1Pin  = rpio.Pin(6)
	leftmotorIn2Pin  = rpio.Pin(13)
	rightmotorIn1Pin = rpio.Pin(19)
	rightmotorIn2Pin = rpio.Pin(26)
	leftmotorPWMPin  = rpio.Pin(18)
	rightmotorPWMPin = rpio.Pin(12)
)

func init() {
	if err := rpio.Open(); err != nil {
		fmt.Println("Fehler beim Öffnen der GPIO-Pins:", err)
		return
	}
	leftmotorIn1Pin.Output()
	leftmotorIn2Pin.Output()
	rightmotorIn1Pin.Output()
	rightmotorIn2Pin.Output()
	leftmotorPWMPin.Output()
	rightmotorPWMPin.Output()

	leftmotorPWMPin.Pwm()
	leftmotorPWMPin.Freq(100) // PWM-Frequenz auf 100 Hz setzen
	leftmotorPWMPin.DutyCycle(0, DC_MAX)

	rightmotorPWMPin.Pwm()
	rightmotorPWMPin.Freq(100) // PWM-Frequenz auf 100 Hz setzen
	rightmotorPWMPin.DutyCycle(0, DC_MAX)

	setMotors(0, 0)
}

// SetMotors steuert die Geschwindigkeit der linken und rechten Motoren
// powerLeft und powerRight liegen zwischen -1 und 1
func setMotors(powerLeft, powerRight float32) {
	pwmLeft := int32(powerLeft * DC_MAX)
	pwmRight := int32(powerRight * DC_MAX)

	leftmotorIn1Pin.Low()
	leftmotorIn2Pin.Low()
	rightmotorIn1Pin.Low()
	rightmotorIn2Pin.Low()

	if pwmLeft < 0 {
		leftmotorIn2Pin.High()
	} else if pwmLeft > 0 {
		leftmotorIn1Pin.High()
	}
	leftmotorPWMPin.DutyCycle(0, uint32(abs(pwmLeft)))

	if pwmRight < 0 {
		rightmotorIn2Pin.High()
	} else if pwmRight > 0 {
		rightmotorIn1Pin.High()
	}
	rightmotorPWMPin.DutyCycle(0, uint32(abs(pwmRight)))
}

// StopMotors stoppt beide Motoren
func stopMotors() {
	leftmotorIn1Pin.Low()
	leftmotorIn2Pin.Low()
	rightmotorIn1Pin.Low()
	rightmotorIn2Pin.Low()
	leftmotorPWMPin.DutyCycle(0, 0)
	rightmotorPWMPin.DutyCycle(0, 0)
}

// Exit beendet die GPIO-Ressourcen
func exit() {
	stopMotors()
	rpio.Close()
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}
