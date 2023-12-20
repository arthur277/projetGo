package librairie

import (
	"encoding/json"
	"io/ioutil"
	"sort"
)

// Entry représente un mot et sa définition.
type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

// Dictionary représente un dictionnaire avec un chemin de fichier et des entrées.
type Dictionary struct {
	FilePath   string
	Entries    []Entry
	AddChan    chan Entry
	RemoveChan chan string
}

// NewDictionary crée une nouvelle instance de Dictionary avec le chemin du fichier.
func NewDictionary(filePath string) *Dictionary {
	d := &Dictionary{
		FilePath:   filePath,
		AddChan:    make(chan Entry),
		RemoveChan: make(chan string),
	}
	go d.processOperations()
	return d
}

func (d *Dictionary) processOperations() {
	for {
		select {
		case entry := <-d.AddChan:
			d.Entries = append(d.Entries, entry)
		case word := <-d.RemoveChan:
			var newEntries []Entry
			for _, entry := range d.Entries {
				if entry.Word != word {
					newEntries = append(newEntries, entry)
				}
			}
			d.Entries = newEntries
		}
		d.saveToFile()
	}
}

// Add ajoute un mot avec sa définition au dictionnaire.
func (d *Dictionary) Add(word, definition string) error {
	entry := Entry{Word: word, Definition: definition}
	d.AddChan <- entry
	return nil
}

// Get récupère la définition d'un mot du dictionnaire.
func (d *Dictionary) Get(word string) (string, bool) {
	for _, entry := range d.Entries {
		if entry.Word == word {
			return entry.Definition, true
		}
	}
	return "", false
}

// Remove supprime un mot du dictionnaire.
func (d *Dictionary) Remove(word string) error {
	d.RemoveChan <- word
	return nil
}

// List renvoie la liste des mots dans le dictionnaire, triés par ordre alphabétique.
func (d *Dictionary) List() ([]string, error) {
	var words []string
	for _, entry := range d.Entries {
		words = append(words, entry.Word)
	}
	sort.Strings(words)
	return words, nil
}

// saveToFile sauvegarde les entrées du dictionnaire dans un fichier JSON.
func (d *Dictionary) saveToFile() {
	jsonData, err := json.MarshalIndent(d.Entries, "", "  ")
	if err != nil {
		// Gestion de l'erreur (par exemple, log ou fmt)
		return
	}
	ioutil.WriteFile(d.FilePath, jsonData, 0644)
}
