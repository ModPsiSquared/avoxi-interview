# avoxi-interview

## Make file instructions
Use the makefile included for common operations

### make build
    builds a local copy in the dist folder

### make run
    runs the service locally defaults to 10,000

### make dist/avoxiinterview
    builds a fresh docker image with contents of the dist folder

### make docker-run
    Runs the container created above.

### make tests
    Runs the tests.

## Concurrency
Can the library (github.com/oschwald/geoip2-golang) support multiple opens or reads of
the geolocation database?

Could not find anything really in the documentation to indicate a concurrency problem in either
proved it to myself with dirty test.  I did not check into the tests:

```
func TestDatabasePull_concurrent(t *testing.T) {
	ips := []net.IP{net.ParseIP("82.55.213.9"), net.ParseIP("2.16.108.0"), net.ParseIP("2.57.172.0")}
	countrys := []string{"Italy", "Spain", "luxembourg"}

	for i := 1; i < 200; i++ {
		go func(l int){
			ip := ips [ l % 2 ]
			country := countrys [l % 2]

			db, err := geoip2.Open("./data/GeoLite2-Country_20210413/GeoLite2-Country.mmdb")
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
		}(i)
	}
}
```

### Solution for GeoLite2 database update

1) We need to persist a record of the path to the database.  Changes to path of live db cannot while process of checkIP.
2) Call to a new endpoint called "updatedb" could be used to start the process of pulling back a database and 
   unzipping it next to the file of the current database.
3) Once file copy completed a mutex will need to be used to prevent attempted database reads while new path to new database is set (persisted)
4) Finally, we can clean up the now outdated database by simple file deletion.  



