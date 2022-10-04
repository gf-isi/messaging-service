package main

import (
	"context"
	"os"

	"github.com/coneno/logger"
	"github.com/influenzanet/messaging-service/pkg/grpc/email_client_emulator"
)

// Config is the structure that holds all global configuration data
type Config struct {
	Port                     string
	ServerConfigPath         string
	HighPrioServerConfigPath string
	EmailClientEmulatorPath  string
}

func initConfig() Config {
	conf := Config{}
	conf.Port = os.Getenv("EMAIL_CLIENT_SERVICE_LISTEN_PORT")
	conf.ServerConfigPath = os.Getenv("MESSAGING_CONFIG_FOLDER") + "/smtp-servers.yaml"
	conf.HighPrioServerConfigPath = os.Getenv("MESSAGING_CONFIG_FOLDER") + "/high-prio-smtp-servers.yaml"
	conf.EmailClientEmulatorPath = os.Getenv("EMAIL_CLIENT_EMULATOR_PATH")
	return conf
}

func main() {
	conf := initConfig()

	ctx := context.Background()
	if err := email_client_emulator.RunServer(
		ctx,
		conf.Port,
		conf.EmailClientEmulatorPath,
	); err != nil {
		logger.Error.Fatal(err)
	}
}