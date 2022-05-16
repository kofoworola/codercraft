package rover

type Rover struct {
	x      int
	y      int
	facing string
}

func NewRover(x, y int, direction string) *Rover {
	return &Rover{x, y, direction}
}

func (r *Rover) Execute(instructionList string) {
	switch r.facing {
	case "E":
		if instructionList == "F" {
			r.x++
		} else if instructionList == "B" {
			r.x--
		} else if instructionList == "R" {
			r.facing = "S"
		} else if instructionList == "L" {
			r.facing = "N"
		}
	case "S":
		if instructionList == "F" {
			r.y--
		} else if instructionList == "B" {
			r.y++
		} else if instructionList == "R" {
			r.facing = "W"
		} else if instructionList == "L" {
			r.facing = "E"
		}
	case "N":
		if instructionList == "F" {
			r.y++
		} else if instructionList == "B" {
			r.y--
		} else if instructionList == "R" {
			r.facing = "E"
		} else if instructionList == "L" {
			r.facing = "W"
		}
	case "W":
		if instructionList == "F" {
			r.x--
		} else if instructionList == "B" {
			r.x++
		} else if instructionList == "R" {
			r.facing = "N"
		} else if instructionList == "L" {
			r.facing = "S"
		}
	}
}
