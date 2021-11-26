package config

import "github.com/Ali-Farhadnia/clientGRPC/models/modelpb"

type AppConfig struct {
	GRPCconfig       GrpcConfig
	Grpc_CRUD_client modelpb.CRUDClient
	Functions        map[string]func(string) (string, error)
}
type GrpcConfig struct {
	Host string
	Port string
}

var App AppConfig
