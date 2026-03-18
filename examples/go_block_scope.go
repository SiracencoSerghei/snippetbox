package examples

import (
	"errors"
	"strings"
)

func splitEmail(email string) (string, string, error) {
	email = strings.TrimSpace(email) // видаляємо пробіли спереду/ззаду

	username, domain, found := strings.Cut(email, "@")
	username = strings.TrimSpace(username)
	domain = strings.TrimSpace(domain)

	if !found || username == "" || domain == "" {
		return "", "", errors.New("invalid email format")
	}
	return username, domain, nil
}