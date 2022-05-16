package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoverPositionNoInstruction(t *testing.T) {
	rover := NewRover(1, 1, "N")
	assert.Equal(t, &Rover{1, 1, "N"}, rover)
}

func TestRoverPositionMoveForwardOneStepNorth(t *testing.T) {
	rover := NewRover(1, 1, "N")
	rover.Execute("F")
	assert.Equal(t, &Rover{1, 2, "N"}, rover)
}

func TestRoverPositionMoveForwardOneStepEast(t *testing.T) {
	rover := NewRover(1, 1, "E")
	rover.Execute("F")
	assert.Equal(t, &Rover{2, 1, "E"}, rover)
}

func TestRoverPositionMoveForwardOneStepSouth(t *testing.T) {
	rover := NewRover(1, 1, "S")
	rover.Execute("F")
	assert.Equal(t, &Rover{1, 0, "S"}, rover)
}

func TestRoverPositionMoveForwardOneStepWest(t *testing.T) {
	rover := NewRover(1, 1, "W")
	rover.Execute("F")
	assert.Equal(t, &Rover{0, 1, "W"}, rover)
}

func TestRoverPositionMoveBackwardOneStepNorth(t *testing.T) {
	rover := NewRover(1, 1, "N")
	rover.Execute("B")
	assert.Equal(t, &Rover{1, 0, "N"}, rover)
}

func TestRoverPositionMoveBackwardOneStepEast(t *testing.T) {
	rover := NewRover(1, 1, "E")
	rover.Execute("B")
	assert.Equal(t, &Rover{0, 1, "E"}, rover)
}

func TestRoverPositionMoveBackwardOneStepSouth(t *testing.T) {
	rover := NewRover(1, 1, "S")
	rover.Execute("B")
	assert.Equal(t, &Rover{1, 2, "S"}, rover)
}

func TestRoverPositionMoveBackwardsOneStepWest(t *testing.T) {
	rover := NewRover(1, 1, "W")
	rover.Execute("B")
	assert.Equal(t, &Rover{2, 1, "W"}, rover)
}

func TestRoveDirectionTurnRightFacingNorth(t *testing.T) {
	rover := NewRover(1, 1, "N")
	rover.Execute("R")
	assert.Equal(t, &Rover{1, 1, "E"}, rover)
}

func TestRoverDirectionTurnRightFacingEast(t *testing.T) {
	rover := NewRover(1, 1, "E")
	rover.Execute("R")
	assert.Equal(t, &Rover{1, 1, "S"}, rover)
}

func TestRoverDirectionTurnRightFacingSouth(t *testing.T) {
	rover := NewRover(1, 1, "S")
	rover.Execute("R")
	assert.Equal(t, &Rover{1, 1, "W"}, rover)
}

func TestRoverDirectionTurnRightFacingWest(t *testing.T) {
	rover := NewRover(1, 1, "W")
	rover.Execute("R")
	assert.Equal(t, &Rover{1, 1, "N"}, rover)
}

func TestRoverDirectionTurnLeftFacingNorth(t *testing.T) {
	rover := NewRover(1, 1, "N")
	rover.Execute("L")
	assert.Equal(t, &Rover{1, 1, "W"}, rover)
}

func TestRoverDirectionTurnLeftFacingEast(t *testing.T) {
	rover := NewRover(1, 1, "E")
	rover.Execute("L")
	assert.Equal(t, &Rover{1, 1, "N"}, rover)
}

func TestRoverDirectionTurnLeftFacingSouth(t *testing.T) {
	rover := NewRover(1, 1, "S")
	rover.Execute("L")
	assert.Equal(t, &Rover{1, 1, "E"}, rover)
}

func TestRoverDirectionTurnLeftFacingWest(t *testing.T) {
	rover := NewRover(1, 1, "W")
	rover.Execute("L")
	assert.Equal(t, &Rover{1, 1, "S"}, rover)
}
