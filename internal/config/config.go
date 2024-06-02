package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Config struct {
	MajorVersion int
	MinorVersion int
	PatchVersion int
}

type version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

var instance *Config

// Instance returns a populated configuration. If the configuration instance has not yet been created, a new
// instance will be created and returned.
func Intance() *Config {
	if instance == nil {
		instance = createInstance()
	}

	return instance
}

func createInstance() *Config {

	version := readVersionConfig()

	return &Config{
		MajorVersion: version.Major,
		MinorVersion: version.Minor,
		PatchVersion: version.Patch,
	}
}

func readVersionConfig() *version {
	jsonFile, err := os.Open("version.json")
	if err != nil {
		fmt.Println("Unable to read version.json file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var version version

	json.Unmarshal(byteValue, &version)

	return &version
}

// GetVersion returns the semantic version string (major.minor.patch)
func (c *Config) GetVersion() string {
	return strconv.Itoa(c.MajorVersion) + "." + strconv.Itoa(c.MinorVersion) + "." + strconv.Itoa(c.PatchVersion)
}
