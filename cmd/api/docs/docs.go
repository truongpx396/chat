// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/register": {
            "post": {
                "description": "User registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "User registration",
                "operationId": "UserRegister",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Operation Id",
                        "name": "OperationId",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Secret is the Openim key. For details, see the server Config.yaml Secret field.",
                        "name": "UserInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/chat.RegisterUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apistruct.UserRegisterResp"
                        }
                    },
                    "400": {
                        "description": "Errcode is",
                        "schema": {}
                    },
                    "500": {
                        "description": "ERRCODE is 500 generally an internal error of the server",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "apistruct.UserRegisterResp": {
            "type": "object",
            "properties": {
                "chatToken": {
                    "type": "string"
                },
                "imToken": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "chat.RegisterUserInfo": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "areaCode": {
                    "type": "string"
                },
                "birth": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "faceURL": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "chat.RegisterUserReq": {
            "type": "object",
            "properties": {
                "autoLogin": {
                    "type": "boolean"
                },
                "deviceID": {
                    "type": "string"
                },
                "invitationCode": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "platform": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/chat.RegisterUserInfo"
                },
                "verifyCode": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "chat-openim.k8s.magiclab396.com",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Open-IM-Chat API",
	Description:      "Open-IM-Chat API server document, all requests in the document have an OperationId field for link tracking",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
