package maps

import "errors"

type Dictionary map[string]string

// Search takes a key and returns the value associated to it in a dictionary
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
