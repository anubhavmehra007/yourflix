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
	result, err := dbworker.Query("Select * from movies");
	if err != nil {
		panic(err)
	}
	for result.Next() {
		var movie Movie;
		var id int;
		var genres[] int;
		err = result.Scan(&id, &Movie.name, &Movie.director, &Movie.path);
		if err != nil {
			panic(err);
		}
		genres, err := db.Query(fmt.Sprintf("SELECT genre_id from movie_genre where movie_id=%d"),id);
		if err != nil {
			panic(err);
		}
	}
	return nil;
}
func (dbworker *DBWorker) theMovie(id string)  Movie {
	movie := Movie {
		"1",
		"True Sight TI7",
		"Valve",
		[] int {1},
		"/home/anubhav/react/your-flix/backend/content/ti7.mp4" };
	return movie;
}