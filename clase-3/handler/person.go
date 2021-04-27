package handler

import (
	"curso_apis_go/clase-3/model"
	"encoding/json"
	"net/http"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Persona creada correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}
