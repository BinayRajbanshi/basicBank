package api

import (
	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupporteCurrency(currency)
	}
	return false
}
