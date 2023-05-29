package lsm

import (
	"sync"
)

type lsmTree struct {
	// memory cache RW mutex
	mrwLock sync.RWMutex

	// disk storage RW mutex
	drwLock sync.RWMutex
}