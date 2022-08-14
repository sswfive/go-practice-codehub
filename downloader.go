package main

import (
	"fmt"
	"go-practice-codehub/infra"
)

//func main() {
//	resp, err := http.Get("https://www.imooc.com")
//	if err != nil {
//		panic(err)
//	}
//
//	defer resp.Body.Close()
//
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	fmt.Printf("%s\n ", bytes)
//}

// 改进版
//func retrieve(url string) string {
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	return string(bytes)
//}

func getRetriever() retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	//retriever := getRetriever()
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
