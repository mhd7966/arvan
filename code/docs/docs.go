// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
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
        "/charge": {
            "post": {
                "description": "create charge",
                "tags": [
                    "Charge"
                ],
                "summary": "create charge",
                "operationId": "create_charge",
                "parameters": [
                    {
                        "description": "charge code info",
                        "name": "charge_code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inputs.ChargeCode"
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
        "/charge/{charge_code}": {
            "get": {
                "description": "get charge",
                "tags": [
                    "Charge"
                ],
                "summary": "get charge",
                "operationId": "get_charge",
                "parameters": [
                    {
                        "type": "string",
                        "description": "charge code name",
                        "name": "charge_code",
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
        "/charge/{charge_code}/apply": {
            "post": {
                "description": "charge code",
                "tags": [
                    "Charge"
                ],
                "summary": "charge",
                "operationId": "charge",
                "parameters": [
                    {
                        "type": "string",
                        "description": "charge code name",
                        "name": "charge_code",
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
        "/charge/{charge_code}/rollback": {
            "post": {
                "description": "charge rollback",
                "tags": [
                    "Charge"
                ],
                "summary": "charge",
                "operationId": "charge_rollback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "charge code name",
                        "name": "charge_code",
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
        "inputs.ChargeCode": {
            "type": "object",
            "properties": {
                "expiration_date": {
                    "type": "string"
                },
                "max_capacity": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
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
	Title:       "Code API",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
