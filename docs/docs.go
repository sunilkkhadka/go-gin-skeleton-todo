// Code generated by swaggo/swag. DO NOT EDIT
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
        "/profile": {
            "get": {
                "description": "Get one user by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get one user by id",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
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
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/responses.Data"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "$ref": "#/definitions/dtos.GetUserResponse"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Return all the User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all User.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "10",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Page no",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by name",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by type",
                        "name": "Keyword2",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
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
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/responses.DataCount"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "type": "array",
                                                "items": {
                                                    "$ref": "#/definitions/dtos.GetUserResponse"
                                                }
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Enter JSON",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateUserRequestData"
                        }
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateUserRequestData": {
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
        "dtos.GetUserResponse": {
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
        "responses.Data": {
            "type": "object",
            "properties": {
                "data": {}
            }
        },
        "responses.DataCount": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {}
            }
        },
        "responses.Error": {
            "type": "object",
            "properties": {
                "error": {}
            }
        },
        "responses.Success": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Boilerplate API",
	Description:      "An API in Go using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
