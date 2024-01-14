package validationControllers

import (
	"strings"
)

func createSections() []string {
	sections := make([]string, 0)
	sections = append(sections, "query")
	sections = append(sections, "body")
	sections = append(sections, "header")
	return sections
}

func methodSectionDoesNotExist(methodConfig map[string]interface{}, section string) bool {
	if methodConfig[section] == nil || len(methodConfig[section].(map[string]interface{})) == 0 {
		return true
	}
	return false
}

func getQueryItemParts(queryItem string) (string, string) {
	tmpParts := strings.Split(queryItem, ":=:")
	queryItemName := tmpParts[1]
	queryItemPartName := tmpParts[0]
	return queryItemName, queryItemPartName
}

func getFieldFromSectionData(sectionData map[string]interface{}, dataType string) {

}
