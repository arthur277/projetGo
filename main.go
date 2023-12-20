package main

import (
	"goProjet/projetGo/librairie"
	"fmt"
	"net/http"
)

const port = 8030

func main() {
	Librairie := librairie.NewLibrairie()
	librairie.SetupRoutes(Librairie)

	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %s\n", err)
	}
}
