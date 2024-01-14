package validationControllers

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type typeChecker struct{}
type lenNRangeTypeChecker struct{}

func (tc typeChecker) CheckFunctions(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr, returnErrError, parts interface{}) bool {
	typeCheckerType := reflect.TypeOf(tc)
	for i := 0; i < typeCheckerType.NumMethod(); i++ {
		method := typeCheckerType.Method(i)
		if method.Type.NumIn() == 8 && method.Type.NumOut() == 1 && method.Type.Out(0).Kind() == reflect.Bool {
			funcValue := reflect.ValueOf(tc)
			funcName := method.Name
			args := []reflect.Value{
				reflect.ValueOf(queryItemType),
				reflect.ValueOf(itemToCheck),
				reflect.ValueOf(c),
				reflect.ValueOf(queryItem),
				reflect.ValueOf(returnErr),
				reflect.ValueOf(returnErrError),
				reflect.ValueOf(parts),
			}

			result := funcValue.MethodByName(funcName).Call(args)[0].Bool()
			if result {
				return true // If any function returns true, stop and return true
			}
		}
	}
	// default
	queryItemVal := itemToCheck
	c.Locals(queryItem, queryItemVal)
	return false // None of the functions returned true
}
func (tc lenNRangeTypeChecker) CheckFunctions(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr, returnErrError, parts interface{}) {
	typeCheckerType := reflect.TypeOf(tc)
	for i := 0; i < typeCheckerType.NumMethod(); i++ {
		method := typeCheckerType.Method(i)
		if method.Type.NumIn() == 8 && method.Type.NumOut() == 1 && method.Type.Out(0).Kind() == reflect.Bool {
			funcValue := reflect.ValueOf(tc)
			funcName := method.Name
			args := []reflect.Value{
				reflect.ValueOf(queryItemType),
				reflect.ValueOf(itemToCheck),
				reflect.ValueOf(c),
				reflect.ValueOf(queryItem),
				reflect.ValueOf(returnErr),
				reflect.ValueOf(returnErrError),
				reflect.ValueOf(parts),
			}
			_ = funcValue.MethodByName(funcName).Call(args)[0].Bool()
		}
	}
}

