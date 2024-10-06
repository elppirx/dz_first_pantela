package messagesService

type MessageService struct {
	repo messageRepository
}

func NewService(repo messageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

func (s *MessageService) UpdateMessageById(id int, message Message) (Message, error) {
	result, err := s.repo.UpdateMessageById(id, message)
	if err != nil {
		return Message{}, err
	}
	return result, nil
}

func (s *MessageService) DeleteMessageById(id int) (string, error) {
	result, err := s.repo.DeleteMessageById(id)
	if err != nil {
		return "", err
	}
	return result, nil
}
