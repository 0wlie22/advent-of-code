package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var INPUT_FILE = "test_input.txt"

func main() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// gameNr := getGameNr(line)
		getScore(getLineContent(line))
	}

}

func getGameNr(line string) string {
	return strings.Split(strings.Split(line, ":")[0], " ")[1]
}

func getLineContent(line string) string {
	return strings.Split(line, ":")[1]
}

func getScore(line string) int {
	games := strings.Split(line, ";")
	var subgameStruct struct {
		green int
		red   int
		blue  int
	}

	for _, game := range games {
		subgame := strings.Split(game, ",")
		for _, subgame := range subgame {
			subgame = strings.TrimSpace(subgame)
			if strings.Contains(subgame, "green") {
				subgameStruct.green++
			} else if strings.Contains(subgame, "red") {
				subgameStruct.red++
			} else if strings.Contains(subgame, "blue") {
				subgameStruct.blue++
			}
		}

	}

	return 0
}

func getScoreForSet(set string) int {
    takes := strings.Split(set, ",")
    for _, take := range takes {

    return 0
}
