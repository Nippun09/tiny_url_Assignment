package db

var databaseHashcodeToShort map[interface{}]string //we cannot use []byte as key, so using interface instead
var databaseShortToLong map[string]string

func init() {
	databaseHashcodeToShort = make(map[interface{}]string)
	databaseShortToLong = make(map[string]string)
}

func MapLongUrlHashcodeToShort(hascode []byte, shortUrl string) {

	databaseHashcodeToShort[hascode] = shortUrl

}

func GetHashcodeToShortDb() map[interface{}]string {
	return databaseHashcodeToShort
}

func MapShortUrlToLong(shortUrl, longUrl string) {
	databaseShortToLong[shortUrl] = longUrl
}

func GetShortUrlToLongDb() map[string]string {
	return databaseShortToLong
}
