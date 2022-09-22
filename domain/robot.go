package domain

import (
	"fmt"
)

type Robot struct {
	Name      string
	Direction rune
	X, Y      int
}

func (r Robot) Hello() string {
	return fmt.Sprintf("Hello, %s", r.Name)
}

func (r *Robot) Move(cmd rune) error {
	switch cmd {
	case 'F':
		switch r.Direction {
		case 'S':
			r.Y++
		case 'E':
			r.X++
		case 'N':
			r.Y--
		case 'W':
			r.X--
		}
	case 'B':
		switch r.Direction {
		case 'S':
			r.Y--
		case 'E':
			r.X--
		case 'N':
			r.Y++
		case 'W':
			r.X++
		}
	default:
		return fmt.Errorf("Command %v not recognised", cmd)
	}
	return nil
}
