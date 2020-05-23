package db

import (
	"fmt"

	"gomongo.mqtt/src/config"

	mgo "gopkg.in/mgo.v2"
)

// Database - Estrutura que possui os dados do Banco de Dados
type Database struct {
	Server   string
	Port     string
	Database string
	Session  *mgo.Session
}

// Init - Inicializa os dados do banco de dados
func (d *Database) Init() {

	var config = config.Config{}
	config.Read()
	d.Server = config.Server
	d.Port = config.Port
	d.Database = config.Database

}

// Connect - Retorna uma conexao com o banco
func (d *Database) Connect() (db *mgo.Database) {

	d.Session, _ = mgo.Dial(d.Server)
	db = d.Session.DB(d.Database)
	return

}

// Close - Encerra uma conexao com o banco
func (d *Database) Close() {

	fmt.Println("Fechando conexao...")
	d.Session.Close()
	return

}
