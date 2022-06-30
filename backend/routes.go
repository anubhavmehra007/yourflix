package main
import (
	"strings"
	"time"
	"log"
	"net/http"
	"fmt"
	//"math/rand"
	//"strconv"
	//"encoding/json"
	"github.com/gorilla/mux"
	"os"
	"os/exec"
)
func test(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	var s string = "Hello Browser from GO"
	res.Write([]byte(s))
}
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
func serveFile(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req);
		id := vars["id"];
		var fileName string;
		if id == "1" {
			fileName = "./content/true_sight.mp4";
		} else {
			fileName = "test.mp4";
		}
		fmt.Printf("%s %s\n", id, fileName);
		http.ServeFile(res, req,fileName);
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", showMovies).Methods("GET");
	r.HandleFunc("/genres", showGenres).Methods("GET");
	r.HandleFunc("/movie/{id}", playMovie).Methods("GET");
	r.HandleFunc("/addMovie", addMovie).Methods("POST");
	r.HandleFunc("/subtitle/{id}", sendSubtitle).Methods("GET");
	r.HandleFunc("/subtitle", showSubtitlesList).Methods("GET");
	const serverString : string = fmt.Sprintf("%s %d", def.dbhost, def.port);
	srv := &http.Server{
        Handler:      r,
        Addr:         serverString,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    log.Fatal(srv.ListenAndServe())
}