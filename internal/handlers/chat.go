package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Aidajy111/API-chats-and-messages/internal/models"
	"github.com/Aidajy111/API-chats-and-messages/internal/repo"
	"github.com/gorilla/mux"
)

type ChatHandler struct {
	ChatRepo    *repo.ChatRepository
	MessageRepo *repo.MessageRepository
}

func NewChatHandler(chatRepo *repo.ChatRepository, msgRepo *repo.MessageRepository) *ChatHandler {
	return &ChatHandler{ChatRepo: chatRepo, MessageRepo: msgRepo}
}

// Создание чата - post /chats
func (h *ChatHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload.Title = strings.TrimSpace(payload.Title)

	if len(payload.Title) == 0 {
		http.Error(w, "There is no chat header", http.StatusBadRequest)
		return
	}

	chat := &models.Chat{Title: payload.Title}
	if err := h.ChatRepo.Create(chat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(chat)
}

// получить чат с лимитом /chats/{id}?limit=N
func (h *ChatHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	limit := 20 // default
	if l := r.URL.Query().Get("limit"); l != "" {
		// проверка лимита в рамках 1 до 100
		if val, err := strconv.Atoi(l); err == nil && val > 0 && val <= 100 {
			limit = val
		}
	}

	chat, messages, err := h.ChatRepo.GetByID(id, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := struct {
		Chat     *models.Chat      `json:"chat"`
		Messages []models.Messeges `json:"messages"`
	}{
		Chat:     chat,
		Messages: messages,
	}

	json.NewEncoder(w).Encode(resp)
}

// Удалить /chats/{id}
func (h *ChatHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if err := h.ChatRepo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// отправить сообщение в чат  /chats/{id}/messages/
func (h *ChatHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	exists, err := h.ChatRepo.Check(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "chat not found", http.StatusNotFound)
		return
	}

	var payload struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(payload.Text) == 0 {
		http.Error(w, "text cannot be empty", http.StatusBadRequest)
		return
	}

	message := &models.Messeges{
		ChatID: id,
		Text:   payload.Text,
	}

	if err := h.MessageRepo.Create(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}
