package check

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}

func ValidateGmailAddress(email string) error {
	if !strings.Contains(email, "@") || !strings.HasSuffix(email, "@gmail.com") {
		return fmt.Errorf("Email address %s is not valid", email)
	}
	return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	starting := `^\+998\d{9}$`
	correct, err := regexp.MatchString(starting, phoneNumber)
	if err != nil {
		return fmt.Errorf("Error while validation: %v", err)
	}
	if !correct {
		return fmt.Errorf("Phone number %s is invalid", phoneNumber)
	}
	return nil
}
