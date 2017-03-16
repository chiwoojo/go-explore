package main

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	router := httprouter.New()
	router.GET("/", showBooks)
	// router.GET("/s", handleS)
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}

func showBooks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
