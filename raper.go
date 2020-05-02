package main

import "errors"

// Raper estructura
type Raper struct {
	ID      int
	Name    string
	Country string
	Age     int
}

// Create
func CreateRaper(r Raper) (int64, error) {
	q := `INSERT INTO 
			rapers(name, country, age) 
			VALUES($1, $2, $3);`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(r.Name, r.Country, r.Age)
	if err != nil {
		return 0, err
	}

	i, _ := res.RowsAffected()
	if i != 1 {
		return i, errors.New("error: Se esperaba 1 fila afectada")
	}

	return i, nil
}
