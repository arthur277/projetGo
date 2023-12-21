package librairie

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

// LibrairieEntry représente une entrée de dictionnaire avec un mot et une définition.
type LibrairieEntry struct {
	Mot        string `json:"mot"`
	Definition string `json:"definition"`
}

// Librairie est une structure de données représentant un dictionnaire.
type Librairie struct {
	entries []LibrairieEntry // Utilisation d'une slice au lieu d'une map
	mu      sync.RWMutex
}

// NewLibrairie crée et retourne une nouvelle instance de Librairie.
func NewLibrairie() *Librairie {
	return &Librairie{
		entries: make([]LibrairieEntry, 0),
	}
}

// handleMethodNotAllowed envoie une réponse d'erreur si la méthode HTTP n'est pas autorisée.
func (d *Librairie) handleMethodNotAllowed(w http.ResponseWriter, r *http.Request, allowedMethod string) {
	if r.Method != allowedMethod {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
}

// Add permet d'ajouter une nouvelle entrée au dictionnaire.
func (d *Librairie) Add(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodPost)

	var entry LibrairieEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	d.entries = append(d.entries, entry)
	w.WriteHeader(http.StatusCreated)
}

// Get permet de récupérer la définition d'un mot spécifique dans le dictionnaire.
func (d *Librairie) Get(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodGet)

	mot := r.URL.Query().Get("mot")

	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, entry := range d.entries {
		if entry.Mot == mot {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(entry)
			return
		}
	}

	http.Error(w, "Mot non trouvé", http.StatusNotFound)
}

// Remove permet de supprimer une entrée du dictionnaire en fonction du mot.
func (d *Librairie) Remove(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodDelete)

	mot := r.URL.Query().Get("mot")

	d.mu.Lock()
	defer d.mu.Unlock()

	for i, entry := range d.entries {
		if entry.Mot == mot {
			// Supprimer l'entrée en la retirant de la slice
			d.entries = append(d.entries[:i], d.entries[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// Si la boucle se termine sans trouver l'entrée, renvoyer une erreur
	http.Error(w, "Mot non trouvé", http.StatusNotFound)
}

// RemoveAll permet de supprimer toute la liste du dictionnaire.
func (d *Librairie) RemoveAll(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodDelete)

	d.mu.Lock()
	defer d.mu.Unlock()

	// Réinitialiser la slice à une slice vide
	d.entries = make([]LibrairieEntry, 0)
	w.WriteHeader(http.StatusOK)
}

// List renvoie la liste complète des entrées du dictionnaire.
func (d *Librairie) List(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodGet)

	d.mu.RLock()
	defer d.mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d.entries)
}

// ExportToFile exporte la liste des entrées du dictionnaire vers un fichier JSON.
func (d *Librairie) ExportToFile(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodGet)

	// Encode la liste en JSON
	jsonData, err := json.MarshalIndent(d.entries, "", "  ")
	if err != nil {
		http.Error(w, "Erreur", http.StatusInternalServerError)
		return
	}

	// Écrit le JSON dans un fichier
	file, err := os.Create("liste.json")
	if err != nil {
		http.Error(w, "Erreur", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	file.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
