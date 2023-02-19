package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DocumentName string             `json:"name,omitempty"`
	DocumentType string             `json:"type,omitempty"`
	Legal        bool               `json:"islegal,omitempty"`
	Owner        *Owner             `json:"owner,omitempty"`
}

type Owner struct {
	OwnerName string `json:"name,omitempty"`
	Varified  bool   `json:"verified,omitempty"`
}
