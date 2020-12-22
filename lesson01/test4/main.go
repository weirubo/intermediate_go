package main

import (
	"encoding/json"
	"fmt"
)

// Opus 作品
type Opus struct {
	Type  string
	Title string
}

// Actress 女演员
type Actress struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       map[string]Opus
}

func main() {
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":{
         "2013":{
            "Type":"近代革命剧",
            "Title":"《阿娜尔罕》"
         },
         "2014":{
            "Type":"奇幻剧",
            "Title":"《逆光之恋》"
         },
         "2015":{
            "Type":"爱情剧",
            "Title":"《克拉恋人》"
         }
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
	for index, value := range actress.Opus {
		fmt.Printf("\t日期：%s\n", index)
		fmt.Printf("\t\t分类：%s\n", value.Type)
		fmt.Printf("\t\t标题：%s\n", value.Title)
	}
}
