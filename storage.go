package main

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type Storage struct {
	conn *sql.DB
}

func NewStorage(conn *sql.DB) *Storage {
	return &Storage{
		conn: conn}
}

func (s *Storage) SaveProgram(program Program) error {
	query := "INSERT INTO programs(id, name, nameEn, isPublic, projectID) VALUES ($1, $2, $3, $4, $5)"
	_, err := s.conn.Exec(query, program.ID, program.Name, program.NameEn, program.IsPublic, program.ProjectID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) ProgramByID(id uuid.UUID) (Program, error) {
	query := "SELECT id, name, nameEn, isPublic, projectID FROM programs WHERE id = $1 "

	var program Program
	err := s.conn.QueryRow(query, id).Scan(&program.ID, &program.Name, &program.NameEn, &program.IsPublic, &program.ProjectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Program{}, ErrNotFound
		}
		return Program{}, err
	}
	return program, nil
}
