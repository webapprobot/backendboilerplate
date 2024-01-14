package validationControllers

// var pathsConfig = map[string]interface{}{
// 	"/api/noauth/user": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"sendActivationMail": notRequiredBoolNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 				"fields": struct { // All fields
// 					Email          string `json:"name"` //
// 					FirstName      string `json:"FirstName"`
// 					LastName       string `json:"LastName"`
// 					Telephone      string `json:"Telephone"`
// 					Authority      string `json:"Authority"`
// 					Gender         string `json:"Gender"`
// 					DateOfBirth    string `json:"DateOfBirth"`
// 					activationPage string
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Email": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "asd@cseco.co.ke",
// 						"value":    "",    // the value supplied of the default value
// 						"supplied": false, // if value has been supplied
// 						// "len": map[string]int{
// 						// 	"min": 20,
// 						// 	"max": 30,
// 						// },
// 					},
// 					"FirstName": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",    // the value supplied of the default value
// 						"supplied": false, // if value has been supplied
// 						"len": map[string]int{
// 							"min": 2,
// 							"max": 20,
// 						},
// 					},
// 					"LastName": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",    // the value supplied of the default value
// 						"supplied": false, // if value has been supplied
// 						"len": map[string]int{
// 							"min": 2,
// 							"max": 20,
// 						},
// 					},
// 					"Authority": map[string]interface{}{
// 						"type":     "enum",
// 						"required": true,
// 						"default":  "SYS_USER",
// 						"value":    "",    // the value supplied of the default value
// 						"supplied": false, // if value has been supplied
// 						"enum":     []string{"SYS_ADMIN", "SYS_USER"},
// 					},
// 					"Gender": map[string]interface{}{
// 						"type":     "enum",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",    // the value supplied of the default value
// 						"supplied": false, // if value has been supplied
// 						"enum":     []string{"Male", "Female"},
// 					},
// 					"Telephone":      requiredTelephoneNotSupplied,
// 					"DateOfBirth":    requiredDateNotSupplied,
// 					"activationPage": notRequiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/auth/social/github": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Code":     requiredStringNotSupplied,
// 				"clientId": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/auth/social/facebook": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Code":         requiredStringNotSupplied,
// 				"clientId":     requiredStringNotSupplied,
// 				"redirect_uri": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/auth/social/twitter": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Code":         requiredStringNotSupplied,
// 				"clientId":     requiredStringNotSupplied,
// 				"redirect_uri": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/auth/social/linkedin": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Code":         requiredStringNotSupplied,
// 				"clientId":     requiredStringNotSupplied,
// 				"redirect_uri": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/noauth/activate": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"activateToken": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/noauth/forgotPassword": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Email      string `json:"Email"`
// 					passwdPage string
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Email":      requiredEmailNotSupplied,
// 					"passwdPage": notRequiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/noauth/resetPassword": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY",
// 				"DOFORANOTHER": false,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Code     string `json:"Code"` //
// 					Password string `json:"Password"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Code":     requiredStringNotSupplied,
// 					"Password": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/auth/login": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY",
// 				"DOFORANOTHER": false,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Email    string `json:"Email"` //
// 					Password string `json:"Password"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Email":    requiredEmailNotSupplied,
// 					"Password": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/auth/token": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/subjects": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Subject":   notRequiredStringNotSupplied,
// 				"SubjectId": notRequiredStringNotSupplied,
// 				"Approved":  notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Subject":   notRequiredStringNotSupplied,
// 				"SubjectId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/subject": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Subject string `json:"Subject"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Subject": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Subject   string `json:"Subject"`
// 					SubjectId string `json:"SubjectId"`
// 					Approved  bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Subject":   requiredStringNotSupplied,
// 					"SubjectId": requiredStringNotSupplied,
// 					"Approved":  notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/grades": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Grade":    notRequiredStringNotSupplied,
// 				"GradeId":  notRequiredStringNotSupplied,
// 				"Approved": notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Grade":   notRequiredStringNotSupplied,
// 				"GradeId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/grade": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Grade string `json:"Grade"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Grade": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Grade    string `json:"Grade"`
// 					GradeId  string `json:"GradeId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Grade":    requiredStringNotSupplied,
// 					"GradeId":  requiredStringNotSupplied,
// 					"Approved": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/sizes": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Size": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"SizeId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"Approved": map[string]interface{}{
// 					"type":     "bool",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"query": map[string]interface{}{
// 				"Size": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"SizeId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/size": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Size string `json:"Size"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Size": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Size     string `json:"Size"`
// 					SizeId   string `json:"SizeId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Size":     notRequiredStringNotSupplied,
// 					"SizeId":   notRequiredStringNotSupplied,
// 					"Approved": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/covers": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Cover": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"CoverId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"Approved": map[string]interface{}{
// 					"type":     "bool",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"query": map[string]interface{}{
// 				"Cover": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"CoverId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/cover": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Cover string `json:"Cover"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Cover": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Cover    string `json:"Cover"`
// 					CoverId  string `json:"CoverId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Cover": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"CoverId": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"Approved": map[string]interface{}{
// 						"type":     "bool",
// 						"required": false,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/colors": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Color": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"ColorId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"Approved": map[string]interface{}{
// 					"type":     "bool",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"query": map[string]interface{}{
// 				"Color": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"ColorId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/color": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Color string `json:"Color"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Color": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Color    string `json:"Color"`
// 					ColorId  string `json:"ColorId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Color": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"ColorId": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"Approved": map[string]interface{}{
// 						"type":     "bool",
// 						"required": false,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/publishers": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Publisher": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"PublisherId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"Approved": map[string]interface{}{
// 					"type":     "bool",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"query": map[string]interface{}{
// 				"Publisher": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"PublisherId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/publisher": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Publisher string `json:"Publisher"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Publisher": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Publisher   string `json:"Publisher"`
// 					PublisherId string `json:"PublisherId"`
// 					Approved    bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Publisher": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"PublisherId": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"Approved": map[string]interface{}{
// 						"type":     "bool",
// 						"required": false,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/authors": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Author": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"AuthorId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"Approved": map[string]interface{}{
// 					"type":     "bool",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"query": map[string]interface{}{
// 				"Author": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 				"AuthorId": map[string]interface{}{
// 					"type":     "string",
// 					"required": false,
// 					"default":  "",
// 					"value":    "",
// 					"supplied": false,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/author": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Author string `json:"Author"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Author": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Author   string `json:"Author"`
// 					AuthorId string `json:"AuthorId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Author": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"AuthorId": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"Approved": map[string]interface{}{
// 						"type":     "bool",
// 						"required": false,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 	},
// 	"/api/book": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Isbn        string `json:"Isbn"`
// 					Title       string `json:"Title"`
// 					AuthorId    string `json:"AuthorId"`
// 					PublisherId string `json:"PublisherId"`
// 					SubjectId   string `json:"SubjectId"`
// 					GradeId     string `json:"GradeId"`
// 					CoverId     string `json:"CoverId"`
// 					ColorId     string `json:"ColorId"`
// 					SizeId      string `json:"SizeId"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Isbn":     requiredStringNotSupplied,
// 					"Title":    requiredStringNotSupplied,
// 					"AuthorId": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": map[string]interface{}{ // this is the template to use for all fields
// 					"type":     "string",
// 					"required": true,
// 					"default":  "",    // string. Will be converted to correct type by the function
// 					"value":    "",    // the value supplied of the default value
// 					"supplied": false, // if value has been supplied
// 				},
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct { // All fields
// 					Author   string `json:"Author"`
// 					AuthorId string `json:"AuthorId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{ // this is the template to use for all fields
// 					"Author": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"AuthorId": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"Approved": map[string]interface{}{
// 						"type":     "bool",
// 						"required": false,
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 	},
// }

