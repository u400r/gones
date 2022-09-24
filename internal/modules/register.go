package modules

import "sync"

type ByteSignal interface {
	uint8 | uint16 | uint32 | uint64
}

type SignedByteSignal interface {
	int8 | int16 | int32 | int64
}

type ReadableRegister[T ByteSignal] interface {
	Read() T
}

type WritableRegister[T ByteSignal] interface {
	Write(data T)
	ReadableRegister[T]
}

type Counter[T ByteSignal] interface {
	Increment()
	Increment32()
	Decrement()
	WritableRegister[T]
}

type Flag[T ByteSignal] interface {
	WritableRegister[T]
	Get(nbit uint) bool
	Set(nbit uint)
	Clear(nbit uint)
	Change(nbit uint, flag bool)
}

type ShiftableRegister[T ByteSignal] interface {
	WritableRegister[T]
	Left(carry bool)
	Load(data T)
}

type Register[T ByteSignal] struct {
	data T
	m    sync.RWMutex
}

func NewRegister[T ByteSignal](data T) *Register[T] {
	return &Register[T]{
		data: data,
		m:    sync.RWMutex{},
	}
}

func (r *Register[T]) Read() T {
	r.m.RLock()
	defer r.m.RUnlock()
	return r.data
}

func (r *Register[T]) Write(data T) {
	r.m.Lock()
	defer r.m.Unlock()
	r.data = data
}

func (r *Register[T]) Increment() {
	r.data += 1
}

func (r *Register[T]) Increment32() {
	r.data += 32
}

func (r *Register[T]) Decrement() {
	r.data -= 1
}

func (r *Register[T]) Get(nbit uint) bool {
	return (r.data>>nbit)&0x1 == 0x1
}

func (r *Register[T]) Set(nbit uint) {
	r.data = r.data | 1<<nbit
}

func (r *Register[T]) Clear(nbit uint) {
	r.data = r.data & ^(1 << nbit)
}

func (r *Register[T]) Change(nbit uint, flag bool) {
	if flag {
		r.Set(nbit)
	} else {
		r.Clear(nbit)
	}
}

func (r *Register[T]) Left(carry bool) {
	var carryIn T
	if carry {
		carryIn = 1
	} else {
		carryIn = 0
	}
	r.data = r.data<<1 + carryIn
}

func (r *Register[T]) Load(data T) {
	r.data = r.data | data
}

// cpu status bit
const (
	CARRY     uint = 0
	ZERO      uint = 1
	INTERRUPT uint = 2
	DECIMAL   uint = 3
	BLEAK     uint = 4
	OVERFLOW  uint = 6
	NEGATIVE  uint = 7
)

// ppu ctrl bit
const (
	NametableSelectLow  = 0
	NametableSelectHigh = 1
	IncrementMode       = 2
	SpritetileSelect    = 3
	BgtileSelect        = 4
	SpriteHight         = 5
	PpuMode             = 6
	NmiEnable           = 7
)

// ppu mask bit
const (
	Greyscale            = 0
	BgLeftmostEnable     = 1
	SpriteLeftmostEnable = 2
	BgEnable             = 3
	SpriteEnable         = 4
	RedEmphasis          = 5
	GreenEmphasis        = 6
	BlueEmphasis         = 7
)

// ppu status bit

const (
	SpriteOveflow = 5
	SpriteZeroHit = 6
	Vblank        = 7
)
