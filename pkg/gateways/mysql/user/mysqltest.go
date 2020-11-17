package user

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
}

func NewMySQLTest() (*sql.DB, error) {
	config, err := testDBConfig()
	if err != nil {
		return nil, err
	}
	return NewMySQL(config)
}

func testDBConfig() (Config, error) {
	username, err := user.Current()
	if err != nil {
		return Config{}, err
	}
	return Config{
		Username: username.Username,
		Password: "",
		Host:     "localhost",
		Port:     "3306",
	}, nil
}

func Setup(db *sql.DB, filename string) error {
	return ExecuteSQL(db, filename)
}

func Teardown(db *sql.DB, filename string) error {
	return ExecuteSQL(db, filename)
}

func ExecuteSQL(db *sql.DB, filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	stmts := strings.Split(string(file), ";\n")
	for _, stmt := range stmts {
		if len(stmt) == 0 {
			continue
		}
		if _, err := db.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

func NewMySQL(config Config) (*sql.DB, error) {
	return initConnection(makeConnectionString(config))
}

func initConnection(connectionString string) (*sql.DB, error) {
	db, dbConnErr := sql.Open("mysql", connectionString)
	if dbConnErr != nil {
		return nil, fmt.Errorf("Failed to create database connection: %w", dbConnErr)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(time.Second * 10)

	return db, nil
}

func makeConnectionString(c Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?parseTime=true",
		c.Username, c.Password, c.Host, c.Port,
	)
}
