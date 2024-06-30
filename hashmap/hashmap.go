package main

import (
	"hash/fnv"
)

// Calculate hash of a string.
func calculateHash(str string) int {
	h := fnv.New32a()
	h.Write([]byte(str))
	return int(h.Sum32())
}

// Enum for entry state.
const (
	empty = iota
	occupied
	deleted
)

type Entry struct {
	key   string
	value string
	state int
}

type OpenAddressingHashMap struct {
	Size          int
	OccupiedCount int
	entries       []Entry
}

func NewHashMap(size int) OpenAddressingHashMap {
	entries := make([]Entry, size)
	hashmap := OpenAddressingHashMap{Size: size, entries: entries}
	return hashmap
}

// Linear probing insert to hashmap. Return true if inserted successfully, otherwise false.
func (h *OpenAddressingHashMap) insertNoRehash(key string, value string) (ok bool) {
	hash := calculateHash(key) % h.Size
	for i := 0; i < h.Size; i++ {
		position := (hash + i) % h.Size
		if h.entries[position].state != occupied || h.entries[position].key == key {
			if h.entries[position].state != occupied {
				h.OccupiedCount++ // Increment only if the position was not already occupied
			}
			h.entries[position] = Entry{key: key, value: value, state: occupied}
			return true
		}
	}
	return false
}

// Insert to hashmap, double the size and rehash if load factor is too large.
func (h *OpenAddressingHashMap) Insert(key string, value string) (ok bool) {
	if l := h.GetLoadFactor(); l >= 0.5 {
		h.rehash()
	}
	return h.insertNoRehash(key, value)
}

// Get value for a key. Return value and status. If key is not in the hashmap, the status is false.
func (h *OpenAddressingHashMap) Get(key string) (value string, ok bool) {
	hash := calculateHash(key) % h.Size
	for i := 0; i < h.Size; i++ {
		// Linear probing search, start at the calculated position, if there wasn't a collision, return the value.
		// Ff there was a collision, check the following indices and look for the key, until you reach empty.
		// Empty means there was no key found in the hashmap.
		position := (hash + i) % h.Size // Wrap around if reached the end, so that all slots are checked.
		if h.entries[position].key == key {
			return h.entries[position].value, true
		}
		if h.entries[position].state == empty {
			return "", false
		}
	}
	return "", false
}

// Delete entry for a given key. Deleted values are marked as such in order not to break linear probing.
// Return true if key  deleted successfully, false if there was no such key in the hashmap.
func (h *OpenAddressingHashMap) Delete(key string) (ok bool) {
	hash := calculateHash(key) % h.Size
	for i := 0; i < h.Size; i++ {
		position := (hash + i) % h.Size
		if h.entries[position].state == empty {
			return false
		}
		if h.entries[position].key == key {
			h.entries[position] = Entry{state: deleted}
			h.OccupiedCount-- // Decrement since an entry has been deleted
			return true
		}
	}
	return false
}

// Get load factor, which is the number of entries divided by the capacity of the slice that's holding the entries.
func (h *OpenAddressingHashMap) GetLoadFactor() float32 {
	return float32(h.OccupiedCount) / float32(h.Size)
}

// Double the size of the slice holding the entries and rehash every occupied entry.
func (h *OpenAddressingHashMap) rehash() {
	oldEntries := h.entries

	newSize := h.Size * 2
	newEntries := make([]Entry, newSize)

	// Swap new for old
	h.Size = newSize
	h.entries = newEntries

	for _, entry := range oldEntries {
		if entry.state == occupied {
			h.insertNoRehash(entry.key, entry.value)
		}
	}
}

func main() {
}
