package validation

import "errors"

func ValidateName(name string) error {
	if len(name) > 10 {
		return errors.New("invalid name length")
	}
	return nil
}

func ValidatePhone(phone string) error {
	if len(phone) != 11 {
		return errors.New("invalid phone format")
	}
	return nil
}
