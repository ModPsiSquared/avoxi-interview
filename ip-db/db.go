package ip_db

import (
	"avoxi-interview/utility"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func InitDb() (err error) {
	Ip_Db, err := geoip2.Open(utility.GetDB())
	if err != nil {
		return
	}
	Ip_Db.Close()
	return
}

func GetENCountryForIp(ip net.IP) (country string, err error) {
	ip_db, err := geoip2.Open(utility.GetDB())
	if err != nil {
		return
	}
	defer ip_db.Close()

	record, err := ip_db.Country(ip)
	if err != nil {
		return
	}
	country = record.Country.Names["en"]
	return
}
