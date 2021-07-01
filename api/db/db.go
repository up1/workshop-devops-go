package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

type DbClient interface{
	InquiryAccount(ctx context.Context, accountId string) (AccountData, error)
    CreateConnection(addr string)
	Close()
}

type PostgresClient struct{
	db *sql.DB
}

func NewPostgresClient(address string) *PostgresClient {
	pc := &PostgresClient{}
    pc.CreateConnection(address)
	return pc
}

func (pc *PostgresClient) InquiryAccount(ctx context.Context, accountId string) (AccountData, error) {
	if pc.db == nil {
		return AccountData{}, fmt.Errorf("Connection to DB not established!")
	}

	var account AccountData
	sql := "SELECT id, name FROM accounts WHERE id = $1"
	err := pc.db.QueryRow(sql, accountId).Scan(&account.ID, &account.Name)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	
	return account, nil
}

func (pc *PostgresClient) CreateConnection(addr string) {
	logrus.Infof("Connecting with connection string: '%v'", addr)
    var err error
	pc.db, err = sql.Open("postgres", addr)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	err = pc.db.Ping()
    if err != nil {
        log.Fatal("Failed to ping DB: ", err)
    }
	logrus.Info("Successfully connected to DB")
}

func (pc *PostgresClient) Close() {
	pc.db.Close()
}