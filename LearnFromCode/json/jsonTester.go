package main

import (
	"fmt"
	"encoding/json"
	"flag"
	"goimport/io/ioutil"
)
type Book struct{
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float64
}

func exerciseExample(){
	gobook := &Book{
		"The Go Programming Language",
		[]string{"Huangyuyan", "Lijiaqi", "Linkeying", "Guocaishan"},
		"ituring.com.cn",
		true,
		99.9,
	}
	b,err := json.Marshal(gobook)
	if err != nil {
		return
	}
	var r interface{}
	err = json.Unmarshal(b,&r)
	if err != nil{
		return
	}
	t,ok :=r.(map[string]interface{})
	if ok{
		for k,v :=range t {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i,iv :=range v2 {
					fmt.Println(i,iv)
				}
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}

}
type UrlList struct {
	Url []string `json:"urlList"`
}
func getUrl(filePath string) ([]string, error){
	var (
		urlList UrlList
	)
	js, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("--- getUrl read json file fail, err:%v\n", err)
		return nil, err
	} else {
		err = json.Unmarshal(js, &urlList)
		if err != nil {
			fmt.Printf("--- getUrl unmarshal %v fail, err:%v\n", filePath, err)
			return nil, err
		}
	}
	fmt.Printf("====== looping url list ======\n")
	var cnt int
	for _, url := range urlList.Url {
		fmt.Printf("+++ url%d: %s\n", cnt, url)
		cnt++
	}
	fmt.Printf("====== looping url list end ======\n")
	return urlList.Url, nil
}

func main(){
	cmd := flag.String("cmd", "h", "For more Command.")
	filePath := flag.String("filePath", "./urlList.json", "Url list json file path.")
	flag.Parse()

	switch *cmd {
	case "exExample":
		exerciseExample()
		break
	case "getUrl":
		getUrl(*filePath)
		break
	default:
		fmt.Printf("Usage: ./jsonTester -cmd=\"exExample\"\n")
		fmt.Printf("Usage: ./jsonTester -cmd=\"getUrl\" -filePath=\"./urlList\"\n")
		return

	}
}