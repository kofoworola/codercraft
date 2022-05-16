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

func (r *Rover) facingSouth(instruction string) {
	if instruction == "F" {
		r.y--
	} else if instruction == "B" {
		r.y++
	} else if instruction == "R" {
		r.facing = "W"
	} else if instruction == "L" {
		r.facing = "E"
	}

}

func (r *Rover) facingNorth(instruction string) {
	if instruction == "F" {
		r.y++
	} else if instruction == "B" {
		r.y--
	} else if instruction == "R" {
		r.facing = "E"
	} else if instruction == "L" {
		r.facing = "W"
	}

}

func (r *Rover) Execute(instructionList string) {
	switch r.facing {
	case "E":
		r.facingEast(instructionList)
	case "S":
		r.facingSouth(instructionList)
	case "N":
		r.facingNorth(instructionList)
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
