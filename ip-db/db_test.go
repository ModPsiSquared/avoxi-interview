package ip_db

import (
	"net"
	"reflect"
	"testing"
)

func TestGetCountryForIp(t *testing.T) {
	type args struct {
		ip net.IP
	}

	err := InitDb()
	if err != nil {
		t.FailNow()
	}

	tests := []struct {
		name        string
		args        args
		wantCountry string
		wantErr     bool
	}{
		{
			"Is in list",
			args{ip: net.ParseIP("2.16.146.0")},
			"Italy",
			false,
		},
		{
			"Is in list",
			args{ip: net.ParseIP("18.184.170.128")},
			"Germany",
			false,
		},
		{
			"Is in list",
			args{ip: net.ParseIP("154.73.64.0")},
			"Tanzania",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCountries, err := GetENCountryForIp(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetENCountryForIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCountries, tt.wantCountry) {
				t.Errorf("GetENCountryForIp() gotCountry = %v, want %v", gotCountries, tt.wantCountry)
			}
		})
	}
}
