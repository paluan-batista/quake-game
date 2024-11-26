package main

import (
	"encoding/json"
	"fmt"
	"os"
	"quake-log-parser/internal/service"
)

func main() {
	games, err := service.ParseLogFile("qgames.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(games); err != nil {
		fmt.Println("Error on generate JSON:", err)
	}
}
