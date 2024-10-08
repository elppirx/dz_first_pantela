package usersService

type UserService struct {
	repo userRepository
}

func NewUsersService(repo userRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUserById(id int, user User) (User, error) {
	result, err := s.repo.UpdateUserById(id, user)
	if err != nil {
		return User{}, err
	}
	return result, nil
}

func (s *UserService) DeleteUserById(id int) (string, error) {
	result, err := s.repo.DeleteUserById(id)
	if err != nil {
		return "", err
	}
	return result, nil
}
