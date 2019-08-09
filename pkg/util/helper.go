package util

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zoroastrians "gopkg.in/go-playground/validator.v9/translations/zh"
	"strings"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func init() {
	translator := zh.New()
	uni = ut.New(translator, translator)
	validate = validator.New()
}

/**
验证struct
*/
func Validator(data interface{}, keyMap map[string]string) (errMap map[string]string, ok bool) {
	trans, _ := uni.GetTranslator("zh")
	_ = zoroastrians.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(data)
	errMap = map[string]string{}
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if _, v := keyMap[e.Field()]; v == true {
				errMap[e.StructField()] = strings.Replace(e.Translate(trans), e.Field(), keyMap[e.Field()], 1)
			} else {
				errMap[e.StructField()] = e.Translate(trans)
			}
		}
		return errMap, false
	} else {
		return errMap, true
	}
}
