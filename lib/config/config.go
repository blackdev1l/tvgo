package config

import (
	"fmt"
	"log"
	"os"
	"os/user"

	fs "github.com/blackdev1l/tvgo/lib/filesystem"
)

type Configuration struct {
	Path   string
	daemon bool
}

func (c *Configuration) CheckConfig() {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configPath := usr.HomeDir + "/.config/tvgo"
	if fs.Exists(configPath) == false {
		fmt.Println("Creating tvgo folder in ~/.config")
		os.MkdirAll(configPath, os.ModePerm)
	}

	c.Path = configPath
}
