package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ConfigFile struct {
	MaxPlayers uint32 `json:"max_players"`
	OnlineMode bool   `json:"online_mode"`
	MOTD       string `json:"motd"`
}

func LoadConfig() ConfigFile {
	if !IsConfigCreated() {
		return CreateDefaultConfig()
	}

	// Read the config.json file
	file, _ := os.Open("config.json")

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Close the file
	file.Close()

	// Check if the file is a valid JSON
	config, err := readJSON(content)
	if err != nil {
		fmt.Println("[ZohyMC] The config.json file is not a valid JSON file! Please fix it and restart the server.")
		os.Exit(1)
	}

	fmt.Println("[ZohyMC] Loaded the server configuration successfully!")

	return config
}

func IsConfigCreated() bool {
	file, err := os.Open("config.json")
	file.Close()

	return err == nil
}

func CreateDefaultConfig() ConfigFile {
	config := ConfigFile{
		MaxPlayers: 20,
		OnlineMode: true,
		MOTD:       "ZohyMC\nA Minecraft server written in Go",
	}

	// Convert the struct to indented JSON
	json, err := json.MarshalIndent(config, "", "	")
	if err != nil {
		panic(err)
	}

	// Write the JSON to the config.json file
	err = os.WriteFile("config.json", json, 0644)
	if err != nil {
		panic(err)
	}

	return config
}

func readJSON(data []byte) (config ConfigFile, err error) {
	err = json.Unmarshal(data, &config)
	return
}
