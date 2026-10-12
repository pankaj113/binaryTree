package problem4

// Index is an ordered collection of string keys, each with an int value.
type Index struct {
	// TODO: your fields
}

// NewIndex returns an empty Index.
func NewIndex() *Index {
	// TODO: implement
	return &Index{}
}

// Put stores val under key. An existing key gets its value replaced.
func (ix *Index) Put(key string, val int) {
	// TODO: implement
}

// Get returns the value stored under key, and whether key is there.
func (ix *Index) Get(key string) (int, bool) {
	// TODO: implement
	return -1, true
}

// Len returns how many keys the Index holds.
func (ix *Index) Len() int {
	// TODO: implement
	return -1
}

// Keys returns every key, in sorted order.
func (ix *Index) Keys() []string {
	// TODO: implement
	return []string{"TODO"}
}

// Floor returns the largest key that is <= key.
func (ix *Index) Floor(key string) (string, bool) {
	// TODO: implement
	return "TODO", true
}

// Ceiling returns the smallest key that is >= key.
func (ix *Index) Ceiling(key string) (string, bool) {
	// TODO: implement
	return "TODO", true
}
