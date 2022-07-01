package main
//to represent movies
type Movie struct {
	Id string
	Name string
	Director string
	Genres [] int
	path string
}
//to represent config
type CONFIG struct {
	contentPath string
	dbuser string
	dbpassword string
	dbhost string
	dbport int
	dbname string
	port int
	ffmpeg string
}