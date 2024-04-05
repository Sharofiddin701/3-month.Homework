package check

import (
	"errors"
	"time"
	"strings"
	"unicode"
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
	

	func ValidatePassword(password string) error {
		if len(password) < 8 {
			return errors.New("password length must be at least 8 characters")
		}
	
		var hasUppercase, hasLowercase, hasDigit, hasSymbol bool
	
		for _, char := range password {
			switch {
			case unicode.IsUpper(char):
				hasUppercase = true
			case unicode.IsLower(char):
				hasLowercase = true
			case unicode.IsNumber(char):
				hasDigit = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSymbol = true
			}
		}
	
		if !hasUppercase {
			return errors.New("password must contain at least one uppercase letter")
		}
		if !hasLowercase {
			return errors.New("password must contain at least one lowercase letter")
		}
		if !hasDigit {
			return errors.New("password must contain at least one digit")
		}
		if !hasSymbol {
			return errors.New("password must contain at least one symbol")
		}
	
		return nil
	}
