package main

import (
	"log"
	"os"

	"whooshcat/checkup"
	"whooshcat/cli"
	"whooshcat/fleet"
	//"whooshcat/rules"
)

func main() {
	inspector := checkup.NewRealInspector(2) //минимальная простота для приемки
	f := fleet.NewFleet(inspector)
	//policy := rules.SimplicityThreshold{Min: 6}

	commands := []cli.Command{
		cli.NewAcceptTransportCommand(f),
		cli.NewExitCommand(),
	}

	app := cli.NewApp(commands, os.Stdin, os.Stdout)
	if err := app.Run(); err != nil {
		log.Fatalf("CLI: %v", err)
	}
}
