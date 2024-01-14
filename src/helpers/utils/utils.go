package controllerUtils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	inflection "github.com/jinzhu/inflection"
	strcase "github.com/stoewer/go-strcase"
	"github.com/webappbot/backendboilerplate/local-src/db"
	utils "github.com/webappbot/backendboilerplate/src/utils"
	"gorm.io/gorm/schema"
)

func ToGormFieldName(name string) string {
	// s := schema.NamingStrategy{}
	// return s.ColumnName("", strings.ToLower(name))
	return gorm.ToDBName(name)
}
func ToGormTableName(name string) string {
	s := schema.NamingStrategy{
		SingularTable: false,
	}
	return s.TableName(strings.ToLower(name))
}

func GetTableColumns(table string) map[string]string {
	query := "SELECT column_name, data_type FROM information_schema.columns WHERE table_name = ?"
	rows, err := db.DB.Raw(query, table).Rows()
	defer rows.Close()

	if err != nil {
		// Handle error
	}

	// Define a map to store column names and types
	columns := make(map[string]string)

	// Iterate through the results to build the columns map
	for rows.Next() {
		var columnName, dataType string
		rows.Scan(&columnName, &dataType)
		columns[columnName] = dataType
	}
	return columns
}

func HasApproved(table string) bool {
	columns := GetTableColumns(table)
	approvedColumnExists := false
	for columnName, dataType := range columns {
		if columnName == "approved" && dataType == "boolean" {
			approvedColumnExists = true
			break
		}
	}
	return approvedColumnExists
}

func CheckPrimaryKey(tableName string) string {
	tableName = ToGormTableName(strcase.SnakeCase(tableName))
	query := `
    SELECT column_name
    FROM information_schema.constraint_column_usage
    WHERE constraint_name = (
        SELECT constraint_name
        FROM information_schema.table_constraints
        WHERE table_name = $1
            AND constraint_type = 'PRIMARY KEY'
    );
`
	rows, err := db.DB.Raw(query, tableName).Rows()
	if err != nil {
		// Handle error
	}
	defer rows.Close()

	var primaryKeyColumns []string
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			// Handle error
		}
		primaryKeyColumns = append(primaryKeyColumns, columnName)
	}

	if err := rows.Err(); err != nil {
		// Handle error
	}
	return primaryKeyColumns[0]
}

// / return true if conditions are met:
// either not  approved
// or approved && SYS_ADMIN
func GenericCheckIfApproved(table string, c *fiber.Ctx) bool {
	hasPrimaryKey, primaryKey, primaryKeyVal := PrimaryKeyData(table, c)
	if !hasPrimaryKey {
		return true
	}
	var results map[string]interface{}
	result := db.DB.Table(table).Find(&results, fmt.Sprintf("%s = ?", primaryKey), fmt.Sprintf("%s", primaryKeyVal))
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	// b, _ := json.MarshalIndent(results, "", "    ")
	b, _ := json.Marshal(results)

	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		// panic(err)
		return true
	}
	if f == nil {
		return true
	}
	approved := f.(map[string]interface{})["approved"].(bool)
	if !approved {
		return true
	}
	if c.Locals("auth:=:isAdmin").(bool) == false { // is not SYS_ADMIN
		return false
	}
	return false
}

// / return true row with provided primaryKey val exists
func EnsureExistsComplex(table string, tablesComplex string, c *fiber.Ctx) (bool, string) {
	complexTables := strings.Split(strings.ReplaceAll(tablesComplex, " ", ""), ",")
	existed := true
	for _, table := range complexTables {
		exists, pk := EnsureExists(table, c)
		if exists == false {
			return exists, pk
		}
	}
	return existed, ""
	// hasPrimaryKey, primaryKey, primaryKeyVal := PrimaryKeyData(table, c)
	// if !hasPrimaryKey {
	// 	return false, primaryKey
	// }
	// var results map[string]interface{}
	// result := db.DB.Table(table).Find(&results, fmt.Sprintf("%s = ?", primaryKey), fmt.Sprintf("%s", primaryKeyVal))
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// // b, _ := json.MarshalIndent(results, "", "    ")
	// b, _ := json.Marshal(results)

	// var f interface{}
	// if err := json.Unmarshal(b, &f); err != nil {
	// 	// panic(err)
	// 	return false, primaryKey
	// }
	// if f == nil {
	// 	return false, primaryKey
	// }
	// return true, primaryKey
}
func EnsureExists(table string, c *fiber.Ctx) (bool, string) {
	hasPrimaryKey, primaryKey, primaryKeyVal := PrimaryKeyData(table, c)
	if !hasPrimaryKey {
		return false, primaryKey
	}
	var results map[string]interface{}
	result := db.DB.Table(table).Find(&results, fmt.Sprintf("%s = ?", primaryKey), fmt.Sprintf("%s", primaryKeyVal))
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	// b, _ := json.MarshalIndent(results, "", "    ")
	b, _ := json.Marshal(results)

	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		// panic(err)
		return false, primaryKey
	}
	if f == nil {
		return false, primaryKey
	}
	return true, primaryKey
}

func PrimaryKeyData(table string, c *fiber.Ctx) (bool, string, string) {
	primaryKey := CheckPrimaryKey(table)
	primaryKeyVal := ""
	hasPrimaryKey := false
	// loop through locals to check the primary key. start with body, then query
	if tmp := c.Locals("bodyConf:=:"); tmp != nil {
		fields := tmp.([]reflect.StructField)
		for _, field := range fields {
			fieldName := field.Name
			gormFieldName := ToGormFieldName(fieldName)
			if gormFieldName == primaryKey && c.Locals(fmt.Sprintf("body:=:%s", fieldName)) != nil {
				hasPrimaryKey = true
				primaryKeyVal = c.Locals(fmt.Sprintf("body:=:%s", fieldName)).(string) // check: dataType bad assumption
				c.Locals("field:=:PK", fmt.Sprintf("body:=:%s", fieldName))
			}

			// fmt.Println(fmt.Sprintf("%v:%v:%v:=>%s", fieldName, gormFieldName, primaryKey, primaryKeyVal))
		}
	}
	if !hasPrimaryKey { // check inside query
		if tmp := c.Locals("queryConf:=:"); tmp != nil {
			fields := tmp.(map[string]interface{})
			for fieldName, _ := range fields {
				gormFieldName := ToGormFieldName(fieldName)
				if gormFieldName == primaryKey && c.Locals(fmt.Sprintf("query:=:%s", fieldName)) != nil {
					hasPrimaryKey = true
					primaryKeyVal = c.Locals(fmt.Sprintf("query:=:%s", fieldName)).(string) // check: dataType bad assumption
					c.Locals("field:=:PK", fmt.Sprintf("query:=:%s", fieldName))
				}
				// fmt.Println(fmt.Sprintf("%v:%v:%v:=>%s", fieldName, gormFieldName, primaryKey, primaryKeyVal))
			}
		}
	}
	return hasPrimaryKey, primaryKey, primaryKeyVal
}

