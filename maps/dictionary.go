package maps

// Dictionary is a map for key value pairs for defining terms
type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already defined")
	ErrWordDoesNotExist = DictionaryErr("word is not defined")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search takes a key and returns the value associated to it in a dictionary
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add creates a new key value pair in a Dictionary, only if the key is unique
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update changes the definition for a given key with a new value, only if it exists already
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

// Delete removes an entry from the Dictionary given a key, it does not care if the key exists already or not since the default behavior is desired
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
