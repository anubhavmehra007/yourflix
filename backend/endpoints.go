package main
import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)
func showMovies(dbworker DBWorker) http.HandlerFunc {
	movies := dbworker.allMovies();
	jsonResponse, err := json.Marshal(movies);
	if err != nil {
		panic(err);
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applicaiton/json");
		w.WriteHeader(http.StatusOK);
		w.Write(jsonResponse);
	}
}
func findMovie(dbworker DBWorker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r);
		id := vars["id"];
		movie := dbworker.theMovie(id);
		jsonResponse, err := json.Marshal(movie);
		if err != nil {
			panic(err);
		}
		w.Header().Set("Content-Type", "applicaiton/json");
		w.WriteHeader(http.StatusOK);
		w.Write(jsonResponse);
	}
}
func playMovie(dbworker DBWorker) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) { 
		vars := mux.Vars(r);
		id := vars["id"];
		movie := dbworker.theMovie(id);
		http.ServeFile(w, r,movie.path);
	}
}