package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	for idx, val := range data {
		fmt.Printf("%d - %s (%s@%s)\n", idx + 1, val.Name, val.User, val.Host)
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