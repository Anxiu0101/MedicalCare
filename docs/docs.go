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
        "contact": {
            "name": "Anxiu0101",
            "url": "https://github.com/Anxiu0101/MedicalCare/blob/master/README.md",
            "email": "anxiu.fyc@foxmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/blog": {
            "post": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
                "description": "用于新建文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "文章发布接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章简介",
                        "name": "desc",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "文章内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/info": {
            "get": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
                "description": "用于用户获取其个人资料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户资料获取接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/service.UserService"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
                "description": "用于用户更新其个人资料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户资料更新接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "用户性别",
                        "name": "gender",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "用户年龄",
                        "name": "age",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "联系方式",
                        "name": "tel",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/login": {
            "post": {
                "description": "用于用户登陆，接受用户名和用户密码，将 access token 和 refresh token 作为结果返回，同时返回用户的账户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户登陆接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/model.TokenData"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "User": {
                                                            "$ref": "#/definitions/model.AccountInfo"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/password": {
            "post": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
                "description": "用于用户更新密码，需要新密码做为传入的参数，需要 token 验证",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户重新设置密码接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/register": {
            "post": {
                "description": "用于注册新用户，传入 username 和 password 作为新用户的 username 和 password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/chat/:receiver": {
            "get": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
                "description": "用于聊天室的使用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "聊天室相关接口"
                ],
                "summary": "聊天室接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "聊天对象",
                        "name": "receiver",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/group": {
            "post": {
                "description": "set up a group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "create group",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/group/member": {
            "put": {
                "description": "invite group members",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "invite member",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AccountInfo": {
            "type": "object",
            "properties": {
                "avatars": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "Anxiu"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.TokenData": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "user": {}
            }
        },
        "service.UserService": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 20
                },
                "email": {
                    "type": "string",
                    "example": "email@example.com"
                },
                "gender": {
                    "type": "integer",
                    "example": 1
                },
                "tel": {
                    "type": "string",
                    "example": "188-8888-6666"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "MedicalCare API",
	Description:      "This is MedicalCare Project api docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
