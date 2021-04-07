package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/simia-tech/go-pop3"
)

func write(number int, data []byte, outDir string, user string) {
	os.MkdirAll(outDir, os.FileMode(0777))
	filename := fmt.Sprintf("%s/%s-%d", outDir, user, number)
	ioutil.WriteFile(filename, data, 0644)
}

func parseArgs() (string, string, string, string) {
	args := os.Args[1:]
	if len(args) != 4 {
		panic("Did not supply enough arguments. Usage: harvestpop3 host user pass output_directory.")
	}
	host := args[0]
	user := args[1]
	pass := args[2]
	outDir := args[3]

	fmt.Println(host, user, pass, outDir)
	return host, user, pass, outDir
}

func main() {
	var user string
	var pass string
	var host string
	var outDir string

	host, user, pass, outDir = parseArgs()

	fmt.Println("=== Harvesting Mails ===")
	client, err := pop3.Dial(host)
	if err != nil {
		panic(err)
	}
	err = client.Auth(user, pass)
	if err != nil {
		panic(err)
	}
	messages, _ := client.ListAll()
	for message := range messages {
		str, err := client.Retr(uint32(message) + 1)
		if err != nil {
			panic(err)
		}
		write(message, []byte(str), outDir, user)
	}
	fmt.Println("=== DONE ===")
}
