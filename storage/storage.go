package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Storage() (string, error) {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/sabaody")

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Database can't be closed")
		}
	}(db)

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	_, _ = migrate.NewWithDatabaseInstance(
		"./Migrations",
		"mysql",
		driver,
	)

	if err != nil {
		panic(err)
		return "Can't connect to database", err
	}

	// TODO заменить на человеческое логирование и в целом сделать конект к дб нормальным
	fmt.Println("Connected to MYSQL database")

	return "", nil
}
