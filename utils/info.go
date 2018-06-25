package utils

type PostgresInfo struct {
	User string
	Pass string
}

// GetInfo holds and returns credentials for postgres db
func GetInfo() *PostgresInfo {
	var userInfo PostgresInfo

	userInfo.User = "userID"
	userInfo.Pass = "password"
	return &userInfo
}
