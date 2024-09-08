package gui

import "regexp"

func ValidateTelegramUsername(username string) bool {
	regex := `^[a-zA-Z][a-zA-Z0-9_]{4,31}$`

	re := regexp.MustCompile(regex)

	return re.MatchString(username)
}
