package validationControllers

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/webappbot/backendboilerplate/local-src/z_paths"
)

func isNotEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return !emailRegex.MatchString(e)
}

func isNotTelephoneValid(e string) bool {
	telephoneRegex := regexp.MustCompile(`^254[0-9]{9}$`)
	return !telephoneRegex.MatchString(e)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

/// Cconfigs
func PathConfigs() map[string]interface{} {
	for path, pathConfig := range pathsConfig {
		paths[path] = pathConfig
	}
	pkg := reflect.ValueOf(z_paths.Z_paths{})
	for i := 0; i < pkg.NumMethod(); i++ {
		method := pkg.Method(i)
		pathConfigsTmp := method.Call(nil)[0]
		i := pathConfigsTmp.Interface()
		pathConfigs := i.(map[string]interface{})
		for path, pathConfig := range pathConfigs {
			paths[path] = pathConfig
		}
	}
	return paths
}

func RoutePathConfigs(c *fiber.Ctx) map[string]interface{} {
	paths := map[string]interface{}{}
	route := c.Path()
	route = strings.TrimSuffix(route, "/")
	if _, ok := paths[route]; ok {
		pathConfig := paths[route].(map[string]interface{}) // will throw error if route is not defined
		return pathConfig
	}
	return map[string]interface{}{}
}
func MethodConfig(c *fiber.Ctx) map[string]interface{} {
	pathConfig := RoutePathConfigs(c)
	method := c.Method()
	if _, ok := pathConfig[method]; ok {
		return pathConfig[method].(map[string]interface{})
	}
	return map[string]interface{}{}
}

func getRouteAndMethod(c *fiber.Ctx) (string, string) {
	route := c.Path()
	route = strings.TrimSuffix(route, "/")
	method := c.Method()
	return route, method
}

func getMethodConfig(c *fiber.Ctx) (error, string, map[string]interface{}, map[string]interface{}, string, string) {
	var table string
	route, method := getRouteAndMethod(c)
	paths := PathConfigs()

	if paths[route] == nil {
		return errors.New("cerror"), method, nil, nil, "", table
	}
	if paths[route].(map[string]interface{}) == nil {
		return errors.New("cerror"), method, nil, nil, "", table
	}
	pathConfig := paths[route].(map[string]interface{}) // will throw error if route is not defined
	if pathConfig[method] == nil {
		return errors.New("cerror"), method, nil, nil, "", table
	}

	methodConfig := pathConfig[method].(map[string]interface{})
	permission := methodConfig["permission"].(map[string]interface{})
	permissionUser := permission["user"].(string)

	if permission["Table"] != nil {
		table = permission["Table"].(string)
	}

	return nil, method, methodConfig, permission, permissionUser, table

}

func getMethodConfigAlone(c *fiber.Ctx) map[string]interface{} {
	_, _, methodConfig, _, _, _ := getMethodConfig(c)
	return methodConfig
}

func processQueryAndHeader(c *fiber.Ctx, section string, queryConfig map[string]interface{}) {
	if section != "query" && section != "header" {
		return
	}
	methodConfig := getMethodConfigAlone(c)
	queryConfigData := methodConfig[section].(map[string]interface{})
	if section == "query" {
		c.Locals("queryConf:=:", queryConfigData)
	}
	for queryItemName, parts := range queryConfigData {
		var value interface{}
		supplied := false
		defaultVal := parts.(map[string]interface{})["default"] //.(string) // interface
		typeOf := parts.(map[string]interface{})["type"].(string)
		required := parts.(map[string]interface{})["required"].(bool)

		// // get correct type for default value
		// var defaultVal interface{}
		// switch v := val.(type) {
		// case int:
		// 	defaultVal = v
		// case string:
		// 	defaultVal = v
		// case float64:
		// 	defaultVal = v
		// case float32:
		// 	defaultVal = v
		// case bool:
		// 	defaultVal = v
		// default:
		// 	defaultVal = fmt.Sprintf("%v", v)
		// }

		// set value if it exists
		// get the correct type
		var itemVal interface{} // condense
		if section == "header" {
			itemVal = string(c.Request().Header.Peek(queryItemName))
		} else {
			itemVal = c.Query(queryItemName)
		}

		// if len(itemVal) > 0 { // set supplied
		// 	value = itemVal
		// 	supplied = true
		// } else {
		// 	if len(defaultVal) > 0 {
		// 		value = defaultVal
		// 		supplied = true
		// 	}
		// }
		if len(itemVal.(string)) > 0 { // set supplied
			value = itemVal
			supplied = true
		} else {
			if len(defaultVal.(string)) > 0 {
				value = defaultVal
				supplied = true
			}
		}

		queryConfig[section+":=:"+queryItemName] = map[string]interface{}{
			"type":     typeOf,
			"required": required,
			"default":  defaultVal,
			"value":    value,
			"supplied": supplied,
		}

	}

}

func processBody(c *fiber.Ctx, section string, queryConfig map[string]interface{}) {
	// if methodConfig["body"] == nil || len(methodConfig["body"].(map[string]interface{})) == 0 {
	// 	break
	// }
	if section != "body" {
		return
	}
	methodConfig := getMethodConfigAlone(c)
	bodyFields := methodConfig["body"].(map[string]interface{})["fields"]
	bodyFieldSettings := methodConfig["body"].(map[string]interface{})["fieldSettings"].(map[string]interface{})
	// bodyFieldsOptional := methodConfig["body"].(map[string]interface{})["optionalFields"]
	// bodyFieldsDefaultValues := methodConfig["body"].(map[string]interface{})["defaultValues"].(map[string]interface{})
	fields := reflect.VisibleFields(reflect.TypeOf(bodyFields))
	// fieldsOptional := reflect.VisibleFields(reflect.TypeOf(bodyFieldsOptional))
	c.Locals("bodyConf:=:", fields)
	// get all fields,
	// get optional fields
	// get default values
	if err := c.BodyParser(&bodyFields); err != nil {
	}
	for _, field := range fields {
		required := true
		var value interface{}
		supplied := false
		var defaultVal interface{}
		fieldType := fmt.Sprintf("%v", field.Type)
		enum := make([]interface{}, 0)
		var length interface{}
		var range_ interface{}
		// set required
		if required_ := bodyFieldSettings[field.Name]; required_ != nil {
			if bodyFieldSettings[field.Name].(map[string]interface{})["required"] != nil {
				required = bodyFieldSettings[field.Name].(map[string]interface{})["required"].(bool)
			}
			if bodyFieldSettings[field.Name].(map[string]interface{})["default"] != nil {
				defaultVal = fmt.Sprintf("%v", bodyFieldSettings[field.Name].(map[string]interface{})["default"]) // interface{}
				// defaultVal = bodyFieldSettings[field.Name].(map[string]interface{})["default"].(string)
			}
			if bodyFieldSettings[field.Name].(map[string]interface{})["type"] != nil {
				fieldType = bodyFieldSettings[field.Name].(map[string]interface{})["type"].(string)
			}
			if bodyFieldSettings[field.Name].(map[string]interface{})["enum"] != nil {
				enum = bodyFieldSettings[field.Name].(map[string]interface{})["enum"].([]interface{})
			}
			if bodyFieldSettings[field.Name].(map[string]interface{})["len"] != nil {
				length = bodyFieldSettings[field.Name].(map[string]interface{})["len"]
			}
			if bodyFieldSettings[field.Name].(map[string]interface{})["range"] != nil {
				range_ = bodyFieldSettings[field.Name].(map[string]interface{})["range"]
			}
		}

		if strings.ToLower(field.Name) == "email" {
			fieldType = "email"
		}
		tmpFields := bodyFields.(map[string]interface{})
		if tmpFields[field.Name] != nil {
			// if len(tmpFields[field.Name].(string)) > 0 {
			if len(fmt.Sprintf("%v", tmpFields[field.Name])) > 0 {
				value = tmpFields[field.Name]
				// value = fmt.Sprintf("%v", tmpFields[field.Name])
				supplied = true
			} else {
				// if len(defaultVal) > 0 {
				if len(fmt.Sprintf("%v", defaultVal)) > 0 {
					value = defaultVal
					// value = fmt.Sprintf("%v", defaultVal)
					supplied = true
				}
			}
		} else {
			if len(fmt.Sprintf("%v", defaultVal)) > 0 {
				value = defaultVal
				supplied = true
			}
		}
		queryConfig["body:=:"+field.Name] = map[string]interface{}{
			"type":     fieldType,
			"required": required,
			"default":  defaultVal,
			"value":    value,
			"supplied": supplied,
			"enum":     enum,
			"len":      length,
			"range":    range_,
		}
	}
}

func valueisRequiredAndSet(c *fiber.Ctx, queryItem string, parts interface{}, returnErr *bool, returnErrError *error) {
	queryItemName, queryItemPartName := getQueryItemParts(queryItem)
	if required := parts.(map[string]interface{})["required"].(bool); required {
		if supplied := parts.(map[string]interface{})["supplied"].(bool); !supplied {
			*returnErr = true
			status := 422
			if queryItemPartName == "header" {
				status = 401
			}
			ret := make(map[string]interface{})
			ret["Message"] = fmt.Sprintf("%s is missing in %s.", queryItemName, queryItemPartName)
			ret[queryItemName] = fmt.Sprintf("%s is missing in %s.", queryItemName, queryItemPartName)
			*returnErrError = c.Status(status).JSON(ret)
			Throw("")
		}
	}
}

func validationErrorMsg(errorType string, field string, value interface{}, fieldType string) map[string]interface{} {
	ret := make(map[string]interface{})
	switch errorType {
	case "type":
		ret["Message"] = fmt.Sprintf("Wrong value(%s) given for %s of type %s", value, field, fieldType)
		ret[field] = fmt.Sprintf("Wrong value(%s) given for %s of type %s", value, field, fieldType)
		break
	case "length":

	}
	return ret
}
func suppliedValueIsCorrectType(c *fiber.Ctx, queryItem string, parts interface{}, returnErr *bool, returnErrError *error) {
	// run only if supplied == true
	if supplied := parts.(map[string]interface{})["supplied"].(bool); !supplied {
		return
	}
	queryItemType := parts.(map[string]interface{})["type"].(string)
	itemToCheck := parts.(map[string]interface{})["value"].(string)
	itemToCheck = strings.TrimSpace(itemToCheck)
	{ // data type
		typeChecker := typeChecker{}
		_ = typeChecker.CheckFunctions(queryItemType, itemToCheck, c, queryItem, returnErr, returnErrError, parts)
	}
	{ // length/range etc
		typeChecker := lenNRangeTypeChecker{}
		typeChecker.CheckFunctions(queryItemType, itemToCheck, c, queryItem, returnErr, returnErrError, parts)
	}
}
