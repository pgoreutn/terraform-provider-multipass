package client

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

var multipass, multipassError = getBinary()

// Client holds all of the information required to connect to multipass vm
type Client struct {
	name   string
	cpu    int
	memory int
	disk   string
}

type Operations interface {
	Create() error

	Exists() error

	AddVm() error

	Delete() error
}

func NewClient(name string, cpu int, memory int, disk string) *Client {

	return &Client{
		name:   name,
		cpu:    cpu,
		memory: memory,
		disk:   disk,
	}
}

func (*Client) Create() error {

	return nil
}

func (cli *Client) Exists() error {
	return cli.multipass_exists()
}

//multipass_exists verify if the multipass bin exist or not.
func (cli *Client) multipass_exists() error {

	if multipassError != nil {
		if _, err := exec.LookPath(multipass); err != nil {
			panic(err)
		}
	}
	_, err := exec.Command(multipass, "info", cli.name).Output()
	if err != nil {
		return err
	}

	return nil
}

func (cli *Client) AddVm() error {

	// multipass launch ubuntu --name master --cpus 2 --mem 2G --disk 8G
	switch os := runtime.GOOS; {
	case os != "windows":

		out, err := exec.Command("multipass", "launch", "ubuntu", "--name", cli.name, "--cpus", "2", "--mem", "2G", "--disk", "8G").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The date is %s\n", out)

		return err
	case os == "Windows":
		fmt.Println("Windows")
	}

	return nil
}

func (cli *Client) Delete() error {
	// multipass launch ubuntu --name master --cpus 2 --mem 2G --disk 8G
	switch os := runtime.GOOS; {
	case os != "windows":

		var _, err = exec.Command("multipass", "delete", cli.name).Output()
		if err != nil {
			panic(err)
			log.Fatal(err)
		}

		_, err = exec.Command("multipass", "purge").Output()
		if err != nil {
			panic(err)
			log.Fatal(err)
		}

		return err
	case os == "Windows":
		fmt.Println("Windows")
	}

	return nil
}

func getBinary() (string, error) {
	switch os := runtime.GOOS; {
	case os != "windows":
		return "multipass", nil

	case os == "Windows":
		return "multipass.exe", nil
	}

	return "", errors.New("No multipass binary for current platfom")
}
