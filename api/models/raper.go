package models

import (
	"errors"

	"github.com/DixonOrtiz/ApiRap/database"
)

//Raper structure, adapted from the database model
type Raper struct {
	ID      int
	Name    string
	Country string
	Age     int
}

//CreateRaper function
//Function that introduce a new 'raper' in the database
func CreateRaper(r Raper) (int64, error) {
	q := `INSERT INTO 
			rapers(name, country, age) 
			VALUES($1, $2, $3);`

	db := database.GetConnection()
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
		return i, errors.New("error: an affected row was expected")
	}

	return i, nil
}
