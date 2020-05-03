package models

import (
	"errors"
	"fmt"

	"github.com/DixonOrtiz/ApiRap/database"
	_ "github.com/lib/pq" //Driver to use postgres db
)

//Raper structure, adapted from the database model
type Raper struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Age     int    `json:"age"`
}

//GetRapers function
//Function that show the rapers in the database
func GetRapers() ([]Raper, error) {
	query := "SELECT id, name, country, age FROM rapers;"
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

//GetRaperByID function
//Function that get a raper from by ID from the database
func GetRaperByID(id int) (Raper, error) {
	query := fmt.Sprintf("SELECT id, name, country, age FROM rapers WHERE id = %d;", id)
	r := Raper{}

	db := database.GetConnection()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return r, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&r.ID, &r.Name, &r.Country, &r.Age)
		if err != nil {
			return r, err
		}
	}

	return r, nil

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

//UpdateRaper function
//Function that allows to update a raper data
func UpdateRaper(r Raper) (int64, error) {
	query := fmt.Sprintf(`UPDATE rapers 
							SET name = $1, 
								country = $2, 
								age = $3 
							WHERE id = $4;`,
	)

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(r.Name, r.Country, r.Age, r.ID)
	if err != nil {
		return 0, err
	}

	i, _ := res.RowsAffected()
	if i != 1 {
		return i, errors.New("error: an affected row was expected")
	}

	return i, nil
}
