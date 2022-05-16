package rover

type Rover struct {
	x      int
	y      int
	facing string
}

func NewRover(x, y int, direction string) *Rover {
	return &Rover{x, y, direction}
}

func (r *Rover) facingEast(instruction rune) {
	if instruction == 'F' {
		r.x++
	} else if instruction == 'B' {
		r.x--
	} else if instruction == 'R' {
		r.facing = "S"
	} else if instruction == 'L' {
		r.facing = "N"
	}
}

func (r *Rover) facingSouth(instruction rune) {
	if instruction == 'F' {
		r.y--
	} else if instruction == 'B' {
		r.y++
	} else if instruction == 'R' {
		r.facing = "W"
	} else if instruction == 'L' {
		r.facing = "E"
	}
}

func (r *Rover) facingNorth(instruction rune) {
	if instruction == 'F' {
		r.y++
	} else if instruction == 'B' {
		r.y--
	} else if instruction == 'R' {
		r.facing = "E"
	} else if instruction == 'L' {
		r.facing = "W"
	}
}

func (r *Rover) facingWest(instruction rune) {
	if instruction == 'F' {
		r.x--
	} else if instruction == 'B' {
		r.x++
	} else if instruction == 'R' {
		r.facing = "N"
	} else if instruction == 'L' {
		r.facing = "S"
	}

}

func (r *Rover) Execute(instructionList string) {
	for _, instruction := range instructionList {
		switch r.facing {
		case "E":
			r.facingEast(instruction)
		case "S":
			r.facingSouth(instruction)
		case "N":
			r.facingNorth(instruction)
		case "W":
			r.facingWest(instruction)
		}

	}
}
