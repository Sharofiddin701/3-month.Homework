package check

import (
	"errors"
	"time"
	"strings"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}
func ValidateEmail(gmail string)error{
atCount:=strings.Count(gmail,"@")
dotCount:=strings.Count(gmail,".")
if atCount==1&&dotCount==1&&(strings.HasSuffix(gmail,"@gmail.com")||strings.HasSuffix(gmail,"@mail.ru") ){return nil}else{
	return errors.New("gmail is not validet")
}
}

func ValidatePhone(phone string)error{
	plusCount:=strings.Count(phone,"+")
	if plusCount==1&&strings.HasPrefix(phone,"+998")&&len(phone)==13{
		return nil
	}else{
		return errors.New("number is not validet")
	}
	}
	
