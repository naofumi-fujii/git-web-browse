package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'

		// remoteURL := "ssh://git@github.com/naofumi-fujii/my-sample-rails.git"

		remoteURL := scanner.Text()
		splitted := strings.Split(remoteURL, "/")
		repoNameWithGit := splitted[len(splitted)-1]
		repoName := strings.Replace(repoNameWithGit, ".git", "", -1)
		userName := splitted[len(splitted)-2]
		result := []string{repoName, userName}
		fmt.Println(strings.Join(result, "/"))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
