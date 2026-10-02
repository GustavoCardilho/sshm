package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Connection struct {
	Name string `json:"name"`
	Host string `json:"host"`
	User string `json:"user"`
}

var pathFile string

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}
}

func run() error {
	if err := setup(); err != nil {
		return err
	}

	if len(os.Args) < 2 {
		return errors.New("uso: sshm new | list | connect <nome>")
	}

	command := os.Args[1]

	switch command {
	case "new":
		err := register()
		if err != nil {
			return err
		}
	case "list":
		err := list()
		if err != nil {
			return err
		}
	case "connect":
		if len(os.Args) < 3 {
			return errors.New("Nome de conexão inválida!")
		}
		name := os.Args[2]
		connect(name)
	default:
		return errors.New("Comando não encontrado!")
	}


	return nil
}

func setup() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	appDir := filepath.Join(configDir, "sshm")
	err = os.MkdirAll(appDir, 0700)
	if err != nil {
		return err
	}

	pathFile = filepath.Join(appDir, "configuration.json")

	f, err := os.Open(pathFile)
	if errors.Is(err, os.ErrNotExist) {
		cont, err := json.Marshal([]Connection{})
		if err != nil {
			return err
		}
		if err := os.WriteFile(pathFile, cont, 0600); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

