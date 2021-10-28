package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (

	urls = "https://yz.chsi.com.cn/zyk/specialityCategory.do"
)

func main() {

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodPost,urls, nil )
	req.Form.Set("method", "subCategoryXk")
	req.Form.Set("key","100812")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(content)
}
