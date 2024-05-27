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
        "/api/v1/profile": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get user profile",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "User Profile",
                "operationId": "GetUserProfile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Data-user_GetUserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserManagementApi"
                ],
                "summary": "All users",
                "operationId": "GetAllUsers",
                "parameters": [
                    {
                        "type": "boolean",
                        "name": "all",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DataCount-user_GetUserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create one user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserManagementApi"
                ],
                "summary": "Create User",
                "operationId": "CreateUser",
                "parameters": [
                    {
                        "description": "Enter JSON",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserRequestData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    }
                }
            }
        },
        "/api/v1/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get user profile",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserManagementApi"
                ],
                "summary": "User Profile",
                "operationId": "GetOneUser",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Data-user_GetUserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ApiError": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "$ref": "#/definitions/ErrResponse-array_api_errors_ErrorContext"
                }
            }
        },
        "Data-user_GetUserResponse": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "$ref": "#/definitions/user.GetUserResponse"
                }
            }
        },
        "DataCount-user_GetUserResponse": {
            "type": "object",
            "required": [
                "count",
                "data"
            ],
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.GetUserResponse"
                    }
                }
            }
        },
        "ErrResponse-array_api_errors_ErrorContext": {
            "type": "object",
            "required": [
                "error",
                "message"
            ],
            "properties": {
                "error": {
                    "type": "string"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api_errors.ErrorContext"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "Message": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "api_errors.ErrorContext": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "user.CreateUserRequestData": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "full_name",
                "gender",
                "password",
                "phone"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "description": "add soft delete in gorm",
                    "allOf": [
                        {
                            "$ref": "#/definitions/gorm.DeletedAt"
                        }
                    ]
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "user.GetUserResponse": {
            "type": "object",
            "required": [
                "email",
                "full_name",
                "gender",
                "password",
                "phone"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "description": "add soft delete in gorm",
                    "allOf": [
                        {
                            "$ref": "#/definitions/gorm.DeletedAt"
                        }
                    ]
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Demo API",
	Description:      "An API in Go using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
