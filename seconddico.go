

import "sort"

type Dictionary map[string]string

func NewDictionary() Dictionary {
	return make(Dictionary)
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Get(word string) (string, bool) {
	definition, ok := d[word]
	return definition, ok
}

func (d Dictionary) Remove(word string) {
	delete(d, word)
}

func (d Dictionary) List() []string {
	var words []string
	for word := range d {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}