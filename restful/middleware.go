// package main
// // ttt
// import (
// 	"log"
// 	"time"

// 	echojwt "github.com/labstack/echo-jwt/v4"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func main() {
// 	e := echo.New()

// 	// 1. Logger : 모든 HTTP 요청과 응답에 대한 로그 기록
// 	e.Use(middleware.Logger())

// 	// 2. Recover : 패닉복구 서버에서 발생하는 패닉을 복구하고, 서버가 중단되지 않도록 합니다. 패닉이 발생하면 스택 트레이스를 로그에 기록하고, 클라이언트에게 HTTP 500 상태 코드를 반환합니다.
// 	e.Use(middleware.Recover())

// 	// 3. CORS : Cross-Origin Resource Sharing (CORS) 다른 도메인에서 리소스에 접근할 수 있도록 서버의 응답에 헤더를 추가합니다.
// 	e.Use(middleware.CORS())

// 	// 4. JWT : JSON Web Token을 사용한 인증을 지원합니다. 이 미들웨어를 통해 JWT 토큰을 검증하여 보호된 엔드포인트에 접근할 수 있습니다.
// 	e.Use(echojwt.JWT([]byte("secret")))

// 	// 5. BasicAuth : HTTP Basic 인증을 지원합니다. 이 미들웨어를 통해 사용자 이름과 비밀번호를 검증하여 보호된 엔드포인트에 접근할 수 있습니다.
// 	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
// 		if username == "jhpark" && password == "1234" {
// 			return true, nil
// 		}
// 		return false, nil
// 	}))

// 	// 6. Secure : 보안 관련 헤더를 추가합니다. 이 미들웨어를 사용하면 X-XSS-Protection, X-Content-Type-Options, X-Frame-Options, Content-Security-Policy, Strict-Transport-Security 헤더를 추가할 수 있습니다.
// 	e.Use(middleware.Secure())

// 	// 7. BodyLimit : 요청 본문의 크기를 제한합니다. 이 미들웨어를 사용하면 요청 본문의 최대 크기를 제한할 수 있습니다., 대용량 요청으로 인해 과부하가 걸리지 않도록 할 수 있습니다.
// 	e.Use(middleware.BodyLimit("2M"))

// 	// 8. Gzip : Gzip 압축을 지원합니다. 이 미들웨어를 사용하면 클라이언트가 Gzip 압축을 지원하는 경우 응답을 압축하여 전송할 수 있습니다.
// 	e.Use(middleware.Gzip())

// 	// 9. RequestID : 요청에 대한 고유한 ID를 생성합니다. 이 미들웨어를 사용하면 요청에 대한 고유한 ID를 생성하여 로그에 기록할 수 있습니다.
// 	e.Use(middleware.RequestID())

// 	// 10. Timeout : 요청에 대한 타임아웃을 설정합니다. 이 미들웨어를 사용하면 요청에 대한 타임아웃을 설정하여 요청이 지정된 시간 내에 완료되지 않으면 요청을 취소할 수 있습니다.
// 	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
// 		Skipper:      middleware.DefaultSkipper,
// 		ErrorMessage: "custom timeout error message returns to client",
// 		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
// 			log.Println(c.Path())
// 		},
// 		Timeout: 30 * time.Second, // 30초 설정
// 	}))

// 	// // 라우트 설정
// 	// e.GET("/", func(c echo.Context) error {
// 	// 	return c.String(http.StatusOK, "Hello, World!")
// 	// })

// 	// // 서버 실행
// 	// e.Start(":8080")
// }
