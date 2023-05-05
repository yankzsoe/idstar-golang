package helper

import (
	"ginweb/app/models"
	config "ginweb/configs"
)

var db = config.GetDB()

func Get(id int) (error, *models.Todo) {
	todo := &models.Todo{}
	if err := db.Where("ID = ?", id).Find(&todo).Error; err != nil {
		return err, nil
	}
	return nil, todo
}

func GetList() (error, []*models.Todo) {
	todos := []*models.Todo{}
	if err := db.Find(&todos).Error; err != nil {
		return err, nil
	}
	return nil, todos
}

func Save(todo models.Todo) error {
	if err := db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func Update(id int, todo models.Todo) error {
	if err := db.Where("ID = ?", id).Updates(todo).Error; err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	if err := db.Delete(&models.Todo{}, "ID = ?", id).Error; err != nil {
		return err
	}
	return nil
}
