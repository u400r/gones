package modules

func UnsignedAdd[T BitSignal](a T, b T, c bool) (T, bool) {
	var carryIn T
	if c {
		carryIn = T(1)
	} else {
		carryIn = T(0)
	}
	result := a + b + carryIn
	carryOut := result < a
	return result, carryOut
}

func SignedAdd[T SignedBitSignal](a T, b T, c bool) (T, bool) {
	var carryIn T
	if c {
		carryIn = T(1)
	} else {
		carryIn = T(0)
	}

	result := a + b + carryIn
	overflow := false
	if (a > 0 && b > 0 && result < 0) || (a < 0 && b < 0 && result > 0) {
		overflow = true
	}
	return result, overflow
}

func UnsignedSub[T BitSignal](a T, b T, c bool) (T, bool) {
	var carryIn T
	if c {
		carryIn = T(0)
	} else {
		carryIn = T(1)
	}
	result := a - b - carryIn
	carryOut := result >= a
	return result, carryOut
}

func SignedSub[T SignedBitSignal](a T, b T, c bool) (T, bool) {
	var carryIn T
	if c {
		carryIn = T(0)
	} else {
		carryIn = T(1)
	}

	result := a - b - carryIn
	overflow := false
	if (a > 0 && b > 0 && result < 0) || (a < 0 && b < 0 && result > 0) {
		overflow = true
	}
	return result, overflow
}
