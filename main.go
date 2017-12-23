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
		remoteURL := scanner.Text()
		splitted := strings.Split(remoteURL, "/")
		repoNameWithGit := splitted[len(splitted)-1]
		repoName := strings.Replace(repoNameWithGit, ".git", "", -1)
		userName := splitted[len(splitted)-2]
		hostNameWithGit := splitted[len(splitted)-3]
		hostName := strings.Replace(hostNameWithGit, "git@", "", -1)
		result := []string{hostName, userName, repoName}
		fmt.Println(strings.Join(result, "/"))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
