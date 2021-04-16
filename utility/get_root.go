package utility

import (
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func AddBaseDirectory(s string) (ret string) {
	r := RootDir()
	return path.Join(r, s)
}

func GetDB() string {
	return AddBaseDirectory("data/GeoLite2-Country_20210413/GeoLite2-Country.mmdb")
}

func GetConfigFolder() string {
	return AddBaseDirectory("data")
}
