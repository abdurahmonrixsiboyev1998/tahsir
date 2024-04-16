package validateuser

import (
	"errors"
	"strings"
)


func ValidateUser(Name string,Age int, Email string) []error{
	var err []error

	if len(Name) < 6 {
        err = append(err, errors.New("Name: length cannot be less than a 6 characters"))
    }
	if Age <= 0 || Age > 120 {
        err = append(err, errors.New("Age: couldn't be more than 120"))
    }

	if Email == "" {
		err = append(err, errors.New("Email: cannot be empty"))
	} else {
		if !strings.Contains(Email, "@") || !strings.Contains(Email, ".") {
			err = append(err, errors.New("Email: invalid format"))
		}
	}

	return err
}
