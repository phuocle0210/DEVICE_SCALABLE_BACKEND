package common

import "time"

type SqlModelDefault struct {
	Id        int        `gorm:"id;primaryKey;autoIncrement:true;unique"`
	Status    int        `gorm:"-" json:"status"`
	CreatedAt *time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updated_at"`
}

func (sql *SqlModelDefault) InsertDefaultModel() {
	now := time.Now().UTC()

	sql.Status = 1
	sql.CreatedAt = &now
	sql.UpdatedAt = &now
}
