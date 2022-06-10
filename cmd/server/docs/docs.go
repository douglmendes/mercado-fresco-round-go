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
        "termsOfService": "https://developers.mercadolivre.com.br/pt_br/termos-e-condicoes",
        "contact": {
            "name": "API Support",
            "url": "https://developers.mercadolivre.com.br/support"
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
        "/api/v1/employees": {
            "get": {
                "description": "get employees",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "List employees",
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
                "description": "Create a new employee in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Create a new employee",
                "parameters": [
                    {
                        "description": "Employee to be created",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.requestEmployee"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/employees.Employee"
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
        "/api/v1/employees/{id}": {
            "get": {
                "description": "Get a employee from the system searching by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Get a employee by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/employees.Employee"
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
                "description": "Delete a employee from the system, selecting by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Delete a employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Successfully deleted"
                    },
                    "400": {
                        "description": "Bad Request",
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
            "patch": {
                "description": "update employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Update employee",
                "parameters": [
                    {
                        "description": "Employee to create",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.requestEmployee"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/employees.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
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
            }
        },
        "/api/v1/sections": {
            "get": {
                "description": "List all sections currently in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "List all sections",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/sections.Section"
                            }
                        }
                    },
                    "204": {
                        "description": "Empty database"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new section in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "Create a new section",
                "parameters": [
                    {
                        "description": "Section to be created",
                        "name": "section",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.sectionsRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/sections.Section"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/sections/{id}": {
            "get": {
                "description": "Get a section from the system searching by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "Get a section by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Section id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sections.Section"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a section from the system, selecting by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "Delete a section",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Section id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Successfully deleted"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a section in the system, selecting by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "Update a section",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Section id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Section to be updated (all fields are optional)",
                        "name": "section",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controllers.sectionsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sections.Section"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/sellers.Seller"
                            }
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
                        "name": "seller",
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
                            "$ref": "#/definitions/sellers.Seller"
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
                            "$ref": "#/definitions/sellers.Seller"
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
                            "$ref": "#/definitions/sellers.Seller"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
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
                        "description": "desc",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/response.Response"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "$ref": "#/definitions/warehouses.Warehouse"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
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
                            "$ref": "#/definitions/warehouses.Warehouse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
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
                            "$ref": "#/definitions/warehouses.Warehouse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
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
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
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
                            "$ref": "#/definitions/warehouses.Warehouse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
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
        "controllers.requestEmployee": {
            "type": "object",
            "properties": {
                "card_number_id": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "warehouse_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.sectionsRequest": {
            "type": "object",
            "required": [
                "current_capacity",
                "current_temperature",
                "maximum_capacity",
                "minimum_capacity",
                "minimum_temperature",
                "product_type_id",
                "section_number",
                "warehouse_id"
            ],
            "properties": {
                "current_capacity": {
                    "type": "integer"
                },
                "current_temperature": {
                    "type": "integer"
                },
                "maximum_capacity": {
                    "type": "integer"
                },
                "minimum_capacity": {
                    "type": "integer"
                },
                "minimum_temperature": {
                    "type": "integer"
                },
                "product_type_id": {
                    "type": "integer"
                },
                "section_number": {
                    "type": "integer"
                },
                "warehouse_id": {
                    "type": "integer"
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
        },
        "employees.Employee": {
            "type": "object",
            "properties": {
                "card_number_id": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "warehouse_id": {
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                }
            }
        },
        "sections.Section": {
            "type": "object",
            "properties": {
                "current_capacity": {
                    "type": "integer"
                },
                "current_temperature": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "maximum_capacity": {
                    "type": "integer"
                },
                "minimum_capacity": {
                    "type": "integer"
                },
                "minimum_temperature": {
                    "type": "integer"
                },
                "product_type_id": {
                    "type": "integer"
                },
                "section_number": {
                    "type": "integer"
                },
                "warehouse_id": {
                    "type": "integer"
                }
            }
        },
        "sellers.Seller": {
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
                "id": {
                    "type": "integer"
                },
                "telephone": {
                    "type": "string"
                }
            }
        },
        "warehouses.Warehouse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
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
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Mercado Fresco",
	Description:      "This API Handle MELI fresh products",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
