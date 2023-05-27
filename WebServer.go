package main

import (
	"fmt"
	//フォーマットI/Oを扱うパッケージで、PrintlnやFprintのような関数,HTTPレスポンスの生成やサーバーの起動情報の出力を提供
	"net/http"
	//HTTPクライアントとサーバーの実装を提供するパッケージ
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
	//w（レスポンスライター）に文字列を書き込む
	//r:リクエストの情報を保持する構造体
}

func main() {
	http.HandleFunc("/", helloHandler)
	//指定されたパスへのHTTPリクエストを処理

	fmt.Println("Starting server at port 8080")
	//コンソールにメッセージ表示

	if err := http.ListenAndServe(":8080", nil); err != nil {
		//:= 演算子はGoで変数の宣言と同時に初期化を行うためのもの
		//nilは基本的に他言語のnullと同じ。nilは型を持つという違いがある
		fmt.Println("Failed to start server: ", err)
	}
}
