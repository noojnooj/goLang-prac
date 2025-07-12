// main 패키지 지정
package main

// fmt는 Go의 표준 출력 라이브러리
import (
	"fmt"
	"net/http"
	"os"
	// godotenv는 .env 파일에서 환경 변수를 로드하는 라이브러리
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// HTTP 요청에 대한 응답을 작성
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}

// 실행이 시작되는 함수
// main 함수는 프로그램 내에서 하나만 존재해야함.
func main() {
	godotenv.Load()           // .env 파일에서 환경 변수 로드
	port := os.Getenv("PORT") // 환경 변수에서 PORT 값을 가져옴

	http.HandleFunc("/", handler) // 루트 경로에 핸들러를 등록

	fmt.Print("서버 실행 콘솔, http://localhost:", port)
	http.ListenAndServe(":"+port, nil)
}

//
