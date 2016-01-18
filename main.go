package main

import (
	"fmt"
	"os"

	"github.com/ariarijp/gomoku2/commands"
)

func main() {
	cmd := os.Args[1]

	if cmd == "head" {
		checkArgs(3, fmt.Sprintf("Usage: %s head host:port", os.Args[0]))
		url := os.Args[2]
		commands.Head(url)
	} else if cmd == "get" {
		checkArgs(3, fmt.Sprintf("Usage: %s get host:port", os.Args[0]))
		url := os.Args[2]
		commands.Get(url)
	} else if cmd == "client_get" {
		checkArgs(3, fmt.Sprintf("Usage: %s client_get host:port", os.Args[0]))
		url := os.Args[2]
		commands.ClientGet(url)
	} else if cmd == "file_server" {
		checkArgs(3, fmt.Sprintf("Usage: %s file_server path", os.Args[0]))
		path := os.Args[2]
		commands.FileServer(path)
	} else if cmd == "print_env" {
		checkArgs(3, fmt.Sprintf("Usage: %s print_env path", os.Args[0]))
		path := os.Args[2]
		commands.PrintEnv(path)
	} else if cmd == "server_handler" {
		commands.ServerHandler()
	} else {
		fmt.Printf("%s is unsupported command\n", cmd)
		os.Exit(1)
	}
}

func checkArgs(expectedArgc int, usage string) {
	if len(os.Args) != expectedArgc {
		fmt.Println(usage)
		os.Exit(1)
	}
}
