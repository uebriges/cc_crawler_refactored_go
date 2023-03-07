package crawler

import "sync"

type LinkMap struct {
	entries map[string]bool
	sync.RWMutex
}

func NewLinkMap() *LinkMap {
	return &LinkMap{
		entries: map[string]bool{},
	}
}

// ContainsLink checks if link already exists
// in map and returns true/false
func (lm LinkMap) containsLink(newLink string) (containsLink bool) {
	lm.RLock()
	defer lm.RUnlock()
	_, containsLink = lm.entries[newLink]
	return containsLink
}

func (lm LinkMap) add(newLink string) {
	lm.Lock()
	lm.entries[newLink] = false
	lm.Unlock()
}
