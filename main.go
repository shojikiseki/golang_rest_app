package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// 入力を定義
type postHelloInput struct {
	Name string
}

// 出力を定義
type postHelloOutput struct {
	Result string
}

func postHello(w rest.ResponseWriter, req *rest.Request) {
	input := postHelloInput{}
	err := req.DecodeJsonPayload(&input)

	// 入力形式が違ったらエラー
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Nameが空だったらエラー
	if input.Name == "" {
		rest.Error(w, "Name is required", 400)
		return
	}

	log.Printf("%#v", input)

	// 結果を返す部分
	w.WriteJson(&postHelloOutput{
		"hello, " + input.Name,
	})
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/hello", postHello),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server started...")
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9999", api.MakeHandler()))
}
