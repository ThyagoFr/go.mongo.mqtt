package models

import (
	"time"

	"gomongo.mqtt/src/db"
)

// SensorData - Modelo de um dado recebido do BrokerMQTT
type SensorData struct {
	TimeAquisition time.Time `json:"time"`
	Value          float64   `json:"value"`
}

// InsertData - Metodo responsavel por inserir os dados no banco
func (s *SensorData) InsertData() error {
	db := db.Database{}
	db.Init()
	conn := db.Connect()
	err := conn.C("data").Insert(&s)
	db.Close()
	return err
}
