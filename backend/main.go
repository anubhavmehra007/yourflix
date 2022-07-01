package main
import (
	"strings"
	"time"
	"log"
	"net/http"
	"fmt"
	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
	"os"
	"os/exec"
)
func convertSubtitles(filePath string) {
	_, err := os.Stat(filePath);
	if err != nil {
		fmt.Printf("File not found");
		log.Fatal(err);
	} else {
		if strings.Contains(filePath, ".srt") {
			newFileName := strings.Replace(filePath, "srt", "vtt", 1);
			cmd := exec.Command("ffmpeg", "-i", filePath, newFileName);
			output,err := cmd.Output()
			if err != nil {
		fmt.Printf("Error running Command");
				log.Fatal(err);
			} else {
				fmt.Printf("%s", string(output));
			}
		}
	}

}
func main() {

	dbworker := NewDBWorker();
	defer dbworker.db.Close();
 	def  := CONFIG{
	"/home/anubhav/go-projects/content/",
	"root",
	"root@123", 
	"localhost",
	3306,
	"amvflix",
	8000,
	"ffmpeg" }
	r := mux.NewRouter()
	r.HandleFunc("/movies/", showMovies(dbworker)).Methods("GET","OPTIONS");
	//r.HandleFunc("/genres", showGenres(dbworker)).Methods("GET", "OPTIONS");
	r.HandleFunc("/movie/{id}", findMovie(dbworker)).Methods("GET" ,"OPTIONS");
	r.HandleFunc("/play/{id}", playMovie(dbworker)).Methods("GET", "OPTIONS");
	//r.HandleFunc("/addMovie", addMovie).Methods("POST");
	//r.HandleFunc("/subtitle/{id}", sendSubtitle).Methods("GET");
	//r.HandleFunc("/subtitle", showSubtitlesList).Methods("GET");
	serverString := fmt.Sprintf("%s:%d", def.dbhost, def.port);
	srv := &http.Server{
        Handler:      r,
        Addr:         serverString,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    log.Fatal(srv.ListenAndServe())
}