package rover

type Rover struct {
	x      int
	y      int
	facing string
}

func NewRover(x, y int, direction string) *Rover {
	return &Rover{x, y, direction}
}

func (r *Rover) facingEast(instruction string) {
	if instruction == "F" {
		r.x++
	} else if instruction == "B" {
		r.x--
	} else if instruction == "R" {
		r.facing = "S"
	} else if instruction == "L" {
		r.facing = "N"
	}

}

func (r *Rover) Execute(instructionList string) {
	switch r.facing {
	case "E":
		r.facingEast(instructionList)
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
