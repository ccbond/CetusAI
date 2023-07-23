package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ccbond/cetus-ai/internal/config"
	"github.com/ccbond/cetus-ai/internal/server"
	openai "github.com/sashabaranov/go-openai"
)

func args_parse() {
	flag.StringVar(&config.Env, "e", "local", "Default using local environment configuration.")
	flag.Parse()
}

func handleSignals(sig os.Signal) (exitNow bool) {
	switch sig {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
		return true
	case syscall.SIGUSR1:
		return false
	}
	return false
}

func registerSignal(shutdown chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1}...)
	go func() {
		for sig := range c {
			if handleSignals(sig) {
				close(shutdown)
				return
			}
		}
	}()
}

func main() {
	shutdownChannel := make(chan struct{})
	registerSignal(shutdownChannel)
	args_parse()

	configFilePath := "./conf/" + config.Env + "/config.toml"
	config.Init(configFilePath)
	c := config.Get()
	openaiClient := openai.NewClient(c.OpenAIConfig.ApiKey)
	svcs := &server.Services{
		OpenaiClient: openaiClient,
	}
	server, err := server.NewServer(c, svcs)
	if err != nil {
		panic("Failed to build new server, err: " + err.Error())
	}

	if err := server.Run(); err != nil {
		panic("Failed to run the server, err: " + err.Error())
	}

	// wait for the terminal signal
	<-shutdownChannel
	server.Shutdown()

}
