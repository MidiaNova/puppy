package puppy

import (
	"github.com/MidiaNova/dog"
)

func Bark() string {
	return "hup!"
}

func Barks() string {
	return "hup! hup! hup!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}
