package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	u, err := url.Parse(getGitRemoteURL())
	if err != nil {
		log.Fatal(err)
	}
	url := "https://" + u.Hostname() + u.Path
	openBrowser(url)
}

func getGitRemoteURL() string {

	out, err := exec.Command("git", "remote", "get-url", "origin").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(out)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
