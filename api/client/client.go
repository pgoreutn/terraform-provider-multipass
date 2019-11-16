package client

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"os/exec"
	"runtime"
)

var multipass, multipassError = getBinary()

// Client holds all of the information required to connect to multipass vm
type VMClient struct {
	name   string
	cpu    int
	memory int
	disk   string
}

func (*VMClient)NewClient(d *schema.ResourceData) *VMClient {
	return &VMClient{
		name:   d.Get("name").(string),
		cpu:    d.Get("cpu").(int),
		memory: d.Get("memory").(int),
		disk:   d.Get("disk").(string),
	}
}

func (*VMClient) Create() error {
	return nil
}

func (cli *VMClient) Exists() error {
	return cli.multipass_exists()
}

func (cli *VMClient) multipass_exists() error {

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

func (cli *VMClient) AddVm() error {

	_, err := exec.Command(multipass, "launch", "ubuntu", "--name", cli.name, "--cpus", "2", "--mem", "2G", "--disk", "8G").Output()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (cli *VMClient) Delete() error {
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
