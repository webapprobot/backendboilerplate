package validatorZPaths

var RequiredEmailNotSupplied = map[string]interface{}{
	"type":     "email",
	"required": true,
	"default":  "",
	"value":    "",
	"supplied": false,
}
var NotRequiredEmailNotSupplied = map[string]interface{}{
	"type":     "email",
	"required": false,
	"default":  "",
	"value":    "",
	"supplied": false,
}
var RequiredStringNotSupplied = map[string]interface{}{
	"type":     "string",
	"required": true,
	"default":  "",
	"value":    "",
	"supplied": false,
}
var NotRequiredStringNotSupplied = map[string]interface{}{
	"type":     "string",
	"required": false,
	"default":  "",
	"value":    "",
	"supplied": false,
}
var RequiredTelephoneNotSupplied = map[string]interface{}{
	"type":     "telephone",
	"required": true,
	"default":  "",
	"value":    "",
	"supplied": false,
}

//	var RequiredYearNotSupplied = map[string]interface{}{
//		"type":     "year",
//		"required": true,
//		"default":  "",
//		"value":    "",
//		"supplied": false,
//	}
var RequiredDateNotSupplied = map[string]interface{}{
	"type":     "date",
	"required": true,
	"default":  "",
	"value":    "",
	"supplied": false,
}
var RequiredBoolNotSupplied = map[string]interface{}{
	"type":     "bool",
	"required": true,
	"default":  "",
	"value":    "",
	"supplied": false,
}
var NotRequiredBoolNotSupplied = map[string]interface{}{ // this is the template to use for all fields
	"type":     "bool",
	"required": false,
	"default":  "",    // string. Will be converted to correct type by the function
	"value":    "",    // the value supplied of the default value
	"supplied": false, // if value has been supplied
}
