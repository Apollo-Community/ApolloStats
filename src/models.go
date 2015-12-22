package apollostats

import (
	"strings"
	"time"
)

type Ban struct {
	ID         int64     `gorm:"column:id;primary_key"`
	Timestamp  time.Time `gorm:"column:bantime"`
	CKey       string    `gorm:"column:ckey"`
	CID        string    `gorm:"column:computerid"`
	IP         string    `gorm:"column:ip"`
	Bantype    string    `gorm:"column:bantype"`
	Admin      string    `gorm:"column:a_ckey"`
	Reason     string    `gorm:"column:reason"`
	Duration   int       `gorm:"column:duration"`
	Expiration time.Time `gorm:"column:expiration_time"`
}

func (b *Ban) TableName() string {
	return "ban"
}

func (b *Ban) Expires() string {
	if b.Duration < 0 {
		return "Permanent"
	} else {
		return b.Expiration.Format("2006-01-02 15:04 MST")
	}
}

type AccountItem struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Timestamp time.Time `gorm:"column:time"`
	CKey      string    `gorm:"column:ckey"`
	Item      string    `gorm:"column:item"`
}

func (a *AccountItem) TableName() string {
	return "acc_items"
}

type Death struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Timestamp time.Time `gorm:"column:tod"`
	Name      string    `gorm:"column:name"`
	Job       string    `gorm:"column:job"`
	Room      string    `gorm:"column:pod"`
	Position  string    `gorm:"column:coord"`
	Brute     int       `gorm:"column:bruteloss"`
	Brain     int       `gorm:"column:brainloss"`
	Fire      int       `gorm:"column:fireloss"`
	Oxygen    int       `gorm:"column:oxyloss"`
}

func (d *Death) TableName() string {
	return "death"
}

func (d *Death) RoomName() string {
	// Cleanup the room name
	return strings.Trim(d.Room, "Ã¿")
}
