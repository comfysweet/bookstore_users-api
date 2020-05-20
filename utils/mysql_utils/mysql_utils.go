package mysql_utils

import (
	"github.com/comfysweet/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRow       = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServiceError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicate entry")
	}
	return errors.NewInternalServiceError("error processing request")
}
