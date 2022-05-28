package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Notch-Technologies/dotshake/cmd/dotshake/cmd"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
		log.Fatalf("Failed to load the env vars: %v", err)
	}
}

func main() {
	if err := cmd.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}