package dao

import (
	"cherryang.com/models"
)

func CreateTopic(name string) {
	topic := models.Topic{Name:name}
	result := db.Create(&topic)
	resultProcess(result)
}
