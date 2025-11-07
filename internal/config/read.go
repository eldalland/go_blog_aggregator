package config

import (
	"encoding/json"
	"os"
)
//function reads json data from config, unmarshals into readable data as a struct
func Read()(Config,error){
	configUrl,err := getConfigFilePath("/.gatorconfig.json")
	if err != nil{
		return Config{}, err
	}
	json_config_data, err := os.ReadFile(configUrl)
	if err != nil{
		return Config{}, err
	}
	config_data_struct := Config{}
	json.Unmarshal(json_config_data,&config_data_struct)
	return config_data_struct, nil
}