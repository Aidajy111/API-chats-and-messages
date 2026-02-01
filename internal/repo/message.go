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

func (r *MessageRepository) Create(msg *models.Messages) error {
	return r.db.Create(msg).Error
}

func (r *ChatRepository) Check(id int) (bool, error) {
	var count int64
	err := r.db.Model(&models.Chat{}).
		Where("id = ?", id).
		Count(&count).Error

	return count > 0, err
}
