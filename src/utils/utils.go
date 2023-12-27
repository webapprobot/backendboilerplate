package utils

import (
	"regexp"
	"strings"
)

func RemoveDoubleSpaces(str string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(str, " "))
}

func RemoveAllSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func GetForeignKeyViolationTableName(err error) string {
	re := regexp.MustCompile(`violates foreign key constraint "(.+?)"`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) == 2 {
		constraint := matches[1]
		parts := strings.Split(constraint, "_")
		return parts[len(parts)-1]
	}
	return ""
}
func GetPrimaryKeyViolationTableName(err error) string {
	re := regexp.MustCompile(`duplicate key value violates unique constraint "(.+?)"`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
