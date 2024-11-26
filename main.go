package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/Dobefu/csb/cmd/migrate_db"
	"github.com/Dobefu/csb/cmd/remote_sync"

	_ "github.com/Dobefu/csb/cmd/init"
)

type subCommand struct {
	flag *flag.FlagSet
	desc string
}

var (
	verbose = flag.Bool("verbose", false, "Enable verbose logging")
	quiet   = flag.Bool("quiet", false, "Only log warnings and errors")
)

func init() {
	err := database.Connect()

	if err != nil {
		logger.Fatal("Could not connect to the database: %s", err.Error())
	}

	err = database.DB.Ping()

	if err != nil {
		logger.Fatal("Could not connect to the database: %s", err.Error())
	}
}

func main() {
	flag.Parse()
	applyGlobalFlags()

	args := flag.Args()
	cmdName := args[0]
	var err error

	flag := flag.NewFlagSet(cmdName, flag.ExitOnError)

	switch cmdName {
	case "migrate:db":
		reset := flag.Bool("reset", false, "Migrate from a clean database. Warning: this will delete existing data")
		flag.Parse(args[1:])

		err = migrate_db.Main(*reset)
		break

	case "remote:sync":
		reset := flag.Bool("reset", false, "Synchronise all data, instead of starting from the last sync token")
		flag.Parse(args[1:])

		err = remote_sync.Sync(*reset)
		break
	default:
		break
	}

	if err != nil {
		logger.Fatal(err.Error())
	}
}

func applyGlobalFlags() {
	if *verbose {
		logger.SetLogLevel(logger.LOG_VERBOSE)
	}

	if *quiet {
		logger.SetLogLevel(logger.LOG_WARNING)
	}
}

func getSubCommands() map[string]subCommand {
	return map[string]subCommand{
		"migrate:db": {
			desc: "Migrate or initialise the database",
		},
		"remote:sync": {
			desc: "Synchronise Contentstack data into the database",
		},
	}
}

func listSubCommands() {
	cmds := getSubCommands()

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	for idx, cmd := range cmds {
		fmt.Printf("  %s:\n", idx)
		fmt.Printf("    %s\n", cmd.desc)
	}

	if flag.Lookup("test.v") == nil {
		os.Exit(1)
	}
}
