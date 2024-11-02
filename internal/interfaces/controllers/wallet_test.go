package controllers

// func TestTransferMoneySuccess(t *testing.T) {
// 	GetClaimsFunc = func(c echo.Context) (*services.JWTClaims, error) {
// 		return &services.JWTClaims{
// 			ID:       1,
// 			Username: "user1",
// 		}, nil
// 	}

// 	e := echo.New()

// 	dbMock, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer dbMock.Close()

// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	assert.NoError(t, err)

// 	db.DB = gormDB

// 	requestBody := `{"username": "user2", "amount": 50}`

// 	req := httptest.NewRequest(http.MethodPost, "/transfer", strings.NewReader(requestBody))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	mock.ExpectQuery("SELECT \\* FROM `wallets` WHERE user_id = \\? ORDER BY `wallets`.`id` LIMIT \\?").
// 		WithArgs(1, 1).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "balance"}).AddRow(1, 1, 500))

// 	mock.ExpectQuery("SELECT `wallets`.`id`,`wallets`.`user_id`,`wallets`.`balance`,`wallets`.`created_at`,`wallets`.`updated_at`,`User`.`id` AS `User__id`,`User`.`username` AS `User__username`,`User`.`password` AS `User__password`,`User`.`created_at` AS `User__created_at`,`User`.`updated_at` AS `User__updated_at` FROM `wallets` LEFT JOIN `users` `User` ON `wallets`.`user_id` = `User`.`id` WHERE User.username = \\? ORDER BY `wallets`.`id` LIMIT \\?").
// 		WithArgs("user2", 1).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "balance"}).AddRow(2, 2, 500))

// 	mock.ExpectBegin()

// 	mock.ExpectQuery("SELECT \\* FROM `wallets` WHERE `wallets`.`id` = \\? FOR UPDATE").
// 		WithArgs(1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "balance"}).AddRow(1, 1, 500))

// 	mock.ExpectExec("UPDATE `wallets` SET `balance`=\\?,`updated_at`=\\? WHERE `id` = \\?").
// 		WithArgs(450.00, sqlmock.AnyArg(), 1).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("UPDATE `wallets` SET `balance`=\\?,`updated_at`=\\? WHERE `id` = \\?").
// 		WithArgs(550.00, sqlmock.AnyArg(), 2).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectCommit()

// 	if assert.NoError(t, TransferMoney(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.JSONEq(t, `{"status":"success","message":"Transfer Success"}`, rec.Body.String())
// 	}

// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

// func TestTransferMoneyUnauthorized(t *testing.T) {
// 	GetClaimsFunc = func(c echo.Context) (*services.JWTClaims, error) {
// 		return &services.JWTClaims{
// 			ID:       1,
// 			Username: "user1",
// 		}, errors.New("unauthorized")
// 	}

// 	e := echo.New()

// 	requestBody := `{"username": "user2", "amount": 50}`

// 	req := httptest.NewRequest(http.MethodPost, "/transfer", strings.NewReader(requestBody))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, TransferMoney(c)) {
// 		assert.Equal(t, http.StatusUnauthorized, rec.Code)
// 		assert.JSONEq(t, `{"status":"error","message":"Unauthorized"}`, rec.Body.String())
// 	}
// }
