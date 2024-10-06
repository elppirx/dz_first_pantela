package handlers

import (
	"dz_first_pantela/iternal/messagesService"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *messagesService.MessageService
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

func (h *Handler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	message, err := h.Service.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(message)
}

func (h *Handler) UpdateMessageById(w http.ResponseWriter, r *http.Request) {
	ids := mux.Vars(r)
	id, err := strconv.Atoi(ids["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var message messagesService.Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateMessage, err := h.Service.UpdateMessageById(id, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateMessage)
}

func (h *Handler) DeleteMessageById(w http.ResponseWriter, r *http.Request) {
	ids := mux.Vars(r)
	id, err := strconv.Atoi(ids["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.Service.DeleteMessageById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(result)
}
