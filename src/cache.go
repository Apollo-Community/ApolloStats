package apollostats

import (
	"fmt"
	"time"
)

type Cache struct {
	DB          *DB
	LatestRound *RoundStats
	GameStats   *Stats
	GameModes   []*GameMode
	LastUpdated time.Time
	UpdateTime  time.Duration

	closeChan chan bool
}

func NewCache(db *DB) *Cache {
	return &Cache{
		DB:        db,
		closeChan: make(chan bool),
	}
}

func (c *Cache) close() {
	c.closeChan <- true
}

func (c *Cache) updater() {
	ticker := time.NewTicker(CACHE_UPDATE * time.Minute)
	defer ticker.Stop()

	c.updateCache()
	for {
		select {
		case <-c.closeChan:
			return
		case <-ticker.C:
			c.updateCache()
		}
	}
}

func (c *Cache) updateCache() {
	c.LastUpdated = time.Now()
	c.LatestRound = c.DB.GetLatestRound()
	c.GameStats = c.DB.GetStats()
	c.GameModes = c.DB.AllGameModes()
	c.UpdateTime = time.Since(c.LastUpdated)
	fmt.Println("Cache update took", c.UpdateTime)
}
