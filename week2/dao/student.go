package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type Student struct {
	id      int
	Name    string
}

func (stu Student) Get() (Student, error) {
	var student Student

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {

	}
	defer db.Close()

	query := db.QueryRow("select id, name from status where id = ?", 1)
	err = query.Scan(&student.id,&student.Name)
	if err != nil {
		if err != sql.ErrNoRows {
			return student, errors.Wrap(err, "database error")
		}
	}

	return student, nil
}