package cmd

import (
	"fmt"
	"os"

	"github.com/mpieczaba/nimbus/core"
)

func Execute() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "help")
	}

	switch os.Args[1] {
	case "start":
		app := core.NewApp()

		app.Start()
	default:
		fmt.Println("Nimbus - extensible storage system with quick access to data")
	}
}
