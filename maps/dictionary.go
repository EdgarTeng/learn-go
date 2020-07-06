package maps

import "errors"

// Dictionary is short for map[string]string
type Dictionary map[string]string

// Search is used for search a string from map
func (d Dictionary) Search(word string) string {
	return d[word]
}

// Add used for put a key-value pair to dict
func (d Dictionary) Add(key, word string) error {
	if d.Search(key) != "" {
		return ErrAlreadyExists
	}
	d[key] = word
	return nil
}

// ErrAlreadyExists stands for already have key in dict
var ErrAlreadyExists = errors.New("The key already exists in dictionary")
