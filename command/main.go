package main

import (
	"fmt"

	"github.com/cosaques/patterns/command/src/commands"
	"github.com/cosaques/patterns/command/src/remotes"
)

func main() {
	testTargetControl()
}

func testTargetControl() {
	var rc = remotes.NewTargetControl()
	fmt.Println(rc)

	var lightOn commands.LightOn
	var lightOff commands.LightOff
	rc.SetCommand(0, &lightOn, &lightOff)

	var garageDoorOpen commands.GarageDoorOpen
	var garageDoorClose commands.GarageDoorClose
	rc.SetCommand(1, &garageDoorOpen, &garageDoorClose)
	fmt.Println(rc)

	rc.ButtonOnPressed(0)
	rc.ButtonOnPressed(1)
	rc.ButtonOffPressed(0)
	rc.ButtonOffPressed(1)
}

func testSimpleControl() {
	var rc = new(remotes.SimpleControl)

	var lightOn commands.LightOn
	rc.SetCommand(&lightOn)
	rc.ButtonPressed()

	var garageDoorOpen commands.GarageDoorOpen
	rc.SetCommand(&garageDoorOpen)
	rc.ButtonPressed()
}
