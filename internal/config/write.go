package config
//alters config.json by passing in struct data with new username
import (
	"encoding/json"
	"os"
)
func write(c *Config)error{
	configUrl,err := getConfigFilePath("/.gatorconfig.json")
	if err != nil{
		return err
	}
	data_to_write, err := json.Marshal(c)
	if err !=nil{
		return err
	}
	err = os.WriteFile(configUrl,data_to_write,0660)
	if err != nil{
		return err
	}
	return nil
}