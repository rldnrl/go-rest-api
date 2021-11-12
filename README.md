# Go로 만든 Rest API
## Stack: Fiber, Gorm

---

## 개발하면서 막혔던 문제
1. Fiber 설치하면서 발생했던 문제<br />
Fiber를 설치를 했는데, 모듈을 불러올 수 없는 문제에 부딪혔다. 먼저 Go 프로젝트에 mod를 설정해야한다. 명령어는 아래와 같다.

```bash
go mod init github.com/<github_repo_name>
```

그리고 다시 fiber와 gorm을 설치하고 불러오면 잘 작동하게 된다.

2. GORM과 DB를 연결했는데, DB에 있는 내용을 JSON 형태로 응답하지 않는 문제<br />

문제가 발생한 코드
```go
const DNS = "root:root124@tcp(127.0.0.1:3306)/godb"
```

수정한 코드
```go
const DNS = "root:root124@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
```