package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	util "github.com/vipin030/automan/src/utils"
	"regexp"
	"testing"
)

// Testcreate test the creation of vehicle type
func TestCreate(t *testing.T) {
	DB, mock = Setup()
	assert := assert.New(t)
	const newID = 1
	const sql = `INSERT INTO "vehicle_types" ("name","user_id","created_at","updated_at") VALUES ($1,$2,$3,$4) RETURNING "vehicle_types"."id"`
	vtype := &VehicleType{
		Name:      "SUV",
		UserID:    1,
		CreatedAt: util.GetNow(),
		UpdatedAt: util.GetNow(),
	}

	mock.ExpectBegin() // start transaction
	mock.ExpectQuery(regexp.QuoteMeta(sql)).
		WithArgs(vtype.Name, vtype.UserID, vtype.CreatedAt, vtype.UpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newID))

	mock.ExpectCommit() // commit transaction

	resp := vtype.Create()
	assert.Equal(resp["status"].(bool), true)
}

func TestFindAll(t *testing.T) {
	DB, mock = Setup()
	assert := assert.New(t)
	const sql = `SELECT * FROM "vehicle_types"`
	vtypeMockRows := sqlmock.NewRows([]string{"id", "name", "user_id", "created_at", "updated_at"}).
		AddRow("1", "SUV 1", 1, util.GetNow(), util.GetNow())
	mock.ExpectQuery(regexp.QuoteMeta(sql)).
		WillReturnRows(vtypeMockRows)
	vtypes, err := FindAll()
	if err != nil {
		t.Fail()
	}
	assert.Equal(len(vtypes), 1)
}
