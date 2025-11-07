package rssfeed
import ("fmt"
"io"
"html"
"encoding/xml"
"net/http"
"context")
//gets RSSFeed data from http request, then unmarshals returned xml data into struct
func FetchFeed(ctx context.Context, feedURL string)(*RSSFeed, error){
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx,"GET",feedURL,nil)
	if err != nil{
		return &RSSFeed{}, fmt.Errorf("error creating request: %s",err)
	}
	req.Header.Set("User-Agent","gator")
	res, err := client.Do(req)
	if err != nil{
		return &RSSFeed{}, fmt.Errorf("error making request: %s",err)
	}
	defer res.Body.Close()
	//read all response data
	body, err := io.ReadAll(res.Body,)
	if err != nil{
		return &RSSFeed{}, fmt.Errorf("error reading res.body: %s",err) 
	}
	xmlData := RSSFeed{}
	xml.Unmarshal(body,&xmlData)
	xmlData.Channel.Title = html.UnescapeString(xmlData.Channel.Title)
	xmlData.Channel.Description = html.UnescapeString(xmlData.Channel.Description)
	for i := range xmlData.Channel.Item{
		xmlData.Channel.Item[i].Title = html.UnescapeString(xmlData.Channel.Item[i].Title)
		xmlData.Channel.Item[i].Description = html.UnescapeString(xmlData.Channel.Item[i].Description)
	}
	return &xmlData, nil
	}