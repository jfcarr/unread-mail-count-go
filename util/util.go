package util

import (
	"encoding/json"
	"errors"
	"jfcarr/unread-mail-count/models"
	"log"
	"os"
	"path/filepath"
)

// Get the working directory for the current executable.
func GetExePath() string {
	exPath, err := os.Executable()

	if err != nil {
		return ""
	}

	exPathDir := filepath.Dir(exPath)
	return exPathDir
}

// Retrieve a password from an environment variable.
func GetPasswordFromEnvironment(envVariable string) string {
	password := os.Getenv(envVariable)

	return password
}

// Check to see if the specified file exists.
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !errors.Is(err, os.ErrNotExist)
}

// Generate a fully-qualified path to the given file.
func QualifiedFile(fileName string) string {
	return filepath.Join(GetExePath(), fileName)
}

// Load list of accounts to check from the app .json file.
func LoadAccounts() []models.Account {
	var configFile string
	if FileExists(QualifiedFile("unread-mail-count-live.json")) {
		configFile = QualifiedFile("unread-mail-count-live.json")
	} else {
		configFile = QualifiedFile("unread-mail-count.json")
	}

	jsonFile, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	var data []models.Account

	if err := json.NewDecoder(jsonFile).Decode(&data); err != nil {
		log.Fatal(err)
	}

	return data
}
