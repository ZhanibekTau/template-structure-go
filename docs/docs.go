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
        "/api/v1/products/getAllProducts": {
            "get": {
                "description": "Get products to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get All products",
                "operationId": "get-products",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/products/insertProducts": {
            "post": {
                "description": "Add new products to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Insert products",
                "operationId": "insert-products",
                "parameters": [
                    {
                        "description": "List of products to insert",
                        "name": "products",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ProductList"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Product": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "discountPercentage": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                },
                "thumbnail": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ProductList": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
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
	Title:            "Subbox API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
