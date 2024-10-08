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
        "/api/v1/channels": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1",
                    "channels"
                ],
                "summary": "Get YouTube channels",
                "parameters": [
                    {
                        "type": "string",
                        "description": "channel topic",
                        "name": "topic",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "channels count",
                        "name": "max_results",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "channels language",
                        "name": "language",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "csv filename (if want to save results to csv)",
                        "name": "filename",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v1handlers.getChannelsResp"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apiutils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apiutils.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1",
                    "channels"
                ],
                "summary": "Insert YouTube channels to database",
                "parameters": [
                    {
                        "description": "data for searching channels to insert",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1handlers.insertChannelsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apiutils.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apiutils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "tags": [
                    "default"
                ],
                "summary": "Healthcheck route",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "apiutils.ErrorResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                }
            }
        },
        "v1handlers.getChannelsResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "subscriptions": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                }
            }
        },
        "v1handlers.insertChannelsReq": {
            "type": "object",
            "required": [
                "max_results",
                "topic"
            ],
            "properties": {
                "language": {
                    "type": "string"
                },
                "max_results": {
                    "type": "integer"
                },
                "topic": {
                    "type": "string"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