func (tc typeChecker) checkBool(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if queryItemType != "bool" {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	itemToCheck = strings.ToLower(itemToCheck)
	queryItemVal, err := strconv.ParseBool(itemToCheck)
	if err != nil {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "bool"))
		Throw("")
	}
	if backStr := strconv.FormatBool(queryItemVal); backStr != itemToCheck {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "bool"))
		Throw("")
	} else {
		c.Locals(queryItem, queryItemVal)
	}
	return true
}
func (tc typeChecker) checkInt(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if queryItemType != "int" {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	itemToCheck = strings.ToLower(itemToCheck)
	queryItemVal, err := strconv.ParseInt(itemToCheck, 10, 32)
	if err != nil {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "int"))
		Throw("")
	}
	if backStr := strconv.FormatInt(queryItemVal, 10); backStr != itemToCheck {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "int"))
		Throw("")
	} else {
		c.Locals(queryItem, queryItemVal)
	}
	return true
}
func (tc typeChecker) checkFloat(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if queryItemType != "float" {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	itemToCheck = strings.ToLower(itemToCheck)
	queryItemVal, err := strconv.ParseFloat(itemToCheck, 64)
	if err != nil {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "float"))
		Throw("")
	}
	if backStr := strconv.FormatFloat(queryItemVal, 'f', 16, 64); backStr[:len(backStr)-(len(backStr)-len(itemToCheck))] != itemToCheck {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "float"))
	} else {
		c.Locals(queryItem, queryItemVal)
	}
	return true
}
func (tc typeChecker) checkDate(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if queryItemType != "date" {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	itemToCheck = strings.ToLower(itemToCheck)
	itemToCheck = strings.Replace(itemToCheck, "-", "/", -1)
	_, err := time.Parse("2006/01/02", itemToCheck)
	if err != nil {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "date"))
		Throw("")
	} else {
		c.Locals(queryItem, itemToCheck)
	}
	return true
}
func (tc typeChecker) checkEmail(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if queryItemType != "email" {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	queryItemVal := itemToCheck
	err := isNotEmailValid(itemToCheck)
	if err {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "email"))
		Throw("")
	} else {
		c.Locals(queryItem, queryItemVal)
	}
	return true
}
func (tc typeChecker) checkTelephone(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if queryItemType != "telephone" {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	queryItemVal := itemToCheck
	err := isNotTelephoneValid(itemToCheck)
	if err {
		*returnErr = true
		*returnErrError = c.Status(422).JSON(validationErrorMsg("type", queryItemName, itemToCheck, "telephone"))
		Throw("")
	} else {
		c.Locals(queryItem, queryItemVal)
	}
	return true
}
func (tc typeChecker) checkEnum(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if queryItemType != "telephone" {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	queryItemVal := itemToCheck
	enums := parts.(map[string]interface{})["enum"].([]string)
	exists := contains(enums, itemToCheck)
	if !exists {
		*returnErr = true
		ret := make(map[string]interface{})
		ret["Message"] = fmt.Sprintf("Value of %s(%s) not in %s", queryItemName, queryItemVal, enums)
		ret[queryItemName] = fmt.Sprintf("Value of %s(%s) not in %s", queryItemName, queryItemVal, enums)
		*returnErrError = c.Status(422).JSON(ret)
		Throw("")
	} else {
		c.Locals(queryItem, queryItemVal)
	}
	return true
}

func (tc lenNRangeTypeChecker) checkLength(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if parts.(map[string]interface{})["len"] == nil {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	lenVals := parts.(map[string]interface{})["len"].(map[string]int)
	strLen := len(itemToCheck)
	max, _ := lenVals["max"]
	min, _ := lenVals["min"]
	if max >= strLen && strLen >= min {
		// c.Locals(queryItem, itemToCheck)
	} else {
		*returnErr = true
		ret := make(map[string]interface{})
		ret["Message"] = fmt.Sprintf("Value of %s(%s) should be between %d and %d characters", queryItemName, itemToCheck, min, max)
		ret[queryItemName] = fmt.Sprintf("Value of %s(%s) should be between %d and %d characters", queryItemName, itemToCheck, min, max)
		*returnErrError = c.Status(422).JSON(ret)
		Throw("")
	}
	return true
}
func (tc lenNRangeTypeChecker) checkRange(queryItemType string, itemToCheck string, c *fiber.Ctx, queryItem string, returnErr *bool, returnErrError *error, parts interface{}) bool {
	if parts.(map[string]interface{})["range"] == nil {
		return false
	}
	queryItemName, _ := getQueryItemParts(queryItem)
	rangeVals := parts.(map[string]interface{})["range"].(map[string]interface{})
	max, _ := strconv.ParseFloat(fmt.Sprintf("%v", rangeVals["max"]), 64)
	min, _ := strconv.ParseFloat(fmt.Sprintf("%v", rangeVals["min"]), 64)
	val, _ := strconv.ParseFloat(itemToCheck, 32)
	if max >= val && val >= min {

	} else {
		*returnErr = true
		ret := make(map[string]interface{})
		ret["Message"] = fmt.Sprintf("Value of %s(%s) should be between %s and %s", queryItemName, itemToCheck, regexp.MustCompile("[.]*0+$").ReplaceAllString(fmt.Sprintf("%f", min), ""), regexp.MustCompile("[.]*0+$").ReplaceAllString(fmt.Sprintf("%f", max), ""))
		ret[queryItemName] = fmt.Sprintf("Value of %s(%s) should be between %s and %s", queryItemName, itemToCheck, regexp.MustCompile("[.]*0+$").ReplaceAllString(fmt.Sprintf("%f", min), ""), regexp.MustCompile("[.]*0+$").ReplaceAllString(fmt.Sprintf("%f", max), ""))
		*returnErrError = c.Status(422).JSON(ret)
		Throw("")
	}
	return true
}
