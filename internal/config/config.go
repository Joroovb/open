package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Program struct {
	Path      string
	FileTypes []string
}

var ErrConfigParsingError = errors.New("error reading config file")

func Default() []Program {
	return []Program{
		{
			Path:      "imv",
			FileTypes: []string{".png", ".jpg", ".jpeg", ".webp", ".gif"},
		},
		{
			Path:      "mpv",
			FileTypes: []string{".mp4", ".avi"},
		},
		{
			Path:      "zathura",
			FileTypes: []string{".pdf"},
		},
	}
}

func Get() []Program {
	file, err := os.Open("config.yml")
	if err != nil {
		return Write()
	}

	var progs []Program
	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&progs)
	if err != nil {
		fmt.Fprint(os.Stderr, ErrConfigParsingError)
		os.Exit(1)
	}

	return progs
}

func Write() []Program {
	config := Default()

	data, _ := yaml.Marshal(config)
	err := os.WriteFile("config.yml", data, 0644)
	if err != nil {
		os.Exit(1)
	}

	return Get()
}
