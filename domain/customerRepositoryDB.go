package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pilseong/banking/errs"
	"github.com/pilseong/banking/logger"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	customers := make([]Customer, 0)
	err := d.db.Select(&customers, findAllSql)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customers not found")
		} else {
			logger.Error("Error while querying DB " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB error")
		}
	}
	return customers, nil
}

func (d CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	findOneSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.db.Get(&c, findOneSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	db_user := os.Getenv("DB_USER")
	db_passwd := os.Getenv("DB_PASSWD")
	db_addr := os.Getenv("DB_ADDR")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_passwd, db_addr, db_port, db_name)
	db, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDB{db}
}
