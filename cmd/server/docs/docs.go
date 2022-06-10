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
        "/api/v1/sellers": {
            "get": {
                "description": "get sellers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sellers"
                ],
                "summary": "List sellers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.request"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create sellers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sellers"
                ],
                "summary": "Create sellers",
                "parameters": [
                    {
                        "description": "Seller to create",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sellers/{id}": {
            "get": {
                "description": "get seller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sellers"
                ],
                "summary": "List seller",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Seller ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.request"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete seller",
                "tags": [
                    "Sellers"
                ],
                "summary": "Delete seller",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Seller ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/controllers.request"
                        }
                    }
                }
            },
            "patch": {
                "description": "update seller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sellers"
                ],
                "summary": "Update seller",
                "parameters": [
                    {
                        "description": "Seller to create",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.request"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Seller ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.request"
                        }
                    }
                }
            }
        },
        "/api/v1/warehouses": {
            "get": {
                "description": "List all available warehouses",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Warehouses"
                ],
                "summary": "List warehouses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.whRequest"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create one warehouse",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Warehouses"
                ],
                "summary": "Create warehouses",
                "parameters": [
                    {
                        "description": "Warehouse to create",
                        "name": "warehouses",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.whRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/warehouses/{id}": {
            "get": {
                "description": "Read one warehouse",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Warehouses"
                ],
                "summary": "Warehouse",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Warehouse ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a warehouse by ID",
                "tags": [
                    "Warehouses"
                ],
                "summary": "Delete warehouse",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Warehouse ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a warehouse by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Warehouses"
                ],
                "summary": "Update warehouse",
                "parameters": [
                    {
                        "description": "Warehouse to update",
                        "name": "warehouse",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.whRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Warehouse ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
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
        "controllers.request": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "cid": {
                    "type": "integer"
                },
                "company_name": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                }
            }
        },
        "controllers.whRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "minimun_capacity": {
                    "type": "integer"
                },
                "minimun_temperature": {
                    "type": "integer"
                },
                "telephone": {
                    "type": "string"
                },
                "warehouse_code": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
