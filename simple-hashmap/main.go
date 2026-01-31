package main

import (
	"errors"
	"hash/maphash"
)

type Entry struct {
	key   string
	value string
}

type HashMap struct {
	list       []*Entry
	size       uint64
	occupied   uint64
	maxLoad    uint64
	loadFactor float64
	seed       maphash.Seed
}

func NewHashMap(initialSize uint64) (*HashMap, error) {
	if initialSize < 1 {
		return nil, errors.New("Invalid size.")
	}

	return &HashMap{
		list:     make([]*Entry, initialSize),
		size:     initialSize,
		occupied: 0,
		maxLoad:  initialSize * 3 / 4,
		seed:     maphash.MakeSeed(),
	}, nil
}

func (h *HashMap) hash(key string) uint64 {
	return maphash.String(h.seed, key) & (h.size - 1)
}

func (h *HashMap) rehash() {
	oldList := h.list
	h.size *= 2
	h.list = make([]*Entry, h.size)
	h.occupied = 0
	h.maxLoad = h.size * 3 / 4

	for _, entry := range oldList {
		if entry != nil {
			h.insertNoRehash(entry)
		}
	}
}

func (h *HashMap) insertNoRehash(e *Entry) {
	idx := h.hash(e.key)

	for probe := range h.size {
		i := (idx + probe) & (h.size - 1)

		if h.list[i] == nil {
			h.list[i] = e
			h.occupied++
			return
		}
	}

	panic("rehash insert failed")
}

func (h *HashMap) Put(key, value string) error {
	if len(key) == 0 {
		return errors.New("Invalid key")
	}

	if h.occupied >= h.maxLoad {
		h.rehash()
	}

	idx := h.hash(key)

	for probe := range h.size {
		i := (idx + probe) & (h.size - 1)

		entry := h.list[i]

		if entry == nil {
			h.list[i] = &Entry{key, value}
			h.occupied++
			return nil
		}

		if entry.key == key {
			entry.value = value
			return nil
		}
	}

	panic("HashMap is full")
}

func (h *HashMap) Get(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("Invalid key")
	}

	idx := h.hash(key)
	entry := h.list[idx]

	if entry == nil {
		return "", nil
	}

	if entry.key == key {
		return entry.value, nil
	}

	for probe := range h.size {
		i := (idx + probe) & (h.size - 1)

		entry = h.list[i]

		if entry != nil && entry.key == key {
			return entry.value, nil
		}
	}

	return "", nil
}
