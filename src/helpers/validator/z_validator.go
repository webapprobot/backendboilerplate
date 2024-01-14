package validationControllers

import (
	"errors"

	// "reflect"

	// "github.com/webappbot/backendboilerplate/src/controllers/z_paths"

	"github.com/gofiber/fiber/v2"
)

var paths = make(map[string]interface{})

func Validator(c *fiber.Ctx) error {
	returnErr := false
	returnErrError := errors.New("")
	Block{
		Try: func() {
			if err, _, _, _, _, _ := getMethodConfig(c); err != nil {
				return
			}
			// _, method, methodConfig, permission, permissionUser, table := getMethodConfig(c)
			_, _, methodConfig, _, _, _ := getMethodConfig(c)
			sections := createSections()
			for _, section := range sections {
				if methodSectionDoesNotExist(methodConfig, section) {
					continue
				}
				queryConfig := make(map[string]interface{})
				processQueryAndHeader(c, section, queryConfig)
				processBody(c, section, queryConfig)

				for queryItem, parts := range queryConfig {
					valueisRequiredAndSet(c, queryItem, parts, &returnErr, &returnErrError)
					suppliedValueIsCorrectType(c, queryItem, parts, &returnErr, &returnErrError)
				}
			}
			/**
			 * Authentication and access control
			 */
			// check permissions
			// switch permissionUser {
			// case "NOBODY":
			// case "SYS_USER", "SYS_ADMIN":
			// 	var SecretKeyInterface, _ = config.GetConfig("tokensecret")
			// 	var SecretKey = SecretKeyInterface.(string)
			// 	if c.Locals("header:=:X-Authorization") == nil {

			// 		returnErr = true
			// 		// returnErrError = c.Status(401).JSON(fiber.Map{
			// 		// 	"Message": fmt.Sprintf("Missing token"),
			// 		// })
			// 		ret := make(map[string]interface{})
			// 		ret["Message"] = fmt.Sprintf("Missing token")
			// 		ret["Token"] = fmt.Sprintf("Missing token")
			// 		returnErrError = c.Status(401).JSON(ret)
			// 		Throw("")
			// 	}
			// 	tokenStr := c.Locals("header:=:X-Authorization").(string)
			// 	tokenStr = regexp.MustCompile("^Bearer ").ReplaceAllString(tokenStr, "")

			// 	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			// 		return []byte(SecretKey), nil
			// 	})
			// 	if err != nil {
			// 		returnErr = true
			// 		ret := make(map[string]interface{})
			// 		ret["Message"] = fmt.Sprintf("Invalid token supplied")
			// 		ret["Token"] = fmt.Sprintf("Invalid token supplied")
			// 		returnErrError = c.Status(401).JSON(ret)
			// 		Throw("")
			// 	}

			// 	claims := token.Claims.(*jwt.StandardClaims)

			// 	var user authModels.User
			// 	if err := db.DB.Where("user_id=?", claims.Issuer).First(&user).Error; err != nil {
			// 		returnErr = true
			// 		ret := make(map[string]interface{})
			// 		ret["Message"] = fmt.Sprintf("User no longer exists")
			// 		ret["Email"] = fmt.Sprintf("User no longer exists")
			// 		returnErrError = c.Status(401).JSON(ret)
			// 		Throw("")
			// 	}
			// 	c.Locals("auth:=:userId", user.UserId)
			// 	if user.Authority == "SYS_ADMIN" {
			// 		c.Locals("auth:=:isAdmin", true)
			// 	} else {
			// 		c.Locals("auth:=:isAdmin", false)
			// 	}
			// 	if permissionUser == "SYS_ADMIN" {
			// 		authority := user.Authority
			// 		if authority != permissionUser {
			// 			returnErr = true
			// 			// returnErrError = c.Status(401).JSON(fiber.Map{
			// 			// 	"Message": fmt.Sprintf("You are not allowed to perform a %s operation", permissionUser),
			// 			// })
			// 			ret := make(map[string]interface{})
			// 			ret["Message"] = fmt.Sprintf("You are not allowed to perform a %s operation", permissionUser)
			// 			ret["Permission"] = fmt.Sprintf("You are not allowed to perform a %s operation", permissionUser)
			// 			returnErrError = c.Status(401).JSON(ret)
			// 			Throw("")
			// 		}
			// 	}
			// }
			// if table != "" {
			// 	approvedColumnExists := controllerUtils.HasApproved(table)
			// 	if approvedColumnExists {
			// 		if allowedAfterApproved := controllerUtils.GenericCheckIfApproved(table, c); allowedAfterApproved == false {
			// 			returnErr = true
			// 			ret := make(map[string]interface{})
			// 			ret["Message"] = fmt.Sprintf("You are not allowed to perform a %s operation", permissionUser)
			// 			ret["Permission"] = fmt.Sprintf("You are not allowed to perform a %s operation", permissionUser)
			// 			returnErrError = c.Status(401).JSON(ret)
			// 			Throw("")
			// 		}
			// 		// Remove approve field if is not SYS_ADMIN
			// 		if c.Locals("auth:=:isAdmin") != nil && c.Locals("auth:=:isAdmin").(bool) == false { // is not SYS_ADMIN
			// 			c.Locals("body:=:Approved", nil)
			// 			c.Locals("body:=:approved", nil)
			// 			c.Locals("query:=:Approved", nil)
			// 			c.Locals("query:=:approved", nil)
			// 		}
			// 	}
			// }
			// if validations := methodConfig["validations"]; validations != nil {
			// 	tmp := methodConfig["validations"].(map[string]interface{})
			// 	if ensureExists := tmp["EnsureExists"]; ensureExists != nil {
			// 		if ensureExists.(bool) == true { // check if row with providedId exists
			// 			if exists, primaryKey := controllerUtils.EnsureExists(table, c); exists == false {
			// 				returnErr = true
			// 				ret := make(map[string]interface{})
			// 				ret["Message"] = fmt.Sprintf("No record found for provided %s", primaryKey)
			// 				ret[primaryKey] = ret["Message"]
			// 				returnErrError = c.Status(404).JSON(ret)
			// 				Throw("")
			// 			}
			// 		}
			// 	}
			// 	if ensureExists := tmp["EnsureExistsComplex"]; ensureExists != nil {
			// 		if ensureExists.(bool) == true { // check if row with providedId exists
			// 			if permission, ok := methodConfig["permission"]; ok {
			// 				if tablesComplexInterface, ok := permission.(map[string]interface{})["TablesComplex"]; ok {
			// 					tablesComplex := tablesComplexInterface.(string)
			// 					if exists, primaryKey := controllerUtils.EnsureExistsComplex(table, tablesComplex, c); exists == false {
			// 						returnErr = true
			// 						ret := make(map[string]interface{})
			// 						ret["Message"] = fmt.Sprintf("No record found for provided %s", primaryKey)
			// 						ret[primaryKey] = ret["Message"]
			// 						returnErrError = c.Status(404).JSON(ret)
			// 						Throw("")
			// 					}
			// 				}
			// 			}

			// 		}
			// 	}
			// 	if uniqueGroup, ok := tmp["UniqueGroup"]; ok {
			// 		if isUnique := controllerUtils.GenericUniqueGroup(c, table, uniqueGroup.(string)); !isUnique {
			// 			returnErr = true
			// 			ret := make(map[string]interface{})
			// 			ret["Message"] = fmt.Sprintf("Record already exists.")
			// 			returnErrError = c.Status(409).JSON(ret)
			// 			Throw("")
			// 		}
			// 	}
			// }
			// permission = methodConfig["permission"].(map[string]interface{})
			// if _, ok := permission["DOFORANOTHER"]; ok && permission["DOFORANOTHER"].(bool) == false {
			// 	if _, ok = permission["AccessControlTables"]; ok {
			// 		allowed, pk := controllerUtils.GenericDoForAnother(c, table, permission["AccessControlTables"].(string))
			// 		if allowed == 401 {
			// 			returnErr = true
			// 			ret := make(map[string]interface{})
			// 			ret["Message"] = fmt.Sprintf("You are not allowed to %s %s. It does not belong to you.", method, pk)
			// 			ret["Permission"] = ret["Message"]
			// 			returnErrError = c.Status(401).JSON(ret)
			// 			Throw("")
			// 		}
			// 	}
			// }
			// if _, ok := permission["DOFORANOTHERCOMPLEX"]; ok && permission["DOFORANOTHERCOMPLEX"].(bool) == false {
			// 	if _, ok = permission["AccessControlTables"]; ok {
			// 		existsInAllTables := true
			// 		if ExistsInAllTables, ok := permission["ExistsInAllTables"]; ok {
			// 			existsInAllTables = ExistsInAllTables.(bool)
			// 		}
			// 		allowed, pk := controllerUtils.GenericDoForAnotherComplex(c, table, permission["AccessControlTables"].(string), existsInAllTables)
			// 		if allowed == 401 {
			// 			returnErr = true
			// 			ret := make(map[string]interface{})
			// 			ret["Message"] = fmt.Sprintf("You are not allowed to %s %s. It does not belong to you.", method, pk)
			// 			ret["Permission"] = ret["Message"]
			// 			returnErrError = c.Status(401).JSON(ret)
			// 			Throw("")
			// 		}
			// 	}
			// }

			// //
			// if _, ok := methodConfig["patchUID"]; ok {
			// 	if methodConfig["patchUID"].(bool) == true {
			// 		c.Locals("body:=:UserId", c.Locals("auth:=:userId"))
			// 		c.Locals("query:=:UserId", c.Locals("auth:=:userId"))
			// 	}
			// }
			// if generic := methodConfig["generic"]; generic != nil {
			// 	genericVal := generic.(string)
			// 	// switch genericVal {
			// 	// case "PATCH", "POST":
			// 	// 	{
			// 	// 		if methodConfig["model"] == nil {
			// 	// 			returnErr = true
			// 	// 			ret := make(map[string]interface{})
			// 	// 			ret["Message"] = "Interval server error: MSCFG"
			// 	// 			returnErrError = c.Status(500).JSON(ret)
			// 	// 			Throw("")
			// 	// 		}
			// 	// 	}
			// 	// }
			// 	switch genericVal {
			// 	case "DELETE":
			// 		controllerUtils.GenericDelete(table, c)
			// 		returnErr = true
			// 		ret := make(map[string]interface{})
			// 		ret["Message"] = fmt.Sprintf("Record Deleted.")
			// 		returnErrError = c.Status(200).JSON(ret)
			// 		Throw("")
			// 	case "PATCH":
			// 		returnErr = true
			// 		status, updated := controllerUtils.GenericUpdate(table, c, methodConfig["preload"])
			// 		returnErrError = c.Status(status).JSON(updated)
			// 		// returnErrError = c.Status(200).JSON(updated)
			// 		Throw("")
			// 	case "GET":
			// 		updated := controllerUtils.GenericRead(table, c, methodConfig["preload"])
			// 		returnErr = true
			// 		returnErrError = c.Status(200).JSON(updated)
			// 		Throw("")
			// 	case "POST":
			// 		returnErr = true
			// 		if record, error := controllerUtils.GenericPost(table, c, methodConfig["preload"]); error == nil {
			// 			returnErrError = c.Status(200).JSON(record)
			// 		} else {
			// 			ret := make(map[string]interface{})
			// 			ret["Message"] = error
			// 			status := 422
			// 			if error.(map[string]interface{})["Status"] != nil {
			// 				status = error.(map[string]interface{})["Status"].(int)
			// 			}
			// 			if error.(map[string]interface{})["Message"] != nil {
			// 				ret = error.(map[string]interface{})
			// 			}
			// 			returnErrError = c.Status(status).JSON(ret)
			// 		}
			// 		Throw("")
			// 	}
			// }

		},
		Catch: func(e Exception) {
			// fmt.Println("Error:", e)
			// c.Next()
		},
		Finally: func() {
		},
	}.Do()
	if returnErr {
		return returnErrError
	}
	return c.Next()
	// return nil
}
