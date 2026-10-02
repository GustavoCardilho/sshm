package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func save(conn Connection) error {
	old, err := load()

	if len(conn.Name) == 0 {
		return errors.New("Nome vazio!")
	}

	if len(conn.Host) == 0 {
		return errors.New("Host vazio!")
	}

	if len(conn.User) == 0 {
		return errors.New("User vazio!")
	}

	if exist, _ := exists(old, conn.Name); exist == true {
		return errors.New("Name já utilizado!")
	}

	if err != nil {
		return err
	}
	old = append(old, conn)
	out, err := json.MarshalIndent(old, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(pathFile, out, 0600); err != nil {
		return err
	}
	return nil
}

func load() ([]Connection, error){
	data, err := os.ReadFile(pathFile)
	
	if err != nil {
		return nil, err
	}

	var conns []Connection
	if err := json.Unmarshal(data, &conns); err != nil {
		return nil, err
	}

	return conns, nil
}


func ask(q string, b *bufio.Reader, c *string) (error) {
	fmt.Printf("%s:", q)
	r, err := b.ReadString('\n')
	*c = strings.TrimSpace(r)
	return err
}

func exists(conns []Connection, name string) (bool, Connection) {
	var exists bool = false
	var searched Connection
	for _, val := range conns {
		if strings.EqualFold(val.Name, name) == true {
			exists = true
			searched = val
		}
	}
	return exists, searched
}

func connect(name string) error {
	data, err := load()
	if err != nil {
		return err
	}
	conn, err :=  detail(data, name)
	if err != nil {
		return err
	}
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", conn.User, conn.Host))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
	return nil
}


