package db

import (
	"fmt"
	"time"

	"apigolang/src/apigo/utils"

	"github.com/spf13/viper"
)

type Config struct {
	DataSourceName    string
	MaxOpenCons       int
	MaxIdleCons       int
	ConnMaxIdleTime   time.Duration
	UseConnectionPool bool
}

func getEnv(key string) string {
	return viper.GetString(key)
}

func GetDecrypted(key string) string {
	encryptedPass := getEnv(key)
	pass := utils.DecryptAES(encryptedPass)
	if pass == "" {
		return ""
	}

	return pass
}

func buildPostgresDSN(server, port, user, password, dbname string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		server, port, user, password, dbname)
}
