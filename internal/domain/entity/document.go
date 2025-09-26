package entity

import "time"

type Document struct {
	Hash      string    `db:"hash"`
	MetaData  string    `db:"meta_data"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
