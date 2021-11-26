package connections

import (
	"fmt"

	"github.com/Ali-Farhadnia/clientGRPC/config"
	"github.com/Ali-Farhadnia/clientGRPC/models/modelpb"
	"google.golang.org/grpc"
)

//GetCRUDClient() set  modelpb.CRUDClient with client
func GetCRUDClient() (modelpb.CRUDClient, error) {
	fmt.Println("Hello i am a client")
	target := config.App.GRPCconfig.Host + ":" + config.App.GRPCconfig.Port
	cc, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	Client := modelpb.NewCRUDClient(cc)
	return Client, nil
}
