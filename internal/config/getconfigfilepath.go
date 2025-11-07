package config
import "os"
//method to quickly generate path to config.json
func getConfigFilePath(s string)(string,error){
	homeURL,err := os.UserHomeDir()
	if err != nil{
		return "", err
	}
	configUrl := homeURL + s
	return configUrl, nil
}