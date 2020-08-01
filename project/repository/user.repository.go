package repository

import (
	"log"
	"myproject/project/entity"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// IUserConfigRepository this is interface for repository methode
type IUserConfigRepository interface {
	CreateOneRepository(body *entity.User) error
	GetOneRepository(id string) (*entity.User, error)
	DeleteRepository(id string) error
	UpdateRepository(id string, body *entity.User) error
}

// UserConfigRepository : struct of User config repository
type UserConfigRepository struct{}

// ConnectDatabase is to make a connection to the database
func ConnectDatabase() (*xorm.Engine, error) {
	user := "user"
	password := "password"
	engine, err := xorm.NewEngine("mysql", user+":"+password+"@/mydb?charset=utf8")
	err = engine.Sync2(new(entity.User))
	if err != nil {
		log.Println("Connection Failed to Open")
		return nil, err
	}
	err = engine.Sync2(new(entity.User))
	return engine, nil
}

func NewUserConfigRepository() *UserConfigRepository {
	return &UserConfigRepository{}
}

// CreateOneRepository : Repository to create tax config
func (h *UserConfigRepository) CreateOneRepository(body *entity.User) error {
	db, _ := ConnectDatabase()
	_, err := db.Insert(body)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserConfigRepository) GetOneRepository(id string) (*entity.User, error) {
	db, _ := ConnectDatabase()
	var user entity.User
	_, err := db.Where("id = ?", id).Get(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (h *UserConfigRepository) DeleteRepository(id string) error {
	db, _ := ConnectDatabase()
	var user entity.User
	_, err := db.Where("id = ?", id).Delete(&user)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneRepository : Repository to create tax config
func (h *UserConfigRepository) UpdateRepository(id string, body *entity.User) error {
	db, _ := ConnectDatabase()
	_, err := db.Where("id = ?", id).Update(body)
	if err != nil {
		return err
	}
	return nil
}
