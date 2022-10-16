package repositories

import (
	"database/sql"
	"fmt"
	"go-psql/config"
	"go-psql/models"
	"log"
)

func InsertUser(user models.User) int64 {
	db := config.CreateConnection()
	defer db.Close()

	query := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

	var id int64

	err := db.QueryRow(query, user.Name, user.Location, user.Age).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func GetUser(id int64) (models.User, error) {
	db := config.CreateConnection()
	defer db.Close()

	var user models.User

	query := `SELECT * FROM users WHERE userid=$1`
	row := db.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}

func GetAllUsers() ([]models.User, error) {
	db := config.CreateConnection()
	defer db.Close()

	var users []models.User

	query := `SELECT * FROM users`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		users = append(users, user)
	}

	return users, err
}

func UpdateUser(id int64, user models.User) (int64, error) {
	db := config.CreateConnection()
	defer db.Close()

	query := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

	res, err := db.Exec(query, id, user.Name, user.Location, user.Age)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected, err
}

func DeleteUser(id int64) (int64, error) {
	db := config.CreateConnection()
	defer db.Close()

	query := `DELETE FROM users WHERE userid=$1`

	res, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, err
}
