// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/transactions": {
            "post": {
                "description": "charge code",
                "tags": [
                    "Transaction"
                ],
                "summary": "charge",
                "operationId": "charge",
                "parameters": [
                    {
                        "description": "charge body",
                        "name": "charge",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inputs.Charge"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/transactions/{phone_number}": {
            "get": {
                "description": "get transactions",
                "tags": [
                    "Transaction"
                ],
                "summary": "get transactions",
                "operationId": "get_transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "phone number of user",
                        "name": "phone_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/users/{phone_number}": {
            "get": {
                "description": "return user",
                "tags": [
                    "User"
                ],
                "summary": "get user",
                "operationId": "get_user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "phone number of user",
                        "name": "phone_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "inputs.Charge": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "code_type": {
                    "type": "integer",
                    "default": 1
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "inputs.GetTransactionsPagination": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer",
                    "default": 1
                },
                "size": {
                    "type": "integer",
                    "default": 10
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/v0",
	Schemes:     []string{},
	Title:       "Wallet API",
	Description: "I have no specific description",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
