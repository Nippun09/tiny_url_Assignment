package db

var databaseHashcodeToShort map[uint32]string //we cannot use []byte as key, so using interface instead
var databaseShortToLong map[string]string

func init() {
	databaseHashcodeToShort = make(map[uint32]string)
	databaseShortToLong = make(map[string]string)
}

func MapLongUrlHashcodeToShort(hascode uint32, shortUrl string) {

	databaseHashcodeToShort[hascode] = shortUrl

}

func GetHashcodeToShortDb() map[uint32]string {
	return databaseHashcodeToShort
}

func MapShortUrlToLong(shortUrl, longUrl string) {
	databaseShortToLong[shortUrl] = longUrl
}

func GetShortUrlToLongDb() map[string]string {
	return databaseShortToLong
}
