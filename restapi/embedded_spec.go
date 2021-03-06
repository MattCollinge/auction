package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "reverse auction API",
    "title": "auction",
    "contact": {
      "name": "Matt Collinge",
      "url": "https://github.com/mattcollinge/auction",
      "email": "matt.collinge@gmail.com"
    },
    "version": "0.0.1"
  },
  "host": "localhost:3000",
  "basePath": "/",
  "paths": {
    "/auctions/auctionId/{auctionId}": {
      "x-swagger-router-controller": "auctions",
      "get": {
        "description": "Returns the auction given by the auctionId",
        "operationId": "auctionByAuctionId",
        "parameters": [
          {
            "$ref": "#/parameters/auctionId"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/AuctionResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "401": {
            "description": "Unauthorised"
          },
          "404": {
            "description": "No Auction Found"
          },
          "default": {
            "description": "Unexpected Error"
          }
        }
      }
    },
    "/private/ping": {
      "x-swagger-router-controller": "healthCheck",
      "get": {
        "description": "Returns a ` + "`" + `pong` + "`" + ` message for system health checking",
        "produces": [
          "text/html"
        ],
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/PongModel"
            }
          }
        }
      }
    },
    "/swagger": {
      "x-swagger-pipe": "swagger_raw"
    }
  },
  "definitions": {
    "AuctionModel": {
      "description": "this is the format of a Auction object",
      "type": "object",
      "required": [
        "auctionId",
        "auctionItem"
      ],
      "properties": {
        "auctionId": {
          "type": "string"
        },
        "auctionItem": {
          "type": "string"
        }
      }
    },
    "AuctionResponse": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/AuctionModel"
      }
    },
    "PongModel": {
      "description": "The result of a private ping healthcheck request",
      "type": "string"
    }
  },
  "parameters": {
    "auctionId": {
      "type": "string",
      "description": "the auctionId of an Auction",
      "name": "auctionId",
      "in": "path",
      "required": true
    }
  }
}`))
}
