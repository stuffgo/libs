package types

type (
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}

	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	Integer interface {
		Signed | Unsigned
	}

	Float interface {
		~float32 | ~float64
	}

	Complex interface {
		~complex64 | ~complex128
	}

	Ordered interface {
		Integer | Float | ~string
	}
)
