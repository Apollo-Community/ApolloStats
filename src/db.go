package apollostats

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const MAX_ROWS = 200

// DB connection timeout, in seconds
const TIMEOUT = 30

// NOTE: DON'T USE ANY WRITE OPERATIONS ON THE DATABASE!
// We're interfacing with an external, live game database!

type DB struct {
	*gorm.DB
}

// Matches the Writer interface, it will do nothing when Write() is called.
type NullWriter struct {
}

// Do nothing and return n=0 and err=nil.
func (w *NullWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Connect to the DB using a dialer with timeout and deadline set.
func timeout_dial(addr string) (net.Conn, error) {
	t := time.Duration(TIMEOUT) * time.Second
	c, e := net.DialTimeout("tcp", addr, t)
	if e != nil {
		return nil, e
	}

	d := time.Now().Add(t)
	c.SetDeadline(d) // Deadline for read/write
	return c, nil
}

func OpenDB(DSN string, debug bool) (*DB, error) {
	// Register a custom dialer so we can set a connection timeout and deadline
	mysql.RegisterDial("tcp", timeout_dial)

	// HACK: The mysql driver keeps spamming i/o timeouts in the logs when
	// the db connection times out, so let's disable all mysql logging...
	l := log.New(&NullWriter{}, "", 0)
	mysql.SetLogger(l)

	tmp := fmt.Sprintf("%s?parseTime=True&timeout=30s", DSN)
	db, e := gorm.Open("mysql", tmp)
	// Avoid setting LogMode(true), since that will trace log all queries.
	// Only show errors by default, unless silenced completely by debug = false
	if !debug {
		db.LogMode(debug)
	}
	return &DB{&db}, e
}

type Stats struct {
	TotalAccountItems  int64
	TotalBans          int64
	AvgBans            float64
	TotalRounds        int64
	TotalDeaths        int64
	AvgDeaths          float64
	TotalRoundDuration float64
	AvgRoundDuration   float64
	TotalMonkeys       int64
	TotalDamages       int64
}

func (db *DB) GetStats() *Stats {
	var (
		total_acc_items      int64
		total_bans           int64
		avg_bans             float64
		total_rounds         int64
		total_deaths         int64
		avg_deaths           float64
		total_round_duration float64
		avg_round_duration   float64
		total_monkey_deaths  int64
		total_damage_cost    int64
	)

	db.Model(AccountItem{}).Count(&total_acc_items)
	db.Model(Ban{}).Count(&total_bans)
	db.Model(RoundStats{}).Count(&total_rounds)
	db.Model(Death{}).Count(&total_deaths)

	var b Ban
	db.First(&b, 1)
	ban_days := time.Now().Sub(b.Timestamp).Hours() / 24
	avg_bans = float64(total_bans) / float64(ban_days)

	var durations []int64
	var total_minutes int64
	db.Model(RoundStats{}).Pluck("duration", &durations)
	for _, d := range durations {
		total_minutes += d
	}
	avg_round_duration = float64(total_minutes) / float64(total_rounds)
	total_round_duration = float64(total_minutes) / 60.0

	var monkeys []int64
	db.Model(RoundStats{}).Pluck("monkey_deaths", &monkeys)
	for _, m := range monkeys {
		total_monkey_deaths += m
	}

	var damages []int64
	db.Model(RoundStats{}).Pluck("damage_cost", &damages)
	for _, d := range damages {
		total_damage_cost += d
	}

	avg_deaths = float64(total_deaths+total_monkey_deaths) / float64(total_rounds)

	return &Stats{total_acc_items, total_bans, avg_bans, total_rounds, total_deaths, avg_deaths, total_round_duration, avg_round_duration, total_monkey_deaths, total_damage_cost}
}

func (db *DB) AllBans() []*Ban {
	var tmp []*Ban
	db.Order("id desc").Limit(MAX_ROWS).Find(&tmp)
	return tmp
}

func (db *DB) AllAccountItems() []*AccountItem {
	var tmp []*AccountItem
	db.Order("id desc").Limit(MAX_ROWS).Find(&tmp)
	return tmp
}

func (db *DB) AllDeaths() []*Death {
	var tmp []*Death
	db.Order("id desc").Limit(MAX_ROWS).Find(&tmp)
	return tmp
}

func (db *DB) AllRounds() []*RoundStats {
	var tmp []*RoundStats
	db.Order("id desc").Limit(MAX_ROWS).Find(&tmp)
	return tmp
}

func (db *DB) GetRound(id int64) *RoundStats {
	var tmp RoundStats
	db.First(&tmp, id)
	return &tmp
}

func (db *DB) GetLatestRound() *RoundStats {
	var tmp RoundStats
	db.Order("id desc").Limit(1).First(&tmp)
	return &tmp
}

func (db *DB) GetAntags(id int64) []*RoundAntags {
	var tmp []*RoundAntags
	db.Order("round_id desc, name asc").Where("round_id = ?", id).Find(&tmp)
	return tmp
}

func (db *DB) GetAILaws(id int64) []*RoundAILaws {
	var tmp []*RoundAILaws
	db.Order("round_id desc, law asc").Where("round_id = ?", id).Find(&tmp)
	return tmp
}

func (db *DB) GetDeaths(id int64) []*Death {
	var tmp []*Death
	db.Order("round_id desc, name asc").Where("round_id = ?", id).Find(&tmp)
	return tmp
}

func (db *DB) GetCharacter(id int64) *Character {
	var tmp Character
	db.First(&tmp, id)
	return &tmp
}

func (db *DB) SearchCharacter(ckey, name string) []*Character {
	var tmp []*Character
	// Don't any weird behaviours if the user is smart enough to try to use these
	tckey := "%" + strings.Trim(ckey, "_%") + "%"
	tname := "%" + strings.Trim(name, "_%") + "%"
	db.Debug().Order("ckey, name desc").Where("ckey LIKE ? AND name LIKE ?", tckey, tname).Limit(MAX_ROWS).Find(&tmp)
	return tmp
}
