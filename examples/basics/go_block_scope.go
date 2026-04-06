package examples

import (
	"errors"
	"strings"
	"unicode"
)


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

// import (
// 	"errors"
// 	"strings"
// )

// func splitEmail(email string) (string, string, error) {
// 	username, domain, found := strings.Cut(email, "@")
// 	if !found || username == "" || domain == "" {
// 		return "", "", errors.New("invalid email format")
// 	}
// 	return username, domain, nil
// }


func splitEmail(email string) (string, string, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return "", "", errors.New("empty email")
	}

	if strings.Count(email, "@") != 1 {
		return "", "", errors.New("invalid email: must contain exactly one '@'")
	}

	parts := strings.SplitN(email, "@", 2)
	username := strings.TrimSpace(parts[0])
	domain := strings.TrimSpace(parts[1])

	if username == "" {
		return "", "", errors.New("invalid email: missing username")
	}
	if domain == "" {
		return "", "", errors.New("invalid email: missing domain")
	}

	// Перевірка на пробіли всередині username
	for _, r := range username {
		if unicode.IsSpace(r) {
			return "", "", errors.New("invalid email: username contains space")
		}
	}

	// Перевірка на пробіли всередині domain
	for _, r := range domain {
		if unicode.IsSpace(r) {
			return "", "", errors.New("invalid email: domain contains space")
		}
	}

	return username, domain, nil
}