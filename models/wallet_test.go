package models

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestWalletTransferSuccessCommit(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	assert.NoError(t, err)

	wallet := Wallet{
		ID:      1,
		UserID:  1,
		Balance: 5000,
	}

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT \\* FROM `wallets` WHERE `wallets`.`id` = \\? FOR UPDATE").
		WithArgs(wallet.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "balance"}).AddRow(wallet.ID, wallet.UserID, wallet.Balance))
	mock.ExpectExec("UPDATE `wallets` SET `balance`=\\?,`updated_at`=\\? WHERE `id` = \\?").
		WithArgs(wallet.Balance-1000, sqlmock.AnyArg(), wallet.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = wallet.Transfer(gormDB, 1000)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestWalletTransferFailedRollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("failed to create gorm DB connection: %v", err)
	}

	wallet := Wallet{
		ID:      1,
		UserID:  1,
		Balance: 5000,
	}

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT \\* FROM `wallets` WHERE `wallets`.`id` = \\? FOR UPDATE").
		WithArgs(wallet.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "balance"}).AddRow(wallet.ID, wallet.UserID, wallet.Balance))

	mock.ExpectExec("UPDATE `wallets` SET `balance`=\\?,`updated_at`=\\? WHERE `id` = \\?").
		WithArgs(wallet.Balance-1000, sqlmock.AnyArg(), wallet.ID).
		WillReturnError(errors.New("update failed"))

	mock.ExpectRollback()

	err = wallet.Transfer(gormDB, 1000)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestWalletTransferFailedInsufficientFunds(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("failed to create gorm DB connection: %v", err)
	}

	wallet := Wallet{
		ID:      1,
		UserID:  1,
		Balance: 500,
	}

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT \\* FROM `wallets` WHERE `wallets`.`id` = \\? FOR UPDATE").
		WithArgs(wallet.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "balance"}).AddRow(wallet.ID, wallet.UserID, wallet.Balance))

	mock.ExpectRollback()

	err = wallet.Transfer(gormDB, 1000)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
