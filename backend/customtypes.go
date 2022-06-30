package main
//to represent movies
type Movie struct {
	id string
	name string
	director string
	genres [] int
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