package pkg01

import (
	"net/http"
)

var Project_name01_pkg01 string = "project_name01_pkg01_big"
var project_name01_pkg01 string = "project_name01_pkg01_small"

func Sample_server() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.ListenAndServe(":8080", nil)
}
