package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

var email, password string
var mock sqlmock.Sqlmock

func Setup() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	DB, err = gorm.Open("postgres", db)
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	DB.LogMode(true)
	return DB, mock
}

func TestUserCreate(t *testing.T) {
	email = uuid.New().String() + "@gmail.com"
	password = "password"
	user := &User{
		Email:    email,
		Password: password,
		Phone:    "9895736375",
	}
	DB, mock = Setup()

	const newId = 1
	const sql = `INSERT INTO "users" ("email","password","phone") VALUES ($1,$2,$3) RETURNING "users"."id"`
	const sqlSelectAll = `SELECT * FROM "users"  WHERE (email = $1) ORDER BY "users"."id" ASC LIMIT 1`

	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).
		WillReturnRows(sqlmock.NewRows(nil))
	mock.ExpectBegin() // start transaction
	mock.ExpectQuery(regexp.QuoteMeta(sql)).
		WithArgs(user.Email, user.Password, user.Phone).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(newId))

	mock.ExpectCommit() // commit transaction
	GeneratePass = func(password string) (string, error) {
		return "password", nil
	}

	resp, err := user.Create()
	if err != nil {
		t.Fail()
	}
	assert := assert.New(t)
	if status := resp["status"].(bool); !status {
		t.Fail()
	}
	assert.Equal(resp["status"].(bool), true)
}
