package apollostats

import "time"

type Cache struct {
	LatestRound *RoundStats
	GameStats   *Stats
	GameModes   []*GameMode
	LastUpdated time.Time
	UpdateTime  time.Duration

	in        *Instance
	closeChan chan bool
}

func NewCache(i *Instance) *Cache {
	return &Cache{
		in:        i,
		closeChan: make(chan bool),
	}
}

func (c *Cache) close() {
	c.closeChan <- true
}

func (c *Cache) updater() {
	d := time.Duration(CACHE_UPDATE * time.Minute)
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	c.in.logMsg("Running cache updates every %s.", d.String())

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
	c.LatestRound = c.in.DB.GetLatestRound()
	c.GameStats = c.in.DB.GetStats()
	c.GameModes = c.in.DB.AllGameModes()
	c.UpdateTime = time.Since(c.LastUpdated)
	c.in.logMsg("Updated cache at %s (%s)", c.LastUpdated.String(), c.UpdateTime.String())
}
