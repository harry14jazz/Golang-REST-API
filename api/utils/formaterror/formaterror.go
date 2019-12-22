package formaterror

import (
	"errors"
	"strings"
)

//FormatError function to create error message readable
func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("Nickname Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}

	if strings.Contains(err, "hashedPasswprd") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}
