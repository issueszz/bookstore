package model

type BookStore struct {
	Book string `gorm:"size:255; primaryKey; not null; comment:book name"`
	Price int64 `gorm:"not null; comment:book price"`
}
