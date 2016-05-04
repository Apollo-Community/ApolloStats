package apollostats

const VERSION = "0.4"

// Max rows DB will return for all queries.
const MAX_ROWS = 200

// DB connection timeout, in seconds.
const TIMEOUT = 30

// Need to adjust time because the main server is running GMT-5.
const TIMEZONE_ADJUST = "EST"

// Cache update every x minutes.
const CACHE_UPDATE = 60
