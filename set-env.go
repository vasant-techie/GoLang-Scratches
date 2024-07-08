package main

import (
	"fmt"
	"os"
)

func main() {
	accessKey := "your_access_key_here"
	secretKey := "your_secret_key_here"
	sessionToken := "your_session_token_here"

	err := setAWSEnvironmentVariables(accessKey, secretKey, sessionToken)
	if err != nil {
		fmt.Println("Error setting AWS environment variables:", err)
		return
	}

	fmt.Println("AWS environment variables set successfully.")
}

func setAWSEnvironmentVariables(accessKey, secretKey, sessionToken string) error {
	err := os.Setenv("AWS_ACCESS_KEY_ID", accessKey)
	if err != nil {
		return err
	}

	err = os.Setenv("AWS_SECRET_ACCESS_KEY", secretKey)
	if err != nil {
		// Rollback the AWS_ACCESS_KEY_ID environment variable if setting AWS_SECRET_ACCESS_KEY fails
		os.Unsetenv("AWS_ACCESS_KEY_ID")
		return err
	}

	err = os.Setenv("AWS_SESSION_TOKEN", sessionToken)
	if err != nil {
		// Rollback the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables if setting AWS_SESSION_TOKEN fails
		os.Unsetenv("AWS_ACCESS_KEY_ID")
		os.Unsetenv("AWS_SECRET_ACCESS_KEY")
		return err
	}

	return nil
}
