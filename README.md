### 오류

.env 생성 후 자바스크립트처럼 환경변수 불러오려했는데 안됨.

```bash
Hello.go:10:2: no required module provides package github.com/joho/godotenv: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

.env 파일 활용하고 싶다면 godoenv 패키지를 사용해서 파일을 로드해야한다고 한다.

```go
<!-- 추가 -->
import "github.com/joho/godotenv"
```

```bash
<!-- 외부 패키지 설치(godoenv) -->
go get github.com/joho/godotenv
```

```bash
go run Hello.go
```
