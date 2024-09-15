package models

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestWalletTransferSuccessCommit(t *testing.T) {
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
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = wallet.Transfer(gormDB, 1000)
	if err != nil {
		t.Errorf("unexpected error during transfer: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
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
	if err == nil {
		t.Errorf("expected error during transfer, but got none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
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
	if err == nil {
		t.Errorf("expected error 'Insufficient funds', but got none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
