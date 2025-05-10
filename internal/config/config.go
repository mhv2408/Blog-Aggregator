package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	// gets the config file path
	homePath, err := os.UserHomeDir() //gives the Home directory path
	if err != nil {
		fmt.Print(err)
		return "", fmt.Errorf("unable to find the homedir")
	}
	return homePath + "/workspace/github.com/mhv2408/gator/" + configFileName, nil
}

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	json_file_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("unable to find the path")
	}

	json_file, err := os.Open(json_file_path) //opening the json file at the path
	if err != nil {
		return Config{}, err
	}
	defer json_file.Close() //closing the json file after all operations are performed

	configStruct := Config{}                 // return this struct as value
	byteValue, _ := io.ReadAll(json_file)    // converts the json to []byte
	json.Unmarshal(byteValue, &configStruct) //converts the []byte into struct

	return configStruct, nil
}
func write(cfg Config) error {
	// Writes the config struct onto the json
	json_file_path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("unable to find the path")
	}
	file, err := os.Create(json_file_path)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = "harsha"
	return write(*cfg)

}
