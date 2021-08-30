package repository

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type Repository struct {
	Db *sql.DB
}

func (r *Repository) GetData() error {
	var (
		usersList []user
		id        int
		name      string
		pass      string
	)
	// Запрос одной строки
	err := r.Db.QueryRow(`select id, name, pass from users where name=?`, "anton").Scan(&id, &name, &pass)
	if err == sql.ErrNoRows {
		return errors.New("not auth!")
	}
	if err != nil {
		return errors.Wrap(err, "DB error")
	}
	// Запрос нескольких строк
	rows, err := r.Db.Query(`select id, name from users`)
	if err == sql.ErrNoRows {
		return errors.New("not auth!")
	}
	if err != nil {
		return errors.Wrap(err, "DB error")
	}
	defer rows.Close()

	for rows.Next() {
		newRow := new(user)
		err = rows.Scan(&newRow.id, &newRow.name)
		if err != nil {
			return errors.Wrap(err, "Scan error")
		}
		usersList = append(usersList, *newRow)
	}
	fmt.Println(usersList)

	query := `
insert into 
user(name,pass)
values(?,?)`

	res, err := r.Db.Exec(query, "testUser", "123pass")
	if err != nil {
		return errors.Wrap(err, "DB error")
	}
	countRows, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "DB error")
	}
	if countRows < 1 {
		return errors.New("table not updated")
	}
//	res.LastInsertId()
	return nil
}

type user struct {
	id   int
	name string
}
