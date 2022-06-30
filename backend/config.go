package main
type struct CONFIG {
	contentPath string
	dbuser string
	dbpassword string
	dbhost string
	database string
	port int
	ffmpeg string
}
const def = CONFIG{
	"/home/anubhav/go-projects/content/",
	"root",
	"root@123", 
	"localhost",
	"amvflix",
	8000,
	"ffmpeg"
}