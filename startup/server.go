package startup

import (
	"github.com/gorilla/mux"
	"log"
	"lucytech/internal/controller"
	"net/http"
)

func Server() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")

	r.HandleFunc("/api/url-info", controller.GetHTMLContextMetadata).Methods("POST")

	//mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Return all the comments.")
	//})
	//
	//mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(w, "Post a new comment.")
	//})

	// or localhost:8080
	//if err := http.ListenAndServe(":8080", mux); err != nil {
	//	fmt.Println(err.Error())
	//}

	log.Fatal(http.ListenAndServe(":8080", r))
}
