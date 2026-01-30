package repo

import (
	"github.com/Aidajy111/API-chats-and-messages/internal/models"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessagesRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(msg *models.Messeges) error {
	return r.db.Create(msg).Error
}
