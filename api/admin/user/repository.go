package user

import (
	"boilerplate-api/api/user/user"
	"boilerplate-api/lib/config"
	"gorm.io/gorm"
)

// Repository database structure
type Repository struct {
	db     *config.Database
	logger config.Logger
}

// NewRepository creates a new CUser repository
func NewRepository(db *config.Database, logger config.Logger) Repository {
	return Repository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (c Repository) WithTrx(trxHandle *gorm.DB) Repository {
	if trxHandle == nil {
		c.logger.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// Create user
func (c Repository) Create(User user.CUser) error {
	return c.db.DB.Create(&User).Error
}

// GetAllUsers Get All users
func (c Repository) GetAllUsers(pagination Pagination) (users []GetUserResponse, count int64, err error) {
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&user.CUser{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`users`.`name` LIKE ?", searchQuery))
	}

	return users, count, queryBuilder.
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error
}

func (c Repository) GetOneUser(Id int64) (userModel GetUserResponse, err error) {
	return userModel, c.db.DB.
		Model(&userModel).
		Where("id = ?", Id).
		First(&userModel).
		Error
}

func (c Repository) GetOneUserWithEmail(Email string) (user user.CUser, err error) {
	return user, c.db.DB.Model(&user).
		Where("email = ?", Email).
		First(&user).
		Error
}

func (c Repository) GetOneUserWithPhone(Phone string) (user user.CUser, err error) {
	return user, c.db.DB.
		First(&user, "phone = ?", Phone).
		Error

}
