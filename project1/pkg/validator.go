package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTran "github.com/go-playground/validator/v10/translations/en"
	zhTran "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitTrans(locale string) (err error) { //这个包是通过validator实现参数校验和错误翻译
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New() //两个翻译器
		uni := ut.New(enT, zhT, enT)

		var ok bool
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		//注册翻译器
		switch locale {
		case "en":
			err = enTran.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTran.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTran.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}
