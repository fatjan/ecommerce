package service

import (
	"myproject/project/entity"
	"myproject/project/repository"
)

// IUserConfigService this is interface
type IUserConfigService interface {
	CreateOneService(body *entity.User) error
	GetOneService(id string) (*entity.User, error)
	DeleteService(id string) error
	UpdateService(id string, body *entity.User) error
}

// UserConfigService : struct of User config service
type UserConfigService struct {
	Repository repository.IUserConfigRepository
}

// NewUserConfigService : create new User config service
func NewUserConfigService(r *repository.UserConfigRepository) *UserConfigService {
	return &UserConfigService{r}
}

// CreateOneService : Service to create user
func (h *UserConfigService) CreateOneService(body *entity.User) error {

	err := h.Repository.CreateOneRepository(body)

	if err != nil {
		return err
	}
	return nil
}

// GetOneService : Service to get user
func (h *UserConfigService) GetOneService(id string) (*entity.User, error) {
	getOne, err := h.Repository.GetOneRepository(id)

	if err != nil {
		return nil, err
	}
	return getOne, nil
}

// DeleteService : Service to get user
func (h *UserConfigService) DeleteService(id string) error {
	err := h.Repository.DeleteRepository(id)

	if err != nil {
		return err
	}
	return nil
}

// UpdateService : Service to create user
func (h *UserConfigService) UpdateService(id string, body *entity.User) error {

	err := h.Repository.UpdateRepository(id, body)

	if err != nil {
		return err
	}
	return nil
}
