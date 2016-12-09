package survey

import (
	"errors"
	"github.com/pkg/term"
)

// key codes for the common keys
const (
	KeyArrowUp = iota
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyEsc
	KeyEnter
)

// GetChar listens for input from the keyboard and returns the key value as a string
// or one of the Key* enum values.
func GetChar() (val string, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".
		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = KeyArrowUp
		} else if bytes[2] == 66 {
			// Down
			keyCode = KeyArrowDown
		} else if bytes[2] == 67 {
			// Right
			keyCode = KeyArrowRight
		} else if bytes[2] == 68 {
			// Left
			keyCode = KeyArrowLeft
		}
	} else if numRead == 1 {
		ascii := int(bytes[0])

		// if the user sends SIGTERM (ascii 3) or presses esc (ascii 27)
		if ascii == 3 || ascii == 27 {
			// hard close
			err = errors.New("Goodbye.")
		}

		// handle the enter key
		if ascii == 13 {
			keyCode = KeyEnter
		}

		val = string(ascii)
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}
