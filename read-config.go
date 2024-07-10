package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

// Config struct to hold the configuration values
type Config struct {
    RoleARN       string `json:"role_arn"`
    PrincipalARN  string `json:"principal_arn"`
    SAMLAssertion string `json:"saml_assertion"`
}

// Function to load configuration from a file
func loadConfig(file string) (Config, error) {
    var config Config
    configFile, err := os.Open(file)
    if err != nil {
        return config, err
    }
    defer configFile.Close()

    byteValue, _ := ioutil.ReadAll(configFile)
    json.Unmarshal(byteValue, &config)

    return config, nil
}

func main() {
    config, err := loadConfig("config.json")
    if err != nil {
        log.Fatalf("Error loading config file: %v", err)
    }

    fmt.Printf("Role ARN: %s\n", config.RoleARN)
    fmt.Printf("Principal ARN: %s\n", config.PrincipalARN)
    fmt.Printf("SAML Assertion: %s\n", config.SAMLAssertion)

    // Your code to fetch AWS access key, secret key, and session token using the config values
}
