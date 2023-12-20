package main

import (
	"projetGo/librairie"
	"fmt"
	"time"
)

func main() {
	filePath := "dictionary.json"
	d := manipulation_dictionnaire.NewDictionary(filePath)

	// Ajout de mots et de définitions de manière concurrente
	go func() {
		handleError(d.Add("java", "Langage de programmation informatique")
		handleError(d.Add("php", "langage programmation backend")
		handleError(d.Add("css", "combiné avec le html perm de poduire un visuel")
	}()

	// Suppression du mot "css" de manière concurrente
	go func() {
		time.Sleep(500 * time.Millisecond) // Attendre un peu pour simuler un traitement concurrent
		if err := d.Remove("css"); err != nil {
			handleError(err, "Erreur lors de la suppression")
		}
	}()

	// Attente de la fin des opérations
	time.Sleep(2 * time.Second)

	// Récupération de la définition du mot "chat"
	definition, ok := d.Get("java")
	if ok {
		fmt.Println("Définition de 'java':", definition)
	} else {
		fmt.Println("Mot non trouvé.")
	}

	// Liste alphabétique des mots restants
	words, err := d.List()
	if err != nil {
		handleError(err, "Erreur lors de la récupération de la liste")
	}
	fmt.Println("Mots restants:", words)
}

func handleError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %v\n", message, err)
	}
}
