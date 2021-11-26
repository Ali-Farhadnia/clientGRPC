package cmd

import (
	mainfunctions "github.com/Ali-Farhadnia/clientGRPC/MainFunctions"
	"github.com/Ali-Farhadnia/clientGRPC/config"
	"github.com/Ali-Farhadnia/clientGRPC/connections"
)

func ConfigGRPC() error {
	//set grpc config
	config.App.GRPCconfig.Host = "localhost"
	config.App.GRPCconfig.Port = "50051"
	//connect to grpc server
	cli, err := connections.GetCRUDClient()
	if err != nil {
		return err
	}
	//set client accses point in config valu
	config.App.Grpc_CRUD_client = cli
	return nil
}
func ConfigMainFuncs() error {
	config.App.Functions = make(map[string]func(string) (string, error))
	//set functions
	config.App.Functions["InsertOneBook"] = mainfunctions.InsertOneBook
	config.App.Functions["InsertManyBooks"] = mainfunctions.InsertManyBooks
	config.App.Functions["UpdateBook"] = mainfunctions.UpdateBook
	config.App.Functions["DeleteBook"] = mainfunctions.DeleteBook
	config.App.Functions["FindBookByID"] = mainfunctions.FindBookByID
	config.App.Functions["Help"] = mainfunctions.Help

	return nil
}
