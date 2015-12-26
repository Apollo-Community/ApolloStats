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
	RoundID   int64     `gorm:"column:round_id"`
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
	return "deaths"
}

func (d *Death) RoomName() string {
	// TODO: fix this thing in the byond source.
	// Cleanup the room name
	return strings.Trim(d.Room, "ÿ")
}

type RoundAntags struct {
	ID      int64  `gorm:"column:id"`
	RoundID int64  `gorm:"column:round_id"`
	Name    string `gorm:"column:name"`
	Job     string `gorm:"column:job"`
	Role    string `gorm:"column:role"`
	Success bool   `gorm:"column:success"`
}

func (r *RoundAntags) TableName() string {
	return "round_antags"
}

type RoundAILaws struct {
	ID      int64  `gorm:"column:id"`
	RoundID int64  `gorm:"column:round_id"`
	Law     string `gorm:"column:law"`
}

func (r *RoundAILaws) TableName() string {
	return "round_ai_laws"
}

type RoundStats struct {
	ID       int64     `gorm:"column:id"`
	Gamemode string    `gorm:"column:game_mode"`
	Endtime  time.Time `gorm:"column:end_time"`
	Duration int64     `gorm:"column:duration"`

	Productivity    int `gorm:"column:productivity"`
	Deaths          int `gorm:"column:deaths"`
	Clones          int `gorm:"column:clones"`
	DispenseVolume  int `gorm:"column:dispense_volume"`
	BombsExploded   int `gorm:"column:bombs_exploded"`
	Vended          int `gorm:"column:vended"`
	RunDistance     int `gorm:"column:run_distance"`
	BloodMopped     int `gorm:"column:blood_mopped"`
	DamageCost      int `gorm:"column:damage_cost"`
	BreakTime       int `gorm:"column:break_time"`
	MonkeyDeaths    int `gorm:"column:monkey_deaths"`
	SpamBlocked     int `gorm:"column:spam_blocked"`
	PeopleSlipped   int `gorm:"column:people_slipped"`
	DoorsOpened     int `gorm:"column:doors_opened"`
	GunsFired       int `gorm:"column:guns_fired"`
	BeepskyBeatings int `gorm:"column:beepsky_beatings"`
	DoorsWelded     int `gorm:"column:doors_welded"`
	Totalkwh        int `gorm:"column:total_kwh"`
	Artifacts       int `gorm:"column:artifacts"`
	CargoProfit     int `gorm:"column:cargo_profit"`
	TrashVented     int `gorm:"column:trash_vented"`
	AIFollow        int `gorm:"column:ai_follow"`
	Banned          int `gorm:"column:banned"`
}

func (r *RoundStats) TableName() string {
	return "round_stats"
}
