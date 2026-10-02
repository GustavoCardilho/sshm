package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

func register() error {
	var nconn = Connection{}

	reader := bufio.NewReader(os.Stdin)
	if err := ask("Name", reader, &nconn.Name); err != nil {
		return err
	}

	if err := ask("Host", reader, &nconn.Host); err != nil {
		return err
	}

	if err := ask("User", reader, &nconn.User); err != nil {
		return err
	}
	
	
	if err := save(nconn); err != nil {
		return err
	}
	return nil
}

func list() error {
	data, err := load()
	if err != nil {
		return err
	}
	var options []huh.Option[int]
	for idx, val := range data {
		label := fmt.Sprintf("%d - %s (%s@%s)", idx + 1, val.Name, val.User, val.Host)
		options = append(options, huh.NewOption(label, idx))
	}

	var choice int
	err = huh.NewSelect[int]().
		Title("Escolha uma opção").
		Options(options...).
		Value(&choice).
		Run()

	if err != nil {
		return err
	}

	if err := connect(data[choice].Name); err != nil {
		return err
	}

	return nil
}

func detail(conns []Connection, name string) (Connection, error) {
	exist, data := exists(conns, name)
	if exist == false {
		return Connection{}, errors.New("Conexão não encontrada!")
	}
	return data, nil
}