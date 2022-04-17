package main

import (
	_ "github.com/laijunbin/go-migrate-seed-example/migrations"
	"github.com/laijunbin/go-migrate/cmd"
	"github.com/laijunbin/go-migrate/config"
	"github.com/laijunbin/go-migrate/pkg/lib/mysql"
)

func init() {
	config.Config = config.DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "",
		Dbname:   "test",
	}

	config.Migrator = mysql.InitMigrator()
	config.Driver = "mysql"
}

func main() {
	cmd.Execute()
}
