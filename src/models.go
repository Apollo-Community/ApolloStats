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
	BanType    string    `gorm:"column:bantype"`
	BannedJob  string    `gorm:"column:job"`
	Admin      string    `gorm:"column:a_ckey"`
	Reason     string    `gorm:"column:reason"`
	Duration   int64     `gorm:"column:duration"`
	Expiration time.Time `gorm:"column:expiration_time"`
}

func (b *Ban) TableName() string {
	return "ban"
}

func (b *Ban) Ban() string {
	if b.BanType == "PERMABAN" || b.BanType == "TEMPBAN" {
		return "Server"
	} else if b.BanType == "JOB_PERMABAN" || b.BanType == "JOB_TEMPBAN" {
		return strings.Title(b.BannedJob)
	}
	return strings.Title(b.BanType)
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
	Brute     int64     `gorm:"column:bruteloss"`
	Brain     int64     `gorm:"column:brainloss"`
	Fire      int64     `gorm:"column:fireloss"`
	Oxygen    int64     `gorm:"column:oxyloss"`
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
	GameMode string    `gorm:"column:game_mode"`
	EndTime  time.Time `gorm:"column:end_time"`
	Duration int64     `gorm:"column:duration"`

	Productivity    int64 `gorm:"column:productivity"`
	Deaths          int64 `gorm:"column:deaths"`
	Clones          int64 `gorm:"column:clones"`
	DispenseVolume  int64 `gorm:"column:dispense_volume"`
	BombsExploded   int64 `gorm:"column:bombs_exploded"`
	Vended          int64 `gorm:"column:vended"`
	RunDistance     int64 `gorm:"column:run_distance"`
	BloodMopped     int64 `gorm:"column:blood_mopped"`
	DamageCost      int64 `gorm:"column:damage_cost"`
	BreakTime       int64 `gorm:"column:break_time"`
	MonkeyDeaths    int64 `gorm:"column:monkey_deaths"`
	SpamBlocked     int64 `gorm:"column:spam_blocked"`
	PeopleSlipped   int64 `gorm:"column:people_slipped"`
	DoorsOpened     int64 `gorm:"column:doors_opened"`
	GunsFired       int64 `gorm:"column:guns_fired"`
	BeepskyBeatings int64 `gorm:"column:beepsky_beatings"`
	DoorsWelded     int64 `gorm:"column:doors_welded"`
	Totalkwh        int64 `gorm:"column:total_kwh"`
	Artifacts       int64 `gorm:"column:artifacts"`
	CargoProfit     int64 `gorm:"column:cargo_profit"`
	TrashVented     int64 `gorm:"column:trash_vented"`
	AIFollow        int64 `gorm:"column:ai_follow"`
	Banned          int64 `gorm:"column:banned"`
}

func (r *RoundStats) TableName() string {
	return "round_stats"
}
