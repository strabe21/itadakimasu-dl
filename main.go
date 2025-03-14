package main

import (
	"itadakimasu-dl/internal/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
