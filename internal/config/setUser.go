package config
//alters config.json by passing in struct data with new username, see write.go
func (c *Config) SetUser(s string) {
	c.CurrentUsername = s
	write(c)
}
