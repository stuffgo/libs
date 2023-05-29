package types

type (
	Item[T any] interface {
		Value() T
		Less(r Item[T]) bool
		Equal(r Item[T]) bool
	}

	BTS[T Item[T]] interface {
		Get() Item[T]
		GetMin() Item[T]
		GetMax() Item[T]
		Insert(val Item[T])
		Delete(val Item[T])
		Find(val Item[T]) BTS[T]
		Len() int
		Iter() <-chan BTS[T]
	}
)
