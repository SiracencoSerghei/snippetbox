package examples

// func splitEmail(email string) (string, string) {
// 	username, domain := "", ""
// 	for i, r := range email {
// 		if r == '@' {
// 			username = email[:i]
// 			domain = email[i+1:]
// 			break
// 		}
// 	}
// 	return username, domain
// }

import (
	"errors"
	"strings"
)

func splitEmail(email string) (string, string, error) {
	username, domain, found := strings.Cut(email, "@")
	if !found || username == "" || domain == "" {
		return "", "", errors.New("invalid email format")
	}
	return username, domain, nil
}
