package cmd

import (
	"fmt"
	"os"

	"github.com/mpieczaba/nimbus/app"
)

func Execute() {
	fmt.Print("Nimbus - extensible storage system\n\n")

	if len(os.Args) < 2 {
		os.Args = append(os.Args, "help")
	}

	switch os.Args[1] {
	case "start":
		app.New().Start()
	}
}
