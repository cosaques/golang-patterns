package remotes

import "github.com/cosaques/patterns/command/commands"

type SimpleControl struct {
	slot commands.Command
}

func (rc *SimpleControl) SetCommand(c commands.Command) {
	rc.slot = c
}

func (rc *SimpleControl) ButtonPressed() {
	if rc.slot != nil {
		rc.slot.Execute()
	}
}
