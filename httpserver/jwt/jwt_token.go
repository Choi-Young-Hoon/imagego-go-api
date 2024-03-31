package jwt

import "github.com/dgrijalva/jwt-go"

var jwtSigningKey = []byte("asdnadsjkfnsfnaesjflnaejlsnflsj@#$%%!@%@")

type JwtUserCalim struct {
	UserId    string `json:"UserId"`
	IpAddress string `json:"IpAddress"`

	jwt.StandardClaims
}

func GenerateJwtToken(usersId string, ipAddress string) (string, error) {
	claims := JwtUserCalim{
		usersId,
		ipAddress,
		jwt.StandardClaims{
			ExpiresAt: 0,
			Issuer:    "imagego",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSigningKey)
}

func ValidateJwtToken(tokenString string) (*JwtUserCalim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtUserCalim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtUserCalim); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
