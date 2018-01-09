package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/browser"
)

func main() {
	u, err := url.Parse(getGitRemoteURL())
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args
	url := ""
	if len(args) < 2 {
		url = strings.TrimSuffix("https://"+u.Hostname()+u.Path, "\n")
	} else {
		url = strings.TrimSuffix("https://"+u.Hostname()+u.Path, "\n") + "/commit/" + args[1]
	}
	browser.OpenURL(url)
}

func getGitRemoteURL() string {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := string(out)
	r1 := strings.Replace(url, "git@github.com:", "ssh://git@github.com/", -1)
	r2 := strings.Replace(r1, ".git", "", -1)
	return r2
}
