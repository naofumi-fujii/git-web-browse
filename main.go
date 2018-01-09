package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
)

func main() {
	u, err := url.Parse(getGitRemoteURL())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("https://" + u.Hostname() + u.Path)
}

func getGitRemoteURL() string {

	out, err := exec.Command("git", "remote", "get-url", "origin").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(out)
}
