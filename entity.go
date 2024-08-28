package main

import (
	"errors"
	"github.com/google/uuid"
)

type Program struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	NameEn    string    `json:"nameEn"`
	IsPublic  bool      `json:"isPublic"`
	ProjectID uuid.UUID `json:"projectID"`
}

var (
	ErrNotFound = errors.New("not found")
)
