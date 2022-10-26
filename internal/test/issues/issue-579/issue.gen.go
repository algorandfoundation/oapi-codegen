// Package issue_579 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package issue_579

import (
	openapi_types "github.com/algorand/oapi-codegen/pkg/types"
)

// AliasedDate defines model for AliasedDate.
type AliasedDate = openapi_types.Date

// Pet defines model for Pet.
type Pet struct {
	Born   *AliasedDate        `json:"born,omitempty"`
	BornAt *openapi_types.Date `json:"born_at,omitempty"`
}
