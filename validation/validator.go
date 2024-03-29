package validation

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits or underscores -> %s", value)
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters or spaces-> %s", value)
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 200)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("must be a valid email address-> %s", value)
	}
	return nil
}

func ValidateEmailId(id int64) error {
	if id < 1 {
		return fmt.Errorf("%d : invalid email id", id)
	}
	return nil
}

func ValidateSecretCode(code string) error {
	if err := ValidateString(code, 32, 128); err != nil {
		return err
	}
	return nil
}
