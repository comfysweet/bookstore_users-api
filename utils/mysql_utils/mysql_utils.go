package mysql_utils

import (
	"github.com/comfysweet/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRow       = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError("no record matching given id")
		}
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError("email already exists")
		}
		return errors.NewInternalServiceError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServiceError("error processing request")
}
