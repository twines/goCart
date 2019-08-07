/*
@author 如梦一般
@date 2019-08-07 10:53
*/
package util

import (
	"fmt"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"log"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	zh       = zh2.New()
	trans    ut.Translator
)


var codeTable = map[string]map[string]string{
	"required": {
		"Product.Stock":  "Stock",
		"Product.Price":  "Price",
		"Admin.Model.ID": "ID",
		"xxx.Model.ID":   "ID",
	},
	"gt": {

	},
}
var transTable = map[string]string{
	"Stock": "库存量不能输入0",
	"Price": "商品价格",
	"ID":    "ID必须>=0",
}

func ValidatorErrors(err error) []string {
	revs := []string{}
	if err==nil {
		return revs
	}
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	errs := err.(validator.ValidationErrors)
	if errs != nil {
		for _, e := range errs {
			tag := e.Tag()
			namespace := e.StructNamespace()
			if space, ok := codeTable[tag]; ok {
				if codeField, ok := space[namespace]; ok {
					if info, ok := transTable[codeField]; ok {
						log.Print(info)
						revs = append(revs, info)
					} else {
						panic(fmt.Sprintf("请完善%v:%v=>%v对应的中文码", tag, namespace, codeField))
					}
				} else {
					panic(fmt.Sprintf("请完善%v:%v=>%v对应的中文码", tag, namespace, codeField))
				}
			}
			//revs = append(revs, e.Translate(trans))
		}
	}
	return revs
}
func ValidatorHelper(dest interface{}, keysMap func() map[string]map[interface{}]string) (error, bool) {
	uni := ut.New(zh, zh)

	trans, _ := uni.GetTranslator("zh")
	validate = validator.New()

	zh_translations.RegisterDefaultTranslations(validate, trans)

	for key, value := range keysMap() {
		keyMap := value

		validate.RegisterTranslation(key, trans, func(ut ut.Translator) error {
			return ut.Add(key, "{0}必填", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(key, keyMap[fe.Field()])

			return t
		})
	}
	err := validate.Struct(dest)
	if err != nil {
		errs, exists := err.(validator.ValidationErrors)
		errs.Translate(trans)
		return errs, exists
	}
	return nil, false
}
