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
        "/auth/login": {
            "post": {
                "description": "Autentificar usuario",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuario"
                ],
                "summary": "Login Usuario",
                "parameters": [
                    {
                        "description": "Autentificar Usuario",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UsuariosLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Guardar nuevo usuario",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuario"
                ],
                "summary": "Registrar Usuario",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UsuariosRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/list": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Listado de usuarios",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuario"
                ],
                "summary": "Listado de Usuarios",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "types.UsuariosLogin": {
            "type": "object",
            "required": [
                "Correo_electronico",
                "password"
            ],
            "properties": {
                "Correo_electronico": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "types.UsuariosRegister": {
            "type": "object",
            "required": [
                "Correo_electronico",
                "apellidos",
                "celular",
                "codigo_pais",
                "fecha_nacimiento",
                "genero",
                "nacionalidad",
                "nombre",
                "password"
            ],
            "properties": {
                "Correo_electronico": {
                    "type": "string"
                },
                "apellidos": {
                    "type": "string"
                },
                "celular": {
                    "type": "string"
                },
                "codigo_pais": {
                    "type": "string"
                },
                "fecha_nacimiento": {
                    "type": "string"
                },
                "genero": {
                    "type": "string"
                },
                "nacionalidad": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Api urls for Unity Application",
	Description:      "List of all api services",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
