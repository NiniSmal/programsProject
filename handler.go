package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type Handler struct {
	storage *Storage
}

func NewHandler(st *Storage) *Handler {
	return &Handler{
		storage: st,
	}
}

type ProgramStorage interface {
	SaveProgram(program Program) error         //сохранение программы в бд указанного проекта
	ProgramByID(id uuid.UUID) (Program, error) //получение программы по ID
}

//получение программ из стороннего API по полученному ID

func (h *Handler) ProgramsByProjectID(w http.ResponseWriter, r *http.Request) {
	projectID := r.URL.Query().Get("projectID")
	projectID = strings.ToUpper(projectID)

	url := fmt.Sprintf("[Адрес ЕПС API]/api/programs/?projectId=%s", projectID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		_, _ = w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return

	}
	res, err := client.Do(req)
	if err != nil {
		_, _ = w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return

	}
	defer res.Body.Close()

	var programs []Program

	err = json.NewDecoder(res.Body).Decode(&programs)
	if err != nil {
		_, _ = w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	for _, program := range programs {
		err = h.storage.SaveProgram(program)
		if err != nil {
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
	}

	err = json.NewEncoder(w).Encode(programs)
	if err != nil {
		_, _ = w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}
}

// получение программы по полученному ID
func (h *Handler) ProgramByID(w http.ResponseWriter, r *http.Request) {
	programID := r.URL.Query().Get("programID")

	id, err := uuid.Parse(programID)
	if err != nil {
		_, _ = w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	program, err := h.storage.ProgramByID(id)
	if err != nil {
		_, _ = w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	err = json.NewEncoder(w).Encode(program)
	if err != nil {
		_, _ = w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}

}
