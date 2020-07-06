package maps

// Search is used for search a string from map
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
