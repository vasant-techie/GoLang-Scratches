package main

import (
	"fmt"
	"golang.org/x/ini"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	awsCredentialsFile = "~/.aws/credentials"
)

func main() {
	// Read AWS credentials and session token from environment variables
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	sessionToken := os.Getenv("AWS_SESSION_TOKEN")

	// Update credentials file
	if accessKey != "" && secretKey != "" {
		err := updateAWSCredentials(accessKey, secretKey, sessionToken)
		if err != nil {
			fmt.Println("Error updating AWS credentials file:", err)
			return
		}
		fmt.Println("AWS credentials updated successfully.")
	} else {
		fmt.Println("AWS credentials not found in environment variables.")
	}
}

func updateAWSCredentials(accessKey, secretKey, sessionToken string) error {
	// Expand the home directory path
	credsFile := expandHomeDir(awsCredentialsFile)

	// Read the existing credentials file
	data, err := ioutil.ReadFile(credsFile)
	if err != nil {
		return err
	}

	// Load the INI formatted data
	cfg, err := ini.Load(data)
	if err != nil {
		return err
	}

	// Set the credentials for the [default] profile
	section, err := cfg.GetSection("default")
	if err != nil {
		section, err = cfg.NewSection("default")
		if err != nil {
			return err
		}
	}

	// Update the values
	section.Key("aws_access_key_id").SetValue(accessKey)
	section.Key("aws_secret_access_key").SetValue(secretKey)
	if sessionToken != "" {
		section.Key("aws_session_token").SetValue(sessionToken)
	} else {
		section.DeleteKey("aws_session_token")
	}

	// Write the updated credentials back to the file
	err = cfg.SaveTo(credsFile)
	if err != nil {
		return err
	}

	return nil
}

// expandHomeDir expands the `~` symbol to the user's home directory path
func expandHomeDir(path string) string {
	if path[:2] == "~/" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting user home directory:", err)
			return path
		}
		return filepath.Join(homeDir, path[2:])
	}
	return path
}
