// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/login": {
            "post": {
                "description": "用户登录权鉴",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/control.UserLoginService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/control.UserLoginInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "control.UserLoginInfo": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "control.UserLoginService": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 40,
                    "minLength": 8
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
