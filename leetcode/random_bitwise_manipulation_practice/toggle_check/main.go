package main

import "fmt"

type SmartHome struct {
	status uint8
}

func newSmartHome() *SmartHome {
	return &SmartHome{}
}

func (sh *SmartHome) check(item uint8) bool {
	return sh.status&item != 0
}

func (sh *SmartHome) toggle(item uint8) bool {
	sh.status ^= item

	return sh.check(item)
}

func (sh *SmartHome) turnOff(item uint8) bool {
	sh.status = sh.status & ^item

	return sh.check(item)
}

func (sh *SmartHome) turnOn(item uint8) bool {
	sh.status = sh.status | item

	return sh.check(item)
}

func (sh *SmartHome) turnOffAll() {
	sh.status = 0
}

// SmartHome items:
const (
	LIGHT     uint8 = 0b00000001
	AC        uint8 = 0b00000010
	TV        uint8 = 0b00000100
	DOOR_LOCK uint8 = 0b00001000
	KETTLE    uint8 = 0b00010000
	HEATING   uint8 = 0b00100000
	ALARM     uint8 = 0b01000000
	VACUUM    uint8 = 0b10000000
)

func main() {
	sh := newSmartHome()

	fmt.Printf("TV is on: %t\n", sh.check(TV))
	sh.turnOn(TV)
	fmt.Printf("TV is on: %t\n", sh.check(TV))

	fmt.Printf("Kettle is on: %t\n", sh.check(KETTLE))
	sh.turnOn(KETTLE)
	fmt.Printf("Kettle is on: %t\n", sh.check(KETTLE))

	sh.turnOffAll()

	fmt.Printf("TV is on: %t\n", sh.check(TV))
	fmt.Printf("Kettle is on: %t\n", sh.check(KETTLE))
}
