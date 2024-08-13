package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func getErrorMsg(ferry validator.FieldError) string {
	//look at the field tag
	switch ferry.Tag() {
	case "required":
		return "Perlu diisi"
	case "min":
		return "jumlah karakter kurang"
	case "max":
		return "karakter melebihi batas"
	case "email":
		return "format email salah"
	case "one-of":
		return "inputan tidak sesuai aturan"
	case "required_with":
		return "kolom ini juga harus diisi"
	case "numeric":
		return "inputan harus angka"
	}
	return "unknown error"
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorBind(err error, strType interface{}) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		res := make([]ErrorMsg, len(ve))
		tp := reflect.TypeOf(strType)
		if tp != nil {
			for i, e := range ve {
				fieldName := e.Field()
				field, ok := tp.Elem().FieldByName(fieldName)
				if !ok {
					return nil
				}
				fieldJSONName, _ := field.Tag.Lookup("json")
				res[i] = ErrorMsg{Field: fieldJSONName, Message: getErrorMsg(e)}
			}
			return res
		}
	} else {
		res := []ErrorMsg{{Field: "", Message: "Something is wrong with user's input"}}
		return res
	}
	return nil
}
