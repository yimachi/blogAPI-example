package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// テスト全t内で共有する sql.DB型
var testDB *sql.DB
var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// 全テスト共通の前処理を書く
func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}
	return nil
}

// 全テスト共通の後処理を書く
func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}

func setupTestData() error {
	// dbuser、dbname、passは適宜変更して使う
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "dbuser", "dbname", "pass", "-e", "source ./testdata/setupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	// dbuser、dbname、passは適宜変更して使う
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "dbuser", "dbname", "pass", "-e", "source ./testdata/cleanupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
