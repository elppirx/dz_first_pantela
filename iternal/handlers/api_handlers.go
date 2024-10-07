package handlers

import (
	"context"
	"dz_first_pantela/iternal/messagesService"
	"dz_first_pantela/iternal/web/messages"
)

type Handler struct {
	Service *messagesService.MessageService
}

func (h *Handler) GetMessages(_ context.Context, request messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:   &msg.ID,
			Text: &msg.Text,
		}
		response = append(response, message)
	}
	return response, nil
}

func (h *Handler) PostMessages(ctx context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToCreate := messagesService.Message{Text: *messageRequest.Text}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:   &createdMessage.ID,
		Text: &createdMessage.Text,
	}

	return response, nil
}

func (h *Handler) DeleteMessagesId(ctx context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	id := request.Id

	result, err := h.Service.DeleteMessageById(int(id))
	if err != nil {
		return nil, err
	}
	return messages.DeleteMessagesId200JSONResponse(result), nil
}

func (h *Handler) PutMessagesId(ctx context.Context, request messages.PutMessagesIdRequestObject) (messages.PutMessagesIdResponseObject, error) {
	id := request.Id

	messageRequest := request.Body
	messageToUpdate := messagesService.Message{
		Text: *messageRequest.Text,
	}

	updatedMessage, err := h.Service.UpdateMessageById(int(id), messageToUpdate)
	if err != nil {
		return nil, err
	}

	response := messages.PutMessagesId201JSONResponse{
		Id:   &updatedMessage.ID,
		Text: &updatedMessage.Text,
	}

	return response, nil
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

//func (h *Handler) UpdateMessageById(w http.ResponseWriter, r *http.Request) {
//	ids := mux.Vars(r)
//	id, err := strconv.Atoi(ids["id"])
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	var message messagesService.Message
//	err = json.NewDecoder(r.Body).Decode(&message)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	updateMessage, err := h.Service.UpdateMessageById(id, message)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(updateMessage)
//}
//func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
//	var message messagesService.Message
//	err := json.NewDecoder(r.Body).Decode(&message)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	createdMessage, err := h.Service.CreateMessage(message)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(createdMessage)
//}
//
//func (h *Handler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
//	message, err := h.Service.GetAllMessages()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.Header().Set("Content-Type", "aplication/json")
//	json.NewEncoder(w).Encode(message)
//}
//
//func (h *Handler) DeleteMessageById(w http.ResponseWriter, r *http.Request) {
//	ids := mux.Vars(r)
//	id, err := strconv.Atoi(ids["id"])
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	result, err := h.Service.DeleteMessageById(id)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.Header().Set("Content-Type", "aplication/json")
//	json.NewEncoder(w).Encode(result)
//}
