package jwt

import "net/http"

type JwtCalimHandler func(res http.ResponseWriter, req *http.Request, claim *JwtUserCalim)

func JwtVerifyMiddleware(next JwtCalimHandler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		bearer := req.Header.Values("Authorization")

		// bearer의 길이가 8글자가 넘는지 확인
		if len(bearer) < 1 || len(bearer[0]) < 8 {
			http.Error(res, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		// Bearer 에서 jwt 토큰을 추출
		jwtToken := bearer[0][7:]

		claim, err := ValidateJwtToken(jwtToken)
		if err != nil {
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		}

		next(res, req, claim)
	})
}
