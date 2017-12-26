package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		remoteURL := scanner.Text()
		u, err := url.Parse(remoteURL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("https://" + u.Hostname() + u.Path)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
