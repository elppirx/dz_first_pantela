package usersService

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *userRepository) UpdateUserById(id int, user User) (User, error) {
	var oldUser User
	r.db.First(&oldUser, id)
	oldUser.Email = user.Email
	oldUser.Password = user.Password
	r.db.Save(&oldUser)
	return oldUser, nil
}

func (r *userRepository) DeleteUserById(id int) (string, error) {
	var user User
	r.db.First(&user, id)
	r.db.Delete(&user)
	return "Успешное удаление", nil
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserById(id int, user User) (User, error)
	DeleteUserById(id int) (string, error)
}
