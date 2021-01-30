package main

import (
	"fmt"
	"github.com/other_project_name01/other_pkg01"
	"github.com/project_name01/pkg01"
)

func main() {
	fmt.Println("HelloWorld!")
	fmt.Println(pkg01.Project_name01_pkg01)
	// これはエラー
	// fmt.Println(pkg01.message)
	fmt.Println(other_pkg01.Other_project_name01_other_pkg01)
	pkg01.Sample_server()
}
