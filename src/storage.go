package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host		string
	Port		uint
	User		string
	Password	string
	DBName		string
	SSLMode		string
}

type Album struct {
	ID		int			`json:"id" db:"id"`
	Title	string		`json:"title" db:"title"`
	Artist	string		`json:"artist" db:"artist"`
	Price	float64		`json:"price" db:"price"`
}

type Storage interface {
	ConnectToDatabase(dbcfg *DBConfig)
	CloseConnection()
	CreateAlbum(al *Album)
	GetAlbums() []Album
	GetAlbum(id int) *Album
	UpdateAlbum(id int, al *Album)
	DeleteAlbum(id int)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (p *PostgresStorage)ConnectToDatabase(dbcfg *DBConfig) {
	datasrc := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=%s", dbcfg.Host, dbcfg.Port, dbcfg.User, dbcfg.Password, dbcfg.DBName, dbcfg.SSLMode)
	db, err := sql.Open("postgres", datasrc)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	p.db = db
}

func (p *PostgresStorage)CloseConnection() {
	p.db.Close()
}

func (p *PostgresStorage)CreateAlbum(al *Album) {
	p.db.Exec("INSERT INTO albums(id, title, artist, price) VALUES($1,$2,$3,$4);", al.ID, al.Title, al.Artist, al.Price)
}

func (p *PostgresStorage)GetAlbums() []Album {
	var res []Album

	rows, err := p.db.Query("SELECT id, title, artist, price FROM albums;")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var singleAlbum Album
		if err := rows.Scan(&singleAlbum.ID, &singleAlbum.Title, &singleAlbum.Artist, &singleAlbum.Price); err != nil {
			panic(err.Error())
		}
		res = append(res, singleAlbum)
		fmt.Println(singleAlbum)
	}
	return res
}

func (p *PostgresStorage)GetAlbum(id int) *Album {
	var res Album
	row, err := p.db.Query("SELECT id, title, artist, price FROM albums WHERE id=$1;", id)
	if err != nil {
		panic(err.Error())
	}
	for row.Next() {
		row.Scan(&res.ID, &res.Title, &res.Artist, &res.Price)
	}
	return &res
}

func (p *PostgresStorage)DeleteAlbum(id int) {
	p.db.Exec("DELETE FROM albums WHERE id=$1;", id)
}

func (p *PostgresStorage)UpdateAlbum(id int, al *Album) {
	al.ID = id
	p.DeleteAlbum(id)
	p.CreateAlbum(al)
}

