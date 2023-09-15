package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/urfave/cli"
)

func init() {
	time.LoadLocation("UTC")             // ensure all time is in UTC
	runtime.GOMAXPROCS(runtime.NumCPU()) // set the core
}

// NewApp creates a cli app that can be setup using args flags
// gracefull shutdown using termination signals
func NewApp() *cli.App {
	conf := NewDefaultConfig()
	app := cli.NewApp()
	app.Name = Prog
	app.Usage = Description
	app.Version = Version
	app.Flags = GetCommandLineOptions()
	app.UsageText = "faucet [options]"

	app.Action = func(ctx *cli.Context) error {
		if err := ParseCLIOptions(ctx, conf); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		// Alternative Sentry Setup point where we can actually pass along configuration options
		// SetupSentry(conf)

		server, err := NewAppServer(conf)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		defer server.logger.Sync()
		if err := server.Run(); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		// Setup the termination signals
		signalChannel := make(chan os.Signal)
		signal.Notify(signalChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-signalChannel

		return nil
	}

	return app
}

func main() {
	app := NewApp()
	app.Run(os.Args)
}
