package db

func InitDb() error {
	db, err := getConnection()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `
		CREATE TABLE IF NOT EXISTS calculations (
			id SERIAL PRIMARY KEY,
			number1 FLOAT NOT NULL,
			number2 FLOAT NOT NULL,
			operator VARCHAR(10) NOT NULL,
			result FLOAT NOT NULL,
			request_id VARCHAR(36) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`

	_, err = db.Exec(query)
	return err
}
