package authservice

import "errors"

func SignIn(email, password string) (bool, error) {
	// Replace this with actual authentication logic
	if email == "user" && password == "pass" {
		return true, nil
	}
	return false, errors.New("authentication failed")
}
