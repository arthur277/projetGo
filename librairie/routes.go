package librairie

import "net/http"

// SetupRoutes configure les routes pour le dictionnaire
func SetupRoutes(Librairie *Librairie) {
	http.HandleFunc("/add", Librairie.Add)
	http.HandleFunc("/get", Librairie.Get)
	http.HandleFunc("/remove", Librairie.Remove)
	http.HandleFunc("/list", Librairie.List)
	http.HandleFunc("/removeall", Librairie.RemoveAll)
	http.HandleFunc("/export", Librairie.ExportToFile)

}
