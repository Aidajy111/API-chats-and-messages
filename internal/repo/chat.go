package repo

import (
	"github.com/Aidajy111/API-chats-and-messages/internal/models"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) Create(chat *models.Chat) error {
	return r.db.Create(chat).Error
}

func (r *ChatRepository) GetByID(id int, limit int) (*models.Chat, []models.Messages, error) {
	var chat models.Chat
	if err := r.db.First(&chat, "id = ?", id).Error; err != nil {
		return nil, nil, err
	}

	var msgs []models.Messages

	query := r.db.
		Where("chat_id = ?", id).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&msgs).Error; err != nil {
		return nil, nil, err
	}

	for i, j := 0, len(msgs)-1; i < j; i, j = i+1, j-1 {
		msgs[i], msgs[j] = msgs[j], msgs[i]
	}

	return &chat, msgs, nil
}
func (r *ChatRepository) Delete(id int) error {
	return r.db.Delete(&models.Chat{}, id).Error
}
