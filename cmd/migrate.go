package cmd

import (
	"os"
	"simple-wallet/internal/infrastructure/db"

	"github.com/labstack/gommon/log"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations: migrate < up | down >`,
	Run: func(cmd *cobra.Command, args []string) {
		migrateHandler(cmd, args)
	},
}

func init() {
	migrateCmd.SetHelpFunc(func(*cobra.Command, []string) {
		os.Stdout.WriteString(usageCommands)
	})
	rootCmd.AddCommand(migrateCmd)
}

var migrateHandler = func(cmd *cobra.Command, args []string) {
	dbConn := db.Connect()
	dbSql, _ := dbConn.DB()
	goose.SetDialect("mysql")

	migrationDir := "internal/infrastructure/db/migrations"
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
	command := args[0]
	cmdArgs := args[1:]
	if err := goose.Run(command, dbSql, migrationDir, cmdArgs...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}

var usageCommands = `
Run database migrations

Usage:
    kraken migrate [command]

Available Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version
`
