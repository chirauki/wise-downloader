/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"

	"github.com/chirauki/wise-downloader/cmd"
)

func main() {
	cmd := cmd.NewRoot()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
