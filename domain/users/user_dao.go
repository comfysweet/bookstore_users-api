package users

import (
	"fmt"
	"github.com/comfysweet/bookstore_users-api/datasources/mysql/users_db"
	"github.com/comfysweet/bookstore_users-api/logger"
	"github.com/comfysweet/bookstore_users-api/utils/errors"
	"github.com/comfysweet/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT * FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT * FROM users where status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Err("error when trying to prepare get user statement", err)
		return errors.NewInternalServiceError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Password, &user.DateCreating, ); err != nil {
		logger.Err("error when trying to get user by id", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Err("error when trying to prepare save user statement", err)
		return errors.NewInternalServiceError("database error")
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreating, user.Status, user.Password)
	if err != nil {
		logger.Err("error when trying to save user", err)
		return mysql_utils.ParseError(err)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Err("error when trying to insert user", err)
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Err("error when trying to prepare update user statement", err)
		return errors.NewInternalServiceError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id); err != nil {
		logger.Err("error when trying to update user", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Err("error when trying to prepare delete user statement", err)
		return errors.NewInternalServiceError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Err("error when trying to delete user", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Err("error when trying to prepare find by status user statement", err)
		return nil, errors.NewInternalServiceError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Err("error when trying to get rows from users db", err)
		return nil, errors.NewInternalServiceError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Password, &user.DateCreating); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}
	return results, nil
}
