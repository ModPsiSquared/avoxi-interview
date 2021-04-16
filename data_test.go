package main_test

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"testing"
)

func TestDatabasePull(t *testing.T) {
	ip := net.ParseIP("82.55.213.9")
	country := "Italy"

	db, err := geoip2.Open("./data/GeoLite2-Country_20210413/GeoLite2-Country.mmdb")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	record, err := db.Country(ip)
	if err != nil {
		log.Print(err)
	}
	for _, n := range record.Country.Names {
		if n == country {
			return
		}
	}
	t.Fail()
}

func BenchmarkDatabasePull(b *testing.B) {
	ips := []net.IP{net.ParseIP("82.55.213.9"), net.ParseIP("2.16.108.0"), net.ParseIP("2.57.172.0")}
	countrys := []string{"Italy", "Spain", "luxembourg"}

	for i := 0; i < b.N; i++ {
		ip := ips[i%2]
		country := countrys[i%2]

		db, err := geoip2.Open("./data/GeoLite2-Country_20210413/GeoLite2-Countries.mmdb")
		if err != nil {
			log.Panic(err)
		}
		defer db.Close()

		record, err := db.Country(ip)
		if err != nil {
			log.Panic(err)
		}
		for _, n := range record.Country.Names {
			if n == country {
				return
			}
		}
		panic("failed")
	}
}
