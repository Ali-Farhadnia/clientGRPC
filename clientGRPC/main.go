package main

import (
	"log"

	mainfunctions "github.com/Ali-Farhadnia/clientGRPC/MainFunctions"
	"github.com/Ali-Farhadnia/clientGRPC/cmd"
	userinput "github.com/Ali-Farhadnia/clientGRPC/models/UserInput"
)

var input userinput.UserInput

func main() {
	//set log flag
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := cmd.ConfigGRPC()
	if err != nil {
		panic(err)
	}
	err = cmd.ConfigMainFuncs()
	if err != nil {
		panic(err)
	}
	flag := true
	for {
		if flag {
			s, _ := mainfunctions.Help("")
			input.Output = s
			input.ShowOutput()
			flag = false
		}

		err = input.Getinput()
		if err != nil {
			continue
		}

		err = input.ParseInput()
		if err != nil {
			input.ShowOutput()
			continue
		}

		err = input.HandleInput()
		if err != nil {
			continue
		}

		_ = input.ShowOutput()

	}
}
