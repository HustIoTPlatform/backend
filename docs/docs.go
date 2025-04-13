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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/attribute/datas/set/logs": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/attribute/datas/{id}": {
            "get": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/board": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/board/device": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/board/device/total": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/board/home": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/board/tenant": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/board/tenant/device/info": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/board/tenant/user/info": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/board/trend": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/board/user/info": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/board/user/update": {
            "post": {
                "responses": {}
            }
        },
        "/api/v1/board/user/update/password": {
            "post": {
                "responses": {}
            }
        },
        "/api/v1/board/{id}": {
            "get": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/casbin/function": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/casbin/function/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/casbin/user": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/casbin/user/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/command/datas/set/logs": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/data_script": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/data_script/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/datapolicy": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            }
        },
        "/api/v1/device": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/device/active": {
            "put": {
                "responses": {}
            }
        },
        "/api/v1/device/check/{deviceNumber}": {
            "get": {
                "tags": [
                    "设备管理"
                ],
                "responses": {}
            }
        },
        "/api/v1/device/detail/{id}": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/device/group": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/device/group/detail/{id}": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/device/group/relation": {
            "get": {
                "responses": {}
            },
            "post": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/device/group/relation/list": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/device/group/tree": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/device/group/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/device/online/status/ws": {
            "get": {
                "description": "通过WebSocket连接获取实时设备在线状态",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "设备"
                ],
                "summary": "获取设备在线状态",
                "responses": {}
            }
        },
        "/api/v1/device/template": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/device/template/chart": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/device/template/detail/{id}": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/device/template/menu": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/device/template/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/device/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/device_config": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/device_config/batch": {
            "put": {
                "responses": {}
            }
        },
        "/api/v1/device_config/menu": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/device_config/{id}": {
            "get": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/dict": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/dict/column": {
            "post": {
                "responses": {}
            }
        },
        "/api/v1/dict/column/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/dict/enum": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/dict/language": {
            "post": {
                "responses": {}
            }
        },
        "/api/v1/dict/language/{id}": {
            "get": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/event/datas": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/file/up": {
            "post": {
                "tags": [
                    "文件上传"
                ],
                "responses": {}
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "使用邮箱或手机号和密码登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户认证"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录凭证",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.LoginRsp"
                        }
                    },
                    "400": {
                        "description": "错误响应",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/logo": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            }
        },
        "/api/v1/notification/services/config": {
            "post": {
                "responses": {}
            }
        },
        "/api/v1/notification/services/config/e-mail/test": {
            "post": {
                "responses": {}
            }
        },
        "/api/v1/notification/services/config/{type}": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/notification_group": {
            "post": {
                "responses": {}
            }
        },
        "/api/v1/notification_group/list": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/notification_group/{id}": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/notification_history/list": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/open/keys": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/open/keys/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/operation_logs": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/ota/package": {
            "get": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/ota/package/": {
            "put": {
                "responses": {}
            }
        },
        "/api/v1/ota/package/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/ota/task": {
            "get": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/ota/task/detail": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            }
        },
        "/api/v1/ota/task/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/protocol_plugin": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/protocol_plugin/config_form": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/protocol_plugin/device_config_form": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/protocol_plugin/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/role": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/role/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/current/detail/{id}": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/current/keys": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/current/keys/ws": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/current/ws": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/current/{id}": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/history": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/history/page": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/history/pagination": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/set/logs": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/telemetry/datas/statistic": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/ui_elements": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/ui_elements/menu": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/ui_elements/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/user": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/user/detail": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/user/tenant/id": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/user/update": {
            "put": {
                "responses": {}
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        }
    },
    "definitions": {
        "errcode.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "description": "用于存储自定义消息",
                    "type": "string"
                }
            }
        },
        "model.LoginReq": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "Login account (enter email or mobile number)",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 5,
                    "example": "test@test.vn"
                },
                "password": {
                    "description": "Password",
                    "type": "string",
                    "maxLength": 512,
                    "minLength": 6,
                    "example": "123456"
                },
                "salt": {
                    "description": "Random salt (required if RSA encryption is enabled by the super administrator)",
                    "type": "string",
                    "maxLength": 512
                }
            }
        },
        "model.LoginRsp": {
            "type": "object",
            "properties": {
                "expires_in": {
                    "description": "Expiration time (in seconds)",
                    "type": "integer"
                },
                "token": {
                    "description": "Login token",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.1.6",
	Host:             "localhost:9999",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Mitras API",
	Description:      "Mitras API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
