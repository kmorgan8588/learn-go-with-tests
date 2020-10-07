package maps

type Dictionary map[string]string

// Search takes a key and returns the value associated to it in a dictionary
func (d Dictionary) Search(key string) string {
	return d[key]
}
