package api

import (
	ip_db "avoxi-interview/ip-db"
	"avoxi-interview/models"
	"avoxi-interview/utility"
	"log"
	"net/http"
	"strings"
)

func IpCountryCheck(w http.ResponseWriter, r *http.Request) {
	var p models.Payload
	utility.PayloadDecoder(w, r, &p, func() () {
		englishCountryName, err := ip_db.GetENCountryForIp(p.Ip)
		if err != nil {
			log.Printf("Error in getting countries from ip database. %s", p.Ip.String())
		}
		var approved = false
		for _, m := range p.Countries {
			if strings.EqualFold(englishCountryName, m) {
				approved = true
				break
			}
		}
		status := models.NewIpStatus(p.Ip, approved)
		utility.Respond(w, 200, status)
	})
}
