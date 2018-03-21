package main

import (
	"fmt"
	"github.com/johskw/score_management/src/management"
	"os"
)

func main() {
	switch os.Args[1] {
	case "create":
		management.Create()
	case "read":
		management.Read()
	case "delete":
		management.Delete()
	case "average":
		management.Average()
	default:
		fmt.Println(">invalid command: wrong argument")
	}
}
