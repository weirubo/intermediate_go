package main

import (
	"encoding/json"
	"fmt"
)

// Opus 作品
type Opus struct {
	Date  string
	Title string
}

// Actress 女演员
type Actress struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       Opus
}

func main() {
	// JSON嵌套普通JSON
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus": {
         "Date":"2013",
         "Title":"《阿娜尔罕》"
      }
   }`)
	var actress Actress
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	fmt.Printf("\t%s:%s", actress.Opus.Date,

		actress.Opus.Title)

}
