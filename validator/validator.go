package validator

import (
	"errors"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	log "github.com/sirupsen/logrus"
)

var Validator = &MyValidator{}

type MyValidator struct {
	once     sync.Once
	validate *validator.Validate

	translator *ut.UniversalTranslator
	tran       ut.Translator
}

func (v *MyValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *MyValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		zh := zh.New()
		v.translator = ut.New(zh, zh)
		v.tran, _ = v.translator.GetTranslator("zh")

		if err := zh_translations.RegisterDefaultTranslations(v.validate, v.tran); err != nil {
			log.Error(err, "RegisterDefaultTranslations failed")
		} else {
			log.Info("RegisterDefaultTranslations success")
		}

	})
}

func (v *MyValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {

			if errs, ok := err.(validator.ValidationErrors); ok {
				sliceErrs := []string{}
				for _, e := range errs {
					sliceErrs = append(sliceErrs, e.Translate(v.tran))
				}
				return errors.New(strings.Join(sliceErrs, ","))

			}

			return err
		}
	}
	return nil
}
