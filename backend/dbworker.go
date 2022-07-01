package main
import (
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
type DBWorker struct {
	db *sql.DB

}
func NewDBWorker() DBWorker {
 	def  := CONFIG{
	"/home/anubhav/go-projects/content/",
	"root",
	"root@123", 
	"localhost",
	3306,
	"amvflix",
	8000,
	"ffmpeg" }
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",def.dbuser,def.dbpassword, def.dbhost, def.dbport, def.dbname);
	db, err := sql.Open("mysql", dbString);
	if err != nil {
		panic(err);
	}
	db.SetConnMaxLifetime(time.Minute * 3);
	db.SetMaxOpenConns(10);
	db.SetMaxIdleConns(10);
	dbworker  := DBWorker { db }
	return dbworker;
}
func (dbworker *DBWorker) allMovies() [] Movie {
	result, err := dbworker.db.Query("Select * from movies");
	var movies []Movie;
	if err != nil {
		panic(err)
	}
	for result.Next() {
		var movie Movie;
		var id int;
		var genres[] int;
		err = result.Scan(&id, &movie.Name, &movie.Director, &movie.path);
		if err != nil {
			panic(err);
		}
		genresResults, err := dbworker.db.Query(fmt.Sprintf("SELECT genre_id from movie_genre where movie_id='%d'", id));
		if err != nil {
			panic(err);
		}
		for genresResults.Next() {
			var genre int;
			err = genresResults.Scan(&genre);
			if err != nil {
				panic(err);
			}
			genres = append(genres, genre);
		}
		movie.Id = fmt.Sprintf("%d", id);
		movie.Genres = genres;
		movies = append(movies, movie);
	}
	return movies;
}
func (dbworker *DBWorker) theMovie(id string)  Movie {
	var movie Movie;
	var idNum int;
	err := dbworker.db.QueryRow(fmt.Sprintf("SELECT * FROM movies WHERE movie_id = '%s'", id)).Scan(&idNum, &movie.Name, &movie.Director, &movie.path);
	if err != nil {
		panic(err);
	}
	movie.Id = id;
	var genres []int;
	result, err := dbworker.db.Query(fmt.Sprintf("SELECT genre_id from movie_genre WHERE movie_id = '%s'",id));
	if err != nil {
		panic(err);
	}
	for result.Next() {
		var genreId int;
		result.Scan(&genreId);
		genres = append(genres, genreId);
	}
	movie.Genres = genres;
	return movie;
}