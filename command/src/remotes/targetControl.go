package remotes

import (
	"bytes"
	"fmt"

	"github.com/cosaques/patterns/command/src/commands"
)

const slotCount = 7

type targetControl struct {
	onCommands  []commands.Command
	offCommands []commands.Command
}

func (r *targetControl) SetCommand(slot int, on commands.Command, off commands.Command) {
	r.onCommands[slot] = on
	r.offCommands[slot] = off
}

func (r *targetControl) ButtonOnPressed(slot int) {
	r.onCommands[slot].Execute()
}

func (r *targetControl) ButtonOffPressed(slot int) {
	r.offCommands[slot].Execute()
}

func (r *targetControl) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n---- Remote control ----\n")
	for i := 0; i < len(r.onCommands); i++ {
		buffer.WriteString(fmt.Sprintf("[slot %d] On (%T) Off (%T)\n",
			i, r.onCommands[i], r.offCommands[i]))
	}

	return buffer.String()
}

func NewTargetControl() *targetControl {
	var r = new(targetControl)
	r.onCommands, r.offCommands =
		make([]commands.Command, slotCount), make([]commands.Command, slotCount)

	for i := 0; i < slotCount; i++ {
		r.onCommands[i] = new(commands.Empty)
		r.offCommands[i] = new(commands.Empty)
	}

	return r
}
