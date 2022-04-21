package main

import (
	"github.com/AniketKariya/go-client-person/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	db *gorm.DB
}

func GetClient() Client {
	var client Client
	var err error

	client.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("failed to connect to database")
	}

	client.db.AutoMigrate(&models.Person{})

	return client
}

func (c Client) Create(person models.Person) (uint, error) {
	person1 := person
	result := c.db.Create(&person1)
	if result.Error != nil {
		return 0, result.Error
	}
	return person1.ID, nil
}

func (c Client) Read(id uint) (models.Person, error) {
	var person models.Person
	result := c.db.First(&person, id)
	return person, result.Error
}

func (c Client) Update(id uint, newPerson models.Person) (models.Person, error) {
	var person models.Person
	c.db.First(&person, id)

	result := c.db.Model(&person).Updates(newPerson)
	return person, result.Error
}

func (c Client) Delete(id uint) (int64, error) {
	result := c.db.Delete(&models.Person{}, id)
	return result.RowsAffected, result.Error
}
