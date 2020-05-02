package models

import (
	"errors"

	"github.com/DixonOrtiz/ApiRap/database"
	_ "github.com/lib/pq"
)

//Raper structure, adapted from the database model
type Raper struct {
	ID      int
	Name    string
	Country string
	Age     int
}

//GetRapersfunction
//Function that show the rapers in the database
func GetRapers() ([]Raper, error) {
	query := `SELECT id, name, country, age FROM rapers;`
	var rapers []Raper

	db := database.GetConnection()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := Raper{}
		err = rows.Scan(&r.ID, &r.Name, &r.Country, &r.Age)
		if err != nil {
			return nil, err
		}

		rapers = append(rapers, r)
	}

	return rapers, nil
}

//CreateRaper function
//Function that introduce a new 'raper' in the database
func CreateRaper(r Raper) (int64, error) {
	query := `INSERT INTO 
			rapers(name, country, age) 
			VALUES($1, $2, $3);`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
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
