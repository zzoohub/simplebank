package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/zzoopro/simple_bank/util"
)

var validCurrancy validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}