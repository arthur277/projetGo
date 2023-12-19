package main

import (
	"fmt"
	"sort"
)

type Dictionary struct {
	m map[string]string
}

func NewDictionary() Dictionary {
	return Dictionary{m: make(map[string]string)}
}

func (d *Dictionary) Add(word, definition string) {
	d.m[word] = definition
}

func (d *Dictionary) Get(word string) (string, bool) {
	definition, ok := d.m[word]
	return definition, ok
}

func (d *Dictionary) Remove(word string) {
	delete(d.m, word)
}

func (d *Dictionary) List() []string {
	var words []string
	for word := range d.m {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func main() {
	// Créez un nouveau dictionnaire
	d := NewDictionary()

	// Ajoutez des mots et des définitions
	d.Add("chèvre", "Mammifère herbivore, sauvage ou domestique")
	d.Add("lion", "Mammifère (canidé) carnivore aux multiples races, caractérisé par sa facilité à être  par une course rapide, un excellent odorat et par son cri spécifique")
	d.Add("oiseau", "Vertébré ovipare, couvert de plumes et d'écailles cornées, à respiration pulmonaire, homéotherme, aux mâchoires sans dents revêtues d'un bec corné, et aux membres antérieurs, ou ailes, normalement adaptés au vol.")

	// Obtenez la définition d'un mot
	definition, ok := d.Get("lion")
	if ok {
		fmt.Println("Définition de 'lion':", definition)
	}

	// Supprimez un mot
	d.Remove("oiseau")

	// Obtenez la liste triée des mots
	words := d.List()
	for _, word := range words {
		fmt.Println(word)
	}
}