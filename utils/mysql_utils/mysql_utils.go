package mysql_utils

import (
	"github.com/comfysweet/bookstore_utils-go/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServiceError("error parsing database response", err)
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicate entry")
	}
	return errors.NewInternalServiceError("error processing request", errors.NewError("database error"))
}
