package apollostats

import (
	"fmt"
	"log"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const MAX_ROWS = 200

// DB connection timeout, in seconds
const TIMEOUT = 30

// Need to adjust time because the main server is running GMT-5
const TIMEZONE_ADJUST = "EST"

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

	tmp := fmt.Sprintf("%s?parseTime=True&loc=%v&timeout=30s", DSN, TIMEZONE_ADJUST)
	db, e := gorm.Open("mysql", tmp)
	// Avoid setting LogMode(true), since that will trace log all queries.
	// Only show errors by default, unless silenced completely by debug = false
	if !debug {
		db.LogMode(debug)
	}
	return &DB{db}, e
}

type SpeciesCount struct {
	TotalDiona      int64
	TotalHuman      int64
	TotalMachine    int64
	TotalNucleation int64
	TotalSkrell     int64
	TotalTajara     int64
	TotalUnathi     int64
	TotalVox        int64
	TotalWryn       int64
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
	TotalPlayers       int64
	TotalCharacters    int64
	Species            *SpeciesCount
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

	var (
		total_players    int64
		total_characters int64
		total_diona      int64
		total_human      int64
		total_machine    int64
		total_nucleation int64
		total_skrell     int64
		total_tajara     int64
		total_unathi     int64
		total_vox        int64
		total_wryn       int64
	)
	// Yet another issue with gorm: can't use DISTINCT and the Count() func
	// in the same query so have to do some Row() magic.
	row := db.Table("characters").Select("COUNT(DISTINCT(ckey))").Row()
	row.Scan(&total_players)
	db.Model(Character{}).Count(&total_characters)
	db.Model(Character{}).Where("species = ?", "diona").Count(&total_diona)
	db.Model(Character{}).Where("species = ?", "human").Count(&total_human)
	db.Model(Character{}).Where("species = ?", "machine").Count(&total_machine)
	db.Model(Character{}).Where("species = ?", "nucleation").Count(&total_nucleation)
	db.Model(Character{}).Where("species = ?", "skrell").Count(&total_skrell)
	db.Model(Character{}).Where("species = ?", "tajara").Count(&total_tajara)
	db.Model(Character{}).Where("species = ?", "unathi").Count(&total_unathi)
	db.Model(Character{}).Where("species = ?", "vox").Count(&total_vox)
	db.Model(Character{}).Where("species = ?", "wryn").Count(&total_wryn)

	species := &SpeciesCount{
		total_diona, total_human, total_machine, total_nucleation,
		total_skrell, total_tajara, total_unathi, total_vox, total_wryn,
	}

	return &Stats{
		total_acc_items, total_bans, avg_bans, total_rounds, total_deaths,
		avg_deaths, total_round_duration, avg_round_duration,
		total_monkey_deaths, total_damage_cost, total_players,
		total_characters, species,
	}
}

func (db *DB) SearchBans(ckey string) []*Ban {
	var tmp []*Ban
	// Don't want any weird behaviours if the user is smart enough to try to use these
	tckey := "%" + strings.Trim(strings.TrimSpace(ckey), "_%") + "%"
	db.Order("id desc, ckey asc").Where("ckey LIKE ?", tckey).Limit(MAX_ROWS).Find(&tmp)
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
	db.Order("round_id desc, law asc").Group("law").Where(
		"round_id = ?", id).Find(&tmp)
	return tmp
}

func (db *DB) GetDeaths(id int64) []*Death {
	var tmp []*Death
	db.Order("round_id desc, name asc").Group("name").Where(
		"round_id = ?", id).Find(&tmp)
	return tmp
}

func (db *DB) GetCharacter(id int64) *Character {
	var tmp Character
	db.First(&tmp, id)
	return &tmp
}

func (db *DB) SearchCharacter(species, name string) []*Character {
	var tmp []*Character
	species = strings.ToLower(species)
	switch species {
	case "diona", "human", "machine", "nucleation", "skrell", "tajara",
		"unathi", "vox", "wryn":
		// Allow these values only and do nothing more with them
	default:
		species = ""
	}
	tname := "%" + strings.Trim(strings.TrimSpace(name), "_%") + "%"

	// The NOT should filter out weird characters with no names set
	if len(species) > 0 {
		db.Order("name asc").Where("species = ? AND name LIKE ?",
			species, tname).Not("name = ''").Limit(MAX_ROWS).Find(&tmp)
	} else {
		db.Order("name asc").Where("name LIKE ?", tname).Not(
			"name = ''").Limit(MAX_ROWS).Find(&tmp)
	}
	return tmp
}

func (db *DB) AllGameModes() []*GameMode {
	// I feel dirty for having made this...
	var game_modes []string
	var modes GameModeSlice
	var total_rounds int64
	db.Table("round_stats").Pluck("DISTINCT(game_mode)", &game_modes)
	for _, m := range game_modes {
		var g GameMode
		db.Table("round_stats").Where("game_mode=?", m).Select(`
			COUNT(id) as total_rounds,
			AVG(duration) as avg_duration,
			AVG(productivity) as avg_productivity,
			AVG(deaths) as avg_deaths
		`).Find(&g)

		g.Title = strings.Title(m)
		total_rounds += g.TotalRounds
		modes = append(modes, &g)
	}

	var t float64
	for _, m := range modes {
		m.AvgRounds = (float64(m.TotalRounds) / float64(total_rounds)) * 100
		t += m.AvgRounds
	}
	sort.Sort(modes)
	return modes
}