// import booksModels "github.com/webappbot/backendboilerplate/local-src/models/books"

// var book = booksModels.Book{}

// var requiredEmailNotSupplied = map[string]interface{}{
// 	"type":     "email",
// 	"required": true,
// 	"default":  "",
// 	"value":    "",
// 	"supplied": false,
// }
// var notRequiredEmailNotSupplied = map[string]interface{}{
// 	"type":     "email",
// 	"required": false,
// 	"default":  "",
// 	"value":    "",
// 	"supplied": false,
// }
// var requiredStringNotSupplied = map[string]interface{}{
// 	"type":     "string",
// 	"required": true,
// 	"default":  "",
// 	"value":    "",
// 	"supplied": false,
// }
// var notRequiredStringNotSupplied = map[string]interface{}{
// 	"type":     "string",
// 	"required": false,
// 	"default":  "",
// 	"value":    "",
// 	"supplied": false,
// }
// var requiredTelephoneNotSupplied = map[string]interface{}{
// 	"type":     "telephone",
// 	"required": true,
// 	"default":  "",
// 	"value":    "",
// 	"supplied": false,
// }

// var requiredDateNotSupplied = map[string]interface{}{
// 	"type":     "date",
// 	"required": true,
// 	"default":  "",
// 	"value":    "",
// 	"supplied": false,
// }
// var requiredBoolNotSupplied = map[string]interface{}{
// 	"type":     "bool",
// 	"required": true,
// 	"default":  "",
// 	"value":    "",
// 	"supplied": false,
// }
// var notRequiredBoolNotSupplied = map[string]interface{}{ // this is the template to use for all fields
// 	"type":     "bool",
// 	"required": false,
// 	"default":  "",    // string. Will be converted to correct type by the function
// 	"value":    "",    // the value supplied of the default value
// 	"supplied": false, // if value has been supplied
// }
// var notRequiredIntNotSupplied = map[string]interface{}{ // this is the template to use for all fields
// 	"type":     "int",
// 	"required": false,
// 	"default":  "",    // string. Will be converted to correct type by the function
// 	"value":    "",    // the value supplied of the default value
// 	"supplied": false, // if value has been supplied
// }
// var pathsConfig = map[string]interface{}{
// 	"/api/noauth/user": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"sendActivationMail": notRequiredBoolNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 				"fields": struct {
// 					Email          string `json:"name"` //
// 					FirstName      string `json:"FirstName"`
// 					LastName       string `json:"LastName"`
// 					Telephone      string `json:"Telephone"`
// 					Authority      string `json:"Authority"`
// 					Gender         string `json:"Gender"`
// 					DateOfBirth    string `json:"DateOfBirth"`
// 					activationPage string
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Email": map[string]interface{}{
// 					"type":     "string",
// 					"required": true,
// 					"default":  "asd@cseco.co.ke",
// 					"value":    "", "supplied": false, // "len": map[string]int{
// 					// "min": 20,
// 					// "max": 30,
// 					// },
// 				},
// 					"FirstName": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "", "supplied": false,
// 						"len": map[string]int{
// 							"min": 2,
// 							"max": 20,
// 						},
// 					},
// 					"LastName": map[string]interface{}{
// 						"type":     "string",
// 						"required": true,
// 						"default":  "",
// 						"value":    "", "supplied": false, "len": map[string]int{
// 							"min": 2,
// 							"max": 20,
// 						},
// 					},
// 					"Authority": map[string]interface{}{
// 						"type":     "enum",
// 						"required": true,
// 						"default":  "SYS_USER",
// 						"value":    "", "supplied": false, "enum": []string{"SYS_ADMIN", "SYS_USER"},
// 					},
// 					"Gender": map[string]interface{}{
// 						"type":     "enum",
// 						"required": true,
// 						"default":  "",
// 						"value":    "", "supplied": false, "enum": []string{"Male", "Female"},
// 					},
// 					"Telephone":      requiredTelephoneNotSupplied,
// 					"DateOfBirth":    requiredDateNotSupplied,
// 					"activationPage": notRequiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/auth/social/github": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Code":     requiredStringNotSupplied,
// 				"clientId": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/auth/social/facebook": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Code":         requiredStringNotSupplied,
// 				"clientId":     requiredStringNotSupplied,
// 				"redirect_uri": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/auth/social/twitter": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Code":         requiredStringNotSupplied,
// 				"clientId":     requiredStringNotSupplied,
// 				"redirect_uri": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/auth/social/linkedin": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Code":         requiredStringNotSupplied,
// 				"clientId":     requiredStringNotSupplied,
// 				"redirect_uri": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/noauth/activate": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"activateToken": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{ // differs from query and header because we have to get the c.Body() into a struct that we need to have saved somewhere
// 			},
// 		},
// 	},
// 	"/api/noauth/forgotPassword": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Email      string `json:"Email"`
// 					passwdPage string
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Email": requiredEmailNotSupplied,
// 					"passwdPage": notRequiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/noauth/resetPassword": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY",
// 				"DOFORANOTHER": false,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Code     string `json:"Code"` //
// 					Password string `json:"Password"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Code": requiredStringNotSupplied,
// 					"Password": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/auth/login": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY",
// 				"DOFORANOTHER": false,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Email    string `json:"Email"` //
// 					Password string `json:"Password"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Email": requiredEmailNotSupplied,
// 					"Password": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/auth/token": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/subjects": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Subject":   notRequiredStringNotSupplied,
// 				"SubjectId": notRequiredStringNotSupplied,
// 				"Approved":  notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Subject":   notRequiredStringNotSupplied,
// 				"SubjectId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/subject": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Subject string `json:"Subject"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Subject": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Subject   string `json:"Subject"`
// 					SubjectId string `json:"SubjectId"`
// 					Approved  bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Subject": requiredStringNotSupplied,
// 					"SubjectId": requiredStringNotSupplied,
// 					"Approved":  notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/grades": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Grade":    notRequiredStringNotSupplied,
// 				"GradeId":  notRequiredStringNotSupplied,
// 				"Approved": notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Grade":   notRequiredStringNotSupplied,
// 				"GradeId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/grade": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Grade string `json:"Grade"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Grade": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Grade    string `json:"Grade"`
// 					GradeId  string `json:"GradeId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Grade": requiredStringNotSupplied,
// 					"GradeId":  requiredStringNotSupplied,
// 					"Approved": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/sizes": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Size":     notRequiredStringNotSupplied,
// 				"SizeId":   notRequiredStringNotSupplied,
// 				"Approved": notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Size":   notRequiredStringNotSupplied,
// 				"SizeId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/size": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Size string `json:"Size"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Size": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Size     string `json:"Size"`
// 					SizeId   string `json:"SizeId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Size": notRequiredStringNotSupplied,
// 					"SizeId":   notRequiredStringNotSupplied,
// 					"Approved": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/covers": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Cover":    notRequiredStringNotSupplied,
// 				"CoverId":  notRequiredStringNotSupplied,
// 				"Approved": notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Cover":   notRequiredStringNotSupplied,
// 				"CoverId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/cover": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Cover string `json:"Cover"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Cover": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Cover    string `json:"Cover"`
// 					CoverId  string `json:"CoverId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Cover": requiredStringNotSupplied,
// 					"CoverId":  requiredStringNotSupplied,
// 					"Approved": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/colors": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Color":    notRequiredStringNotSupplied,
// 				"ColorId":  notRequiredStringNotSupplied,
// 				"Approved": notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Color":   notRequiredStringNotSupplied,
// 				"ColorId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/color": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Color string `json:"Color"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Color": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Color    string `json:"Color"`
// 					ColorId  string `json:"ColorId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Color": requiredStringNotSupplied,
// 					"ColorId":  requiredStringNotSupplied,
// 					"Approved": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/publishers": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Publisher":   notRequiredStringNotSupplied,
// 				"PublisherId": notRequiredStringNotSupplied,
// 				"Approved":    notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Publisher":   notRequiredStringNotSupplied,
// 				"PublisherId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/publisher": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Publisher string `json:"Publisher"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Publisher": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Publisher   string `json:"Publisher"`
// 					PublisherId string `json:"PublisherId"`
// 					Approved    bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Publisher": requiredStringNotSupplied,
// 					"PublisherId": requiredStringNotSupplied,
// 					"Approved":    notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/authors": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Author":   notRequiredStringNotSupplied,
// 				"AuthorId": notRequiredStringNotSupplied,
// 				"Approved": notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Author":   notRequiredStringNotSupplied,
// 				"AuthorId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/author": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Author string `json:"Author"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Author": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Author   string `json:"Author"`
// 					AuthorId string `json:"AuthorId"`
// 					Approved bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Author": requiredStringNotSupplied,
// 					"AuthorId": requiredStringNotSupplied,
// 					"Approved": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/books/bindings": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"query": map[string]interface{}{
// 				"Binding":   notRequiredStringNotSupplied,
// 				"BindingId": notRequiredStringNotSupplied,
// 				"Approved":  notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Binding":   notRequiredStringNotSupplied,
// 				"BindingId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/books/binding": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Binding string `json:"Binding"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Binding": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Binding   string `json:"Binding"`
// 					BindingId string `json:"BindingId"`
// 					Approved  bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Binding": requiredStringNotSupplied,
// 					"BindingId": requiredStringNotSupplied,
// 					"Approved":  notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/charities": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":        "charities",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "GET",
// 			"query": map[string]interface{}{
// 				"Charity":   notRequiredStringNotSupplied,
// 				"CharityId": notRequiredStringNotSupplied,
// 				"Approved":  notRequiredBoolNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":        "charities",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "DELETE",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Charity":   notRequiredStringNotSupplied,
// 				"CharityId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/charity": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":        "charities",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "POST",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Charity string `json:"Charity"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Charity": requiredStringNotSupplied},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":        "charities",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "PATCH",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Charity   string `json:"Charity"`
// 					CharityId string `json:"CharityId"`
// 					Approved  bool   `json:"Approved"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{"Charity": requiredStringNotSupplied,
// 					"CharityId": requiredStringNotSupplied,
// 					"Approved":  notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/libraries": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":        "libraries",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "GET",
// 			"preload": "User",
// 			"query": map[string]interface{}{
// 				"Library":   notRequiredStringNotSupplied,
// 				"LibraryId": notRequiredStringNotSupplied,
// 				"Personal":  notRequiredBoolNotSupplied,
// 				"UserId":    notRequiredStringNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":               "libraries",
// 				"DOFORANOTHER":        false, // do for another and access control tables work together
// 				"AccessControlTables": "libraries, library_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 			},
// 			"generic": "DELETE",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"LibraryId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/library": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":        "libraries",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic":  "POST",
// 			"patchUID": true,
// 			"preload":  "User",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					UserId   string `json:"UserId"`
// 					Library  string `json:"Library"`
// 					Personal string `json:"Personal"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"UserId":   requiredStringNotSupplied,
// 					"Library":  requiredStringNotSupplied,
// 					"Personal": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":               "libraries",
// 				"DOFORANOTHER":        false, // do for another and access control tables work together
// 				"AccessControlTables": "libraries, library_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 			},
// 			"generic": "PATCH",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					LibraryId string `json:"LibraryId"`
// 					Library   string `json:"Library"`
// 					UserId    string `json:"UserId"`
// 					Personal  bool   `json:"Personal"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"Library":   requiredStringNotSupplied,
// 					"UserId":    requiredStringNotSupplied,
// 					"LibraryId": requiredStringNotSupplied,
// 					"Personal":  notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/shelves": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":        "shelves",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "GET",
// 			"preload": "User",
// 			"query": map[string]interface{}{
// 				"Shelf":   notRequiredStringNotSupplied,
// 				"ShelfId": notRequiredStringNotSupplied,
// 				"Shop":    notRequiredBoolNotSupplied,
// 				"UserId":  notRequiredStringNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":               "shelves",
// 				"DOFORANOTHER":        false, // do for another and access control tables work together
// 				"AccessControlTables": "shelves, shelf_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 			},
// 			"generic": "DELETE",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"ShelfId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/shelf": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":        "shelves",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic":  "POST",
// 			"patchUID": true,
// 			"preload":  "User",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					UserId string `json:"UserId"`
// 					Shelf  string `json:"Shelf"`
// 					Shop   bool   `json:"Shop"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"UserId": requiredStringNotSupplied,
// 					"Shelf":  requiredStringNotSupplied,
// 					"Shop":   notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":               "shelves",
// 				"DOFORANOTHER":        false, // do for another and access control tables work together
// 				"AccessControlTables": "shelves, shelf_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 			},
// 			"generic": "PATCH",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					UserId  string `json:"UserId"`
// 					Shelf   string `json:"Shelf"`
// 					ShelfId string `json:"ShelfId"`
// 					Shop    bool   `json:"Shop"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"UserId":  requiredStringNotSupplied,
// 					"Shelf":   requiredStringNotSupplied,
// 					"ShelfId": requiredStringNotSupplied,
// 					"Shop":    notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/shelfowners": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":        "shelf_owners",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "GET",
// 			"preload": "User, Shelf",
// 			"query": map[string]interface{}{
// 				"ShelfOwnerId": notRequiredStringNotSupplied,
// 				"ShelfId":      notRequiredStringNotSupplied,
// 				"UserId":       notRequiredStringNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":               "shelf_owners",
// 				"DOFORANOTHER":        false, // do for another and access control tables work together
// 				"AccessControlTables": "shelves, shelf_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 			},
// 			"generic": "DELETE",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"ShelfOwnerId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/shelfowner": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":               "shelf_owners",
// 				"DOFORANOTHER":        false,
// 				"AccessControlTables": "shelves, shelf_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"UniqueGroup": "UserId, ShelfId", // requires permission.Table
// 			},
// 			"generic":  "POST",
// 			"patchUID": false,
// 			"preload":  "User, Shelf",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					UserId  string `json:"UserId"`
// 					ShelfId string `json:"ShelfId"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"UserId":  requiredStringNotSupplied,
// 					"ShelfId": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":               "shelf_owners",
// 				"DOFORANOTHER":        false, // do for another and access control tables work together
// 				"AccessControlTables": "shelves, shelf_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 				"UniqueGroup":  "UserId, ShelfId", // requires permission.Table
// 			},
// 			"generic":  "PATCH",
// 			"patchUID": false,
// 			"preload":  "User, Shelf",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					UserId       string `json:"UserId"`
// 					ShelfId      string `json:"ShelfId"`
// 					ShelfOwnerId string `json:"ShelfOwnerId"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"UserId":       requiredStringNotSupplied,
// 					"ShelfId":      requiredStringNotSupplied,
// 					"ShelfOwnerId": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/libraryowners": map[string]interface{}{
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":        "library_owners",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "GET",
// 			"preload": "User, Library",
// 			"query": map[string]interface{}{
// 				"LibraryOwnerId": notRequiredStringNotSupplied,
// 				"LibraryId":      notRequiredStringNotSupplied,
// 				"UserId":         notRequiredStringNotSupplied,
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":               "library_owners",
// 				"DOFORANOTHER":        false, // do for another and access control tables work together
// 				"AccessControlTables": "libraries, library_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 			},
// 			"generic": "DELETE",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"LibraryOwnerId": notRequiredStringNotSupplied,
// 			},
// 		},
// 	},
// 	"/api/libraryowner": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":               "library_owners",
// 				"DOFORANOTHER":        false,
// 				"AccessControlTables": "libraries, library_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"UniqueGroup": "UserId, LibraryId", // requires permission.Table
// 			},
// 			"generic":  "POST",
// 			"patchUID": false,
// 			"preload":  "User, Library",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					UserId    string `json:"UserId"`
// 					LibraryId string `json:"LibraryId"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"UserId":    requiredStringNotSupplied,
// 					"LibraryId": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":                "SYS_USER", // check: change to SYS_ADMIN
// 				"Table":               "library_owners",
// 				"DOFORANOTHER":        false, // do for another and access control tables work together
// 				"AccessControlTables": "libraries, library_owners",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 				"UniqueGroup":  "UserId, LibraryId", // requires permission.Table
// 			},
// 			"generic":  "PATCH",
// 			"patchUID": false,
// 			"preload":  "User, Library",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					UserId         string `json:"UserId"`
// 					LibraryId      string `json:"LibraryId"`
// 					LibraryOwnerId string `json:"LibraryOwnerId"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"UserId":         requiredStringNotSupplied,
// 					"LibraryId":      requiredStringNotSupplied,
// 					"LibraryOwnerId": requiredStringNotSupplied,
// 				},
// 			},
// 		},
// 	},
// 	"/api/book": map[string]interface{}{
// 		"POST": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 				"Table":        "books",
// 			},
// 			"generic": "POST",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Isbn        string `json:"Isbn"`
// 					Title       string `json:"Title"`
// 					AuthorId    string `json:"AuthorId"`
// 					PublisherId string `json:"PublisherId"`
// 					SubjectId   string `json:"SubjectId"`
// 					GradeId     string `json:"GradeId"`
// 					CoverId     string `json:"CoverId"`
// 					ColorId     string `json:"ColorId"`
// 					SizeId      string `json:"SizeId"`
// 					BindingId   string `json:"BindingId"`
// 					Edition     string `json:"Edition"`
// 					Year        string `json:"Year"`
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"Isbn":        requiredStringNotSupplied,
// 					"Title":       requiredStringNotSupplied,
// 					"AuthorId":    requiredStringNotSupplied,
// 					"PublisherId": requiredStringNotSupplied,
// 					"SubjectId":   requiredStringNotSupplied,
// 					"GradeId":     requiredStringNotSupplied,
// 					"CoverId":     requiredStringNotSupplied,
// 					"ColorId":     requiredStringNotSupplied,
// 					"SizeId":      requiredStringNotSupplied,
// 					"BindingId":   requiredStringNotSupplied,
// 					"Edition":     requiredStringNotSupplied,
// 					"Year": map[string]interface{}{
// 						"type":     "int",
// 						"required": true,
// 						"range": map[string]interface{}{
// 							"min": 100,
// 							"max": 2050,
// 						},
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 				},
// 			},
// 		},
// 		"PATCH": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER", // check: change to SYS_ADMIN
// 				"DOFORANOTHER": false,
// 				"Table":        "books",
// 			},
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 			},
// 			"generic": "PATCH",
// 			"preload": "Author, Publisher, Subject, Grade, Cover, Color, Size, Binding",
// 			// "model":   booksModels.BookBase{},
// 			"model": book,
// 			"body": map[string]interface{}{
// 				"fields": struct {
// 					Isbn        string `json:"Isbn"`
// 					Title       string `json:"Title"`
// 					AuthorId    string `json:"AuthorId"`
// 					PublisherId string `json:"PublisherId"`
// 					SubjectId   string `json:"SubjectId"`
// 					GradeId     string `json:"GradeId"`
// 					CoverId     string `json:"CoverId"`
// 					ColorId     string `json:"ColorId"`
// 					SizeId      string `json:"SizeId"`
// 					BindingId   string `json:"BindingId"`
// 					Edition     string `json:"Edition"`
// 					Year        string `json:"Year"`
// 					Approved    bool
// 				}{},
// 				"fieldSettings": map[string]interface{}{
// 					"Isbn":        requiredStringNotSupplied,
// 					"Title":       requiredStringNotSupplied,
// 					"AuthorId":    requiredStringNotSupplied,
// 					"PublisherId": requiredStringNotSupplied,
// 					"SubjectId":   requiredStringNotSupplied,
// 					"GradeId":     requiredStringNotSupplied,
// 					"CoverId":     requiredStringNotSupplied,
// 					"ColorId":     requiredStringNotSupplied,
// 					"SizeId":      requiredStringNotSupplied,
// 					"BindingId":   requiredStringNotSupplied,
// 					"Edition":     requiredStringNotSupplied,
// 					"Year": map[string]interface{}{
// 						"type":     "int",
// 						"required": true,
// 						"range": map[string]interface{}{
// 							"min": 100,
// 							"max": 2050,
// 						},
// 						"default":  "",
// 						"value":    "",
// 						"supplied": false,
// 					},
// 					"Approved": notRequiredBoolNotSupplied,
// 				},
// 			},
// 		},
// 		"DELETE": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "SYS_USER",
// 				"DOFORANOTHER": false,
// 				"Table":        "books",
// 			},
// 			"validations": map[string]interface{}{
// 				"EnsureExists": true,
// 			},
// 			"generic": "DELETE",
// 			"header": map[string]interface{}{
// 				"X-Authorization": requiredStringNotSupplied,
// 			},
// 			"query": map[string]interface{}{
// 				"Isbn": requiredStringNotSupplied,
// 			},
// 		},
// 		"GET": map[string]interface{}{
// 			"permission": map[string]interface{}{
// 				"user":         "NOBODY", // NOBODY, SYS_USER, SYS_ADMIN
// 				"Table":        "books",
// 				"DOFORANOTHER": false,
// 			},
// 			"generic": "GET",
// 			"preload": "Author, Publisher, Subject, Grade, Cover, Color, Size, Binding",
// 			"query": map[string]interface{}{
// 				"Isbn":        notRequiredStringNotSupplied,
// 				"Title":       notRequiredStringNotSupplied,
// 				"AuthorId":    notRequiredStringNotSupplied,
// 				"PublisherId": notRequiredStringNotSupplied,
// 				"SubjectId":   notRequiredStringNotSupplied,
// 				"GradeId":     notRequiredStringNotSupplied,
// 				"CoverId":     notRequiredStringNotSupplied,
// 				"ColorId":     notRequiredStringNotSupplied,
// 				"SizeId":      notRequiredStringNotSupplied,
// 				"BindingId":   notRequiredStringNotSupplied,
// 				"Edition":     notRequiredStringNotSupplied,
// 				"Approved":    notRequiredBoolNotSupplied,
// 				"Year":        notRequiredIntNotSupplied,
// 			},
// 		},
// 	},
// }
