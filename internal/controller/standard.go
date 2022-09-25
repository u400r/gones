package controller

type StandardController struct {
	buttonStatus [8]bool
	currentBit   uint8
	isStrobe     bool
}

const (
	A      = 0
	B      = 1
	START  = 2
	SELECT = 3
	UP     = 4
	DOWN   = 5
	LEFT   = 6
	RIGHT  = 7
)

func NewController() *StandardController {
	return &StandardController{
		buttonStatus: [8]bool{false, false, false, false, false, false, false, false},
		currentBit:   uint8(0),
		isStrobe:     true,
	}
}

func (s *StandardController) Toggle(button int, isPressed bool) {
	s.buttonStatus[button] = isPressed
}

func (s *StandardController) Get() uint8 {
	var bit bool
	if s.isStrobe {
		bit = s.buttonStatus[0]
		s.currentBit = 0
	} else {
		bit = s.buttonStatus[s.currentBit]
		s.currentBit = (s.currentBit + 1) & 0x7
	}
	if bit {
		return 1
	} else {
		return 0
	}
}

func (s *StandardController) ChangeStrobe(value bool) {
	s.isStrobe = value
}
