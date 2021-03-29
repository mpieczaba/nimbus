package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "help")
	}

	switch os.Args[1] {
	case "start":
		app := NewApp()

		app.Start()
	default:
		fmt.Println("Nimbus - extensible storage system focused on quick data access")
	}
}
