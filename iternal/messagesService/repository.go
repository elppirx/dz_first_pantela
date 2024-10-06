package messagesService

import (
	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var message []Message
	err := r.db.Find(&message).Error
	return message, err
}

func (r *messageRepository) UpdateMessageById(id int, message Message) (Message, error) {
	var oldMessage Message
	r.db.First(&oldMessage, id)
	oldMessage.Text = message.Text
	r.db.Save(&oldMessage)
	return oldMessage, nil
}

func (r *messageRepository) DeleteMessageById(id int) (string, error) {
	var message Message
	r.db.First(&message, id)
	r.db.Delete(&message)
	return "Успешное удаление", nil
}

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageById(id int, message Message) (Message, error)
	DeleteMessageById(id int) (string, error)
}
