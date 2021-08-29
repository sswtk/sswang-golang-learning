package main

import (
	"fmt"
	"sswang-golang-learning/go_learning_imook/retriever/mock"
	"sswang-golang-learning/go_learning_imook/retriever/real"
	"time"
)

const url ="http://www.baidu.com"


type Retriever interface {
	Get(url string) string
}

type Poster interface{
	Post(url string, form map[string]string) string
}


func download(r Retriever) string {
	return r.Get(url)
}


func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "ccmouse",
			"course": "golang",
		})
}


type RetrieverPoster interface {
	Retriever
	Poster
}


func session(s RetrieverPoster) string{
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}


func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type){
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}


func main(){
	var r Retriever
	retriever  := mock.Retriever{"This is a fake baidu.com"}
	//fmt.Printf("%T %v\n", r, r)
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	//Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	//fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}
