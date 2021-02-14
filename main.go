package main

import (
	"github.com/mpieczaba/nimbus/cmd"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cmd.Execute()
}
