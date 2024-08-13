package db

func CreateConnection() *Config {
	server := getEnv("database.server")
	port := getEnv("database.port")
	user := getEnv("database.user")
	pass := GetDecrypted("database.password")
	scheme := getEnv("database.scheme")
	dsn := buildPostgresDSN(server, port, user, pass, scheme)

	return &Config{
		DataSourceName:    dsn,
		MaxOpenCons:       5,
		MaxIdleCons:       5,
		ConnMaxIdleTime:   0,
		UseConnectionPool: false}
}
