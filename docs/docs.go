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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/accounts/{accountId}/schedules": {
            "get": {
                "description": "GetFreeSchedule",
                "summary": "GetFreeSchedule",
                "parameters": [
                    {
                        "type": "string",
                        "description": "accountId",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "from",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "to",
                        "name": "to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/interfaces.FreeScheduleResponseJson"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    }
                }
            }
        },
        "/activate/account": {
            "put": {
                "description": "SendActivateMail",
                "summary": "SendActivateMail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization header",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    },
                    "403": {}
                }
            }
        },
        "/activate/account/{code}": {
            "get": {
                "description": "アカウント有効化",
                "summary": "ActivateAccount",
                "parameters": [
                    {
                        "type": "string",
                        "description": "アクティベート用コード",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "SignIn",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "SignIn",
                "parameters": [
                    {
                        "description": "Mail\u0026Password",
                        "name": "SignRequestJson",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.SignRequestJson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "headers": {
                            "set-authorization": {
                                "type": "string",
                                "description": "ログイン用セッショントークン"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "SignUp",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "SignUp",
                "parameters": [
                    {
                        "description": "Mail\u0026Password",
                        "name": "SignRequestJson",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.SignRequestJson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "headers": {
                            "set-authorization": {
                                "type": "string",
                                "description": "ログイン用セッショントークン"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    }
                }
            }
        },
        "/invoices/": {
            "post": {
                "description": "IssueAnInvoice",
                "summary": "IssueAnInvoice",
                "parameters": [
                    {
                        "description": "IssueAnInvoiceRequestJson",
                        "name": "IssueAnInvoiceRequestJson",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.IssueAnInvoiceRequestJson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Invoice"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    }
                }
            }
        },
        "/invoices/{invoiceId}": {
            "get": {
                "description": "GetInvoice",
                "summary": "GetInvoice",
                "parameters": [
                    {
                        "type": "string",
                        "description": "invoiceId",
                        "name": "invoiceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Invoice"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    }
                }
            },
            "post": {
                "description": "MakePayment",
                "summary": "MakePayment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "invoiceId",
                        "name": "invoiceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "決済サービスより発行されたトークン",
                        "name": "MakePaymentRequestJson",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.MakePaymentRequestJson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Invoice"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    }
                }
            }
        },
        "/mentor/": {
            "get": {
                "description": "メンター詳細一覧取得",
                "summary": "GetMentorAccountDetails",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.AccountDetail"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    }
                }
            }
        },
        "/reservations/reserve/{scheduleId}": {
            "post": {
                "description": "Reserve",
                "summary": "Reserve",
                "parameters": [
                    {
                        "type": "string",
                        "description": "scheduleId",
                        "name": "scheduleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "authorization header",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    },
                    "403": {}
                }
            }
        },
        "/reservations/reserved/child": {
            "get": {
                "description": "GetReservedReservationsByChildAccountId",
                "summary": "GetReservedReservationsByChildAccountId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization header",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/interfaces.ReservedReservationResponseJson"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    },
                    "403": {}
                }
            }
        },
        "/reservations/reserved/parent": {
            "get": {
                "description": "GetReservedReservationsByParentAccountId",
                "summary": "GetReservedReservationsByParentAccountId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization header",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/interfaces.ReservedReservationResponseJson"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/json.ErrorMessageJson"
                        }
                    },
                    "403": {}
                }
            }
        }
    },
    "definitions": {
        "domain.AccountDetail": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.Invoice": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "invoiceId": {
                    "type": "string"
                },
                "paid": {
                    "type": "boolean"
                }
            }
        },
        "interfaces.FreeScheduleResponseJson": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string"
                },
                "scheduleId": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                },
                "unitPrice": {
                    "type": "integer"
                }
            }
        },
        "interfaces.IssueAnInvoiceRequestJson": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                }
            }
        },
        "interfaces.MakePaymentRequestJson": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "interfaces.ReservedReservationResponseJson": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "childAccountId": {
                    "type": "string"
                },
                "from": {
                    "type": "string"
                },
                "invoiceId": {
                    "type": "string"
                },
                "paid": {
                    "type": "boolean"
                },
                "parentAccountId": {
                    "type": "string"
                },
                "reservationId": {
                    "type": "string"
                },
                "scheduleId": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        },
        "interfaces.SignRequestJson": {
            "type": "object",
            "properties": {
                "mail": {
                    "type": "string"
                },
                "pass": {
                    "type": "string"
                }
            }
        },
        "json.ErrorMessageJson": {
            "type": "object",
            "properties": {
                "message": {
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
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger Letustalk API",
	Description: "This is a sample server Petstore server.",
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
