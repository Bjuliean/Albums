package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Album struct {
	ID		int			`json:"id" db:"id"`
	Title	string		`json:"title" db:"title"`
	Artist	string		`json:"artist" db:"artist"`
	Price	float64		`json:"price" db:"price"`
}

type Storage interface {
	ConnectToDatabase()
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

func (p *PostgresStorage)ConnectToDatabase() {
	db, err := sql.Open("postgres", "user=user password=qwerty dbname=user sslmode=disable")
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
	p.DeleteAlbum(id)
	p.CreateAlbum(al)
}

