package service

import (
	"bufio"
	"fmt"
	"os"
	"quake-log-parser/schema"
	"regexp"
	"strings"
)

const (
	initGamePattern    = "InitGame"
	clientInfoPattern  = "ClientUserinfoChanged"
	killPattern        = "Kill"
	playerRegexPattern = `n\\(.*?)\\t`
	killRegexPattern   = `:\s(.*?)\skilled\s(.*?)\sby\s(.*)`
	worldKiller        = "<world>"
)

func ParseLogFile(filePath string) (map[string]*schema.Game, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro on open archive: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := make(map[string]*schema.Game)
	var currentGame *schema.Game
	gameCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.Contains(line, initGamePattern):
			gameCount++
			currentGame = &schema.Game{
				Kills:  make(map[string]int),
				Deaths: make(map[string]int),
			}
			games[fmt.Sprintf("jogo_%d", gameCount)] = currentGame

		case strings.Contains(line, clientInfoPattern):
			player := extractPlayerName(line)
			if player != "" && !contains(currentGame.Players, player) {
				currentGame.Players = append(currentGame.Players, player)
			}

		case strings.Contains(line, killPattern):
			currentGame.TotalKills++
			killer, victim, means := extractKillDetails(line)
			if killer != "" && victim != "" && means != "" {
				if killer != worldKiller {
					currentGame.Kills[killer]++
				} else {
					currentGame.Kills[victim]--
				}
				meansName := schema.LoadMeansOfDeath()[means]
				currentGame.Deaths[meansName]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error on read archive: %w", err)
	}

	return games, nil
}

func extractPlayerName(line string) string {
	re := regexp.MustCompile(playerRegexPattern)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func extractKillDetails(line string) (string, string, string) {
	re := regexp.MustCompile(killRegexPattern)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 3 {
		return matches[1], matches[2], matches[3]
	}
	return "", "", ""
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
