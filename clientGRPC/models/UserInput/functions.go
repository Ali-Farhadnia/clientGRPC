package userinput

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Ali-Farhadnia/clientGRPC/config"
)

//Getinput() get input from user
func (input *UserInput) Getinput() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("simpleCRUD#")
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	input.RawInput = text
	return nil
}

//ParseInput() parse input to do the job
func (input *UserInput) ParseInput() error {
	input.RawInput = input.RawInput[:len(input.RawInput)-1]
	res := strings.SplitN(input.RawInput, " ", 2)

	_, valid := config.App.Functions[res[0]]
	if !valid {
		input.Output = "invalid input.\nHelp to get help"
		return errors.New("invalid input")
	}
	input.Key = res[0]
	if len(res) > 1 {
		input.Payload = res[1]
	}

	return nil

}

//HandleInput() call key function and put it output
func (input *UserInput) HandleInput() error {
	res, err := config.App.Functions[input.Key](input.Payload)
	if err != nil {
		return err
	}
	input.Output = res
	return nil
}

//ShowOutput()print out put to the user
func (input *UserInput) ShowOutput() error {
	if input.Output == "" {
		return errors.New("out put is empty")
	}
	fmt.Println(input.Output)
	return nil
}
