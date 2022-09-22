package modules

func UnsignedAdd[T ByteSignal](a T, b T, c bool) (T, bool) {
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

func SignedAdd[T SignedByteSignal](a T, b T, c bool) (T, bool) {
	var carryIn T
	if c {
		carryIn = T(1)
	} else {
		carryIn = T(0)
	}

	result := a + b + carryIn
	result64 := int64(a) + int64(b) + int64(carryIn)
	overflow := false
	if int64(result) != result64 {
		overflow = true
	}
	return result, overflow
}

func UnsignedSub[T ByteSignal](a T, b T, c bool) (T, bool) {
	var carryIn T
	if c {
		carryIn = T(0)
	} else {
		carryIn = T(1)
	}
	result := a - b - carryIn
	carryOut := result <= a
	return result, carryOut
}

func SignedSub[T SignedByteSignal](a T, b T, c bool) (T, bool) {
	var carryIn T
	if c {
		carryIn = T(0)
	} else {
		carryIn = T(1)
	}

	result := a - b - carryIn
	result64 := int64(a) - int64(b) - int64(carryIn)
	overflow := false
	if int64(result) != result64 {
		overflow = true
	}
	return result, overflow
}
