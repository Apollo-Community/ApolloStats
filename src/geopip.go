package apollostats

import (
	"net"
	"sort"

	"github.com/oschwald/maxminddb-golang"
)

type mmCountry struct {
	Continent struct {
		Code  string `maxminddb:"code"`
		Names struct {
			En string `maxminddb:"en"`
		} `maxminddb:"names"`
	} `maxminddb:"continent"`

	Country struct {
		ISOCode string `maxminddb:"iso_code"`
		Names   struct {
			En string `maxminddb:"en"`
		} `maxminddb:"names"`
	} `maxminddb:"country"`
}

func GeoLookup(players []*Player) ([]*Country, error) {
	db, e := maxminddb.Open("GeoLite2-Country.mmdb")
	if e != nil {
		return nil, e
	}
	defer db.Close()

	m := make(map[string]*Country)
	for _, p := range players {
		ip := net.ParseIP(p.IP)
		var c mmCountry
		e = db.Lookup(ip, &c)
		if e != nil {
			continue // skip this player/ip then
		}

		if _, ok := m[c.Country.ISOCode]; ok == true {
			m[c.Country.ISOCode].Hits += 1
		} else {
			m[c.Country.ISOCode] = &Country{
				ISOCode:   c.Country.ISOCode,
				Name:      c.Country.Names.En,
				Continent: c.Continent.Names.En,
				Hits:      1,
			}
		}
	}

	// Make a slice so we can sort it and return the top 10 countries.
	s := make(countrySlice, 0, len(m))
	for _, c := range m {
		s = append(s, c)
	}
	sort.Sort(s)
	max := 10
	if len(s) < max {
		max = len(s)
	}
	return s[:max], nil
}
