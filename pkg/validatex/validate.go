package validatex

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// Validator 全局单例
var validate *validator.Validate

func init() {
	validate = validator.New()

	err := validate.RegisterValidation("notzero", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		if field.Kind() == reflect.Ptr && !field.IsNil() {
			if v, ok := field.Interface().(*int64); ok {
				return *v != 0
			}
		}
		return false
	})
	if err != nil {
		panic(err)
		return
	}
}

func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("字段 %s 验证失败，规则：%s", err.Field(), err.Tag())
		}
	}
	return nil
}
