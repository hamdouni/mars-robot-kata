package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobotString(t *testing.T) {
	testCases := []struct {
		name        string
		robotName   string
		expectedMsg string
	}{
		{
			name:        "happy path",
			robotName:   "Marie",
			expectedMsg: "Hello, Marie",
		},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			r := Robot{Name: c.robotName}
			assert.Equal(t, c.expectedMsg, r.Hello())
		})
	}
}

func TestMove(t *testing.T) {
	cases := []struct {
		name              string
		rob               Robot
		move              rune
		expectedDirection rune
		expectedX         int
		expectedY         int
	}{
		{
			name:              "Forward to South",
			rob:               Robot{Name: "Rob", Direction: 'S', X: 0, Y: 0},
			move:              'F',
			expectedDirection: 'S',
			expectedX:         0,
			expectedY:         1,
		},
		{
			name:              "Forward to East",
			rob:               Robot{Name: "Rob", Direction: 'E', X: 0, Y: 0},
			move:              'F',
			expectedDirection: 'E',
			expectedX:         1,
			expectedY:         0,
		},
		{
			name:              "Forward to North",
			rob:               Robot{Name: "Rob", Direction: 'N', X: 0, Y: 1},
			move:              'F',
			expectedDirection: 'N',
			expectedX:         0,
			expectedY:         0,
		},
		{
			name:              "Forward to West",
			rob:               Robot{Name: "Rob", Direction: 'W', X: 1, Y: 0},
			move:              'F',
			expectedDirection: 'W',
			expectedX:         0,
			expectedY:         0,
		},
		{
			name:              "Backward from South",
			rob:               Robot{Name: "Rob", Direction: 'S', X: 0, Y: 1},
			move:              'B',
			expectedDirection: 'S',
			expectedX:         0,
			expectedY:         0,
		},
		{
			name:              "Backward from East",
			rob:               Robot{Name: "Rob", Direction: 'E', X: 1, Y: 0},
			move:              'B',
			expectedDirection: 'E',
			expectedX:         0,
			expectedY:         0,
		},
		{
			name:              "Backward from North",
			rob:               Robot{Name: "Rob", Direction: 'N', X: 0, Y: 0},
			move:              'B',
			expectedDirection: 'N',
			expectedX:         0,
			expectedY:         1,
		},
		{
			name:              "Backward from West",
			rob:               Robot{Name: "Rob", Direction: 'W', X: 0, Y: 0},
			move:              'B',
			expectedDirection: 'W',
			expectedX:         1,
			expectedY:         0,
		},
	}

	for _, uc := range cases {
		t.Run(uc.name, func(t *testing.T) {
			uc.rob.Move(uc.move)
			if uc.rob.X != uc.expectedX || uc.rob.Y != uc.expectedY || uc.rob.Direction != uc.expectedDirection {
				t.Errorf("Expected robot %v,%v,%v but got %v,%v,%v", uc.expectedX, uc.expectedY, uc.expectedDirection, uc.rob.X, uc.rob.Y, uc.rob.Direction)
			}
		})
	}
}