// / return true row with provided primaryKey val exists
func GenericDelete(table string, c *fiber.Ctx) {
	_, primaryKey, primaryKeyVal := PrimaryKeyData(table, c)

	var results map[string]interface{}
	result := db.DB.Table(table).Delete(&results, fmt.Sprintf("%s = ?", primaryKey), fmt.Sprintf("%s", primaryKeyVal))
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

// convert fields having _lower to lowercase with whitespaces removed
// remove whitespaces from corresponding *_lower fieds
func toLowerCase(data map[string]interface{}) map[string]interface{} {
	for key, val := range data {
		if strings.HasSuffix(key, "_lower") {
			data[key] = strings.ToLower(utils.RemoveAllSpaces(val.(string)))
			field := strings.TrimSuffix(key, "_lower")
			data[field] = utils.RemoveDoubleSpaces(data[field].(string))
		}
	}
	return data
}

// Generic POST
func GenericPost(table string, c *fiber.Ctx, preload interface{}) (map[string]interface{}, interface{}) {
	// _, primaryKey, primaryKeyVal := PrimaryKeyData(table, c)
	primaryKey := CheckPrimaryKey(table)
	data := map[string]interface{}{}
	data[primaryKey] = uuid.Must(uuid.NewRandom()).String()
	fieldMappings := map[string]interface{}{}
	columns := GetTableColumns(table)
	if tmp := c.Locals("bodyConf:=:"); tmp != nil {
		fields := tmp.([]reflect.StructField)
		for _, field := range fields {
			fieldName := field.Name
			fieldValInterface := c.Locals(fmt.Sprintf("body:=:%s", fieldName))
			gormFieldName := ToGormFieldName(fieldName)
			if fieldValInterface != nil {
				fieldVal := fieldValInterface
				if _, ok := columns[fmt.Sprintf("%s_lower", gormFieldName)]; ok {
					data[fmt.Sprintf("%s_lower", gormFieldName)] = fieldVal
				}
				// else {
				data[gormFieldName] = fieldVal
				// }
			}
			fieldMappings[gormFieldName] = fieldName
		}
	}
	if _, ok := columns["created_at"]; ok {
		data["created_at"] = time.Now()
	}
	if _, ok := columns["updated_at"]; ok {
		data["updated_at"] = time.Now()
	}
	data = toLowerCase(data)
	result := db.DB.Table(table).Create(&data)
	if result.Error != nil {

		if foreignKeyErr := utils.GetForeignKeyViolationTableName(result.Error); foreignKeyErr != "" {
			ret := make(map[string]interface{})
			ret["Message"] = fmt.Sprintf("%v does not exist.", foreignKeyErr)
			ret["Status"] = 404
			ret[foreignKeyErr] = ret["Message"]
			return nil, ret
		}
		if foreignKeyErr := utils.GetPrimaryKeyViolationTableName(result.Error); foreignKeyErr != "" {
			ret := make(map[string]interface{})
			ret["Message"] = fmt.Sprintf("Record already exists.")
			ret["Status"] = 404
			return nil, ret
		}
		return nil, result.Error
	}
	primaryKeyVal := data[primaryKey]
	db.DB.Table(table).First(&data, fmt.Sprintf("%s = ?", primaryKey), fmt.Sprintf("%s", primaryKeyVal))
	if preload != nil {
		dataRet := map[string]interface{}{}
		for key, val := range data {
			if fieldMappings[key] != nil {
				dataRet[fieldMappings[key].(string)] = val
			} else {
				dataRet[strcase.UpperCamelCase(key)] = val
			}
		}
		data = LoadAssociation(c, preload.(string), data, dataRet)
	} else {
		dataRet := map[string]interface{}{}
		for key, val := range data {
			if fieldMappings[key] != nil {
				dataRet[fieldMappings[key].(string)] = val
			} else {
				dataRet[strcase.UpperCamelCase(key)] = val
			}
		}
		data = dataRet
	}
	return data, nil
}

func GenericUpdate(table string, c *fiber.Ctx, preload interface{}) (int, map[string]interface{}) {
	_, primaryKey, primaryKeyVal := PrimaryKeyData(table, c)
	data := map[string]interface{}{}
	fieldMappings := map[string]interface{}{}
	db.DB.Table(table).First(&data, fmt.Sprintf("%s = ?", primaryKey), fmt.Sprintf("%s", primaryKeyVal))
	columns := GetTableColumns(table)
	if tmp := c.Locals("bodyConf:=:"); tmp != nil {
		fields := tmp.([]reflect.StructField)
		for _, field := range fields {
			fieldName := field.Name
			fieldValInterface := c.Locals(fmt.Sprintf("body:=:%s", fieldName))
			gormFieldName := ToGormFieldName(fieldName)
			if fieldValInterface != nil {
				fieldVal := fieldValInterface
				// data[gormFieldName] = fieldVal
				if _, ok := columns[fmt.Sprintf("%s_lower", gormFieldName)]; ok {
					data[fmt.Sprintf("%s_lower", gormFieldName)] = fieldVal
				}
				// else {
				data[gormFieldName] = fieldVal
				// }
			}
			fieldMappings[gormFieldName] = fieldName
		}
	}
	if _, ok := columns["updated_at"]; ok {
		data["updated_at"] = time.Now()
	}

	data = toLowerCase(data)
	result := db.DB.Table(table).Where(fmt.Sprintf("%s = ?", primaryKey), fmt.Sprintf("%s", primaryKeyVal)).Updates(&data)
	if result.Error != nil {
		if foreignKeyErr := utils.GetForeignKeyViolationTableName(result.Error); foreignKeyErr != "" {
			ret := make(map[string]interface{})
			ret["Message"] = fmt.Sprintf("%v does not exist.", foreignKeyErr)
			ret[foreignKeyErr] = ret["Message"]
			return 404, ret
		}
		if foreignKeyErr := utils.GetPrimaryKeyViolationTableName(result.Error); foreignKeyErr != "" {
			ret := make(map[string]interface{})
			ret["Message"] = fmt.Sprintf("Record already exists.")
			return 401, ret
		}
	}
	result = db.DB.Table(table).Find(&data, fmt.Sprintf("%s = ?", primaryKey), fmt.Sprintf("%s", primaryKeyVal))
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if preload != nil {
		dataRet := map[string]interface{}{}
		for key, val := range data {
			if fieldMappings[key] != nil {
				dataRet[fieldMappings[key].(string)] = val
			} else {
				// dataRet[key] = val
				dataRet[strcase.UpperCamelCase(key)] = val
			}
		}
		data = LoadAssociation(c, preload.(string), data, dataRet)
	} else {
		dataRet := map[string]interface{}{}
		for key, val := range data {
			if fieldMappings[key] != nil {
				dataRet[fieldMappings[key].(string)] = val
			} else {
				// dataRet[key] = val
				dataRet[strcase.UpperCamelCase(key)] = val
			}
		}
		data = dataRet
	}
	return 200, data
}

func GenericRead(table string, c *fiber.Ctx, preload interface{}) []map[string]interface{} {
	fieldMappings := map[string]interface{}{}
	fields := map[string]interface{}{}
	if tmp := c.Locals("queryConf:=:"); tmp != nil {
		fields_ := tmp.(map[string]interface{})
		for fieldName, _ := range fields_ {
			gormFieldName := ToGormFieldName(fieldName)
			if c.Locals(fmt.Sprintf("query:=:%s", fieldName)) != nil {
				fields[gormFieldName] = c.Locals(fmt.Sprintf("query:=:%s", fieldName))
			}

			fieldMappings[gormFieldName] = fieldName
			// fmt.Println(fmt.Sprintf("%v:%v:%v:=>%s", fieldName, gormFieldName, primaryKey, primaryKeyVal))
		}
	}
	results, err := FindResultsByFields(table, fields)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	if results == nil {
		results = []map[string]interface{}{}
	}
	if preload != nil {
		for index, result := range results {
			dataRet := map[string]interface{}{}
			for key, val := range result {
				if fieldMappings[key] != nil {
					dataRet[fieldMappings[key].(string)] = val
				} else {
					dataRet[strcase.UpperCamelCase(key)] = val
				}
			}
			results[index] = LoadAssociation(c, preload.(string), result, dataRet)
		}
	} else {
		resultsRet := []map[string]interface{}{}
		// fmt.Println(reflect.TypeOf(results).Kind() == reflect.Slice)
		for _, result := range results {
			dataRet := map[string]interface{}{}
			for key, val := range result {
				if fieldMappings[key] != nil {
					dataRet[fieldMappings[key].(string)] = val
				} else {
					dataRet[strcase.UpperCamelCase(key)] = val
				}
			}
			resultsRet = append(resultsRet, dataRet)
		}
		results = resultsRet
	}

	return results
}

func FindResultsByFields(table string, fields map[string]interface{}) ([]map[string]interface{}, error) {
	query := db.DB.Table(table)
	for field, value := range fields {
		query = query.Where(fmt.Sprintf("%s = ?", gorm.ToColumnName(field)), value)
	}
	var results []map[string]interface{}
	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
func FindFirstResultByFields(table string, fields map[string]interface{}, orderBy string) ([]map[string]interface{}, error) {
	query := db.DB.Table(table)
	for field, value := range fields {
		query = query.Where(fmt.Sprintf("%s = ?", gorm.ToColumnName(field)), value)
	}
	var result map[string]interface{}
	var results []map[string]interface{}
	if orderBy == "" {
		orderBy = "created_at desc"
	}
	if err := query.Order(orderBy).Find(&result).Error; err != nil {
		// if err := query.First(&results).Error; err != nil {
		return nil, err

	} else {
		if len(result) > 0 {
			results = append(results, result)
		}

	}

	return results, nil
}

func LoadAssociation(c *fiber.Ctx, tablesStr string, tmpData map[string]interface{}, tmpDataRet map[string]interface{}) map[string]interface{} {
	tablesStr = strings.ReplaceAll(tablesStr, " ", "")
	tables := strings.Split(tablesStr, ",")
	for _, tableName := range tables {
		table := strings.ToLower(strings.ReplaceAll(tableName, " ", ""))
		table = gorm.ToTableName(table)
		table = inflection.Plural(table)
		primaryKey := CheckPrimaryKey(table)
		primaryKey = gorm.ToColumnName(primaryKey)
		primaryKeyVal := tmpData[primaryKey]
		fields := map[string]interface{}{}
		fields[primaryKey] = primaryKeyVal
		if data, err := FindResultsByFields(table, fields); err == nil {
			tmp := map[string]interface{}{}
			if data[0] != nil {
				for key, val := range data[0] {
					tmp[strcase.UpperCamelCase(key)] = val
				}
			}
			tmpDataRet[tableName] = tmp
		}
	}
	return tmpDataRet
}

func GenericDoForAnotherComplex(c *fiber.Ctx, table string, accessControlTables string, existsInAllTables bool) (int, string) {
	_, _, primaryKeyVal := PrimaryKeyData(table, c) // check... v1
	// fields := make(map[string]interface{})
	// fields[primaryKey] = primaryKeyVal
	// fields["user_id"] = c.Locals("auth:=:userId").(string)
	// loop through tables
	tablesStr := strings.ReplaceAll(accessControlTables, " ", "")
	tables := strings.Split(tablesStr, ",")
	for _, tableNameStr := range tables {
		tmp := strings.Split(tableNameStr, ":")
		tableName := tmp[0]
		table := strings.ToLower(strings.ReplaceAll(tableName, " ", "")) // tables
		table = gorm.ToTableName(table)
		table = inflection.Plural(table)
		tmpTable := tmp[1]
		lastUnderscore := strings.LastIndex(tmpTable, "_")
		if lastUnderscore != -1 {
			tmpTable = tmpTable[:lastUnderscore]
		}
		_, primaryKey, primaryKeyVal := PrimaryKeyData(tmpTable, c) // check. Added
		fields := make(map[string]interface{})
		fields[primaryKey] = primaryKeyVal
		fields["user_id"] = c.Locals("auth:=:userId").(string)
		results, err := FindResultsByFields(table, fields)
		if err != nil {
			return 500, primaryKeyVal
		}
		if len(results) > 0 { // if results exist in any single table
			if !existsInAllTables { // results in any single table should return true
				return 0, primaryKeyVal // found in atleast one table
			} else {
				// check all the other tables
			} //
		} else { // missing in one table
			if existsInAllTables { // missing result in one table results in negative
				return 401, primaryKeyVal
			} else {
				// check all the other tables
			}
		}
	}
	if existsInAllTables {
		return 0, primaryKeyVal // nothing found to break this condition
	} else {
		return 401, primaryKeyVal
	}

}

func GenericDoForAnother(c *fiber.Ctx, table string, accessControlTables string) (int, string) {
	_, primaryKey, primaryKeyVal := PrimaryKeyData(table, c)
	fields := make(map[string]interface{})
	fields[primaryKey] = primaryKeyVal
	fields["user_id"] = c.Locals("auth:=:userId").(string)
	// loop through tables
	tablesStr := strings.ReplaceAll(accessControlTables, " ", "")
	tables := strings.Split(tablesStr, ",")
	for _, tableName := range tables {
		table := strings.ToLower(strings.ReplaceAll(tableName, " ", ""))
		table = gorm.ToTableName(table)
		table = inflection.Plural(table)
		results, err := FindResultsByFields(table, fields)
		if err != nil {
			return 500, primaryKeyVal
		}
		if len(results) > 0 {
			return 0, primaryKeyVal
		}
	}
	return 401, primaryKeyVal
}

// https://github.com/500bookclub/backend/issues/24
func GenericUniqueGroup(c *fiber.Ctx, table string, uniqueGroup string) bool {
	fields := map[string]interface{}{}
	uniqueGroupStr := strings.ReplaceAll(uniqueGroup, " ", "")
	uniqueGroupParts := strings.Split(uniqueGroupStr, ",")
	for _, bby := range uniqueGroupParts {
		bbyBy := strcase.UpperCamelCase(bby)
		fields[bbyBy] = c.Locals(fmt.Sprintf("body:=:%s", bby))
	}

	results, err := FindResultsByFields(table, fields)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	if results == nil || len(results) == 0 {
		return true
	}
	return false
}
