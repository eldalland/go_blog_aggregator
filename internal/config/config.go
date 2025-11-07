package config
//struct that mirrors structure of config.json file
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}
