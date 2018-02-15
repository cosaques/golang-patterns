package main

import (
	"github.com/cosaques/patterns/command/src/commands"
	"github.com/cosaques/patterns/command/src/remotes"
)

func main() {
	var rc = new(remotes.SimpleControl)

	var lightOn = new(commands.LightOn)
	rc.SetCommand(lightOn)
	rc.ButtonPressed()

	var garageDoor = new(commands.GarageDoor)
	rc.SetCommand(garageDoor)
	rc.ButtonPressed()
}
