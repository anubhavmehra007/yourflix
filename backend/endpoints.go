package main
import (
	"net/http"
	"github.com/gorilla/mux"
)
func showMovies(dbworker DBWorker) http.HandlerFunc {
	//movies := dbworker.allMovies();
	//process movies and parse json;
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func findMovie(dbworker DBWorker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r);
		id := vars["id"];
		movie := dbworker.theMovie(id);
		http.ServeFile(w, r,movie.path);
	}
}