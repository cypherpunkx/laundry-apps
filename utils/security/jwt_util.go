package security

import (
	"fmt"
	"time"

	"enigmacamp.com/enigma-laundry-apps/config"
	"enigmacamp.com/enigma-laundry-apps/model"
	"github.com/golang-jwt/jwt/v5"
)

// generate token
// verifikasi token

func CreateAccessToken(user model.UserCredential) (string,error) {
	cfg, _ := config.NewConfig()

	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: cfg.ApplicationName,
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod,claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		return "",fmt.Errorf("Failed to create access token : %s", err.Error())
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims,error) {
	fmt.Println("TOken", tokenString)
	cfg, _ := config.NewConfig()
	// parse token yang dikirim dari client 
	token, err := jwt.Parse(tokenString,func(t *jwt.Token)(interface{},error){
		// check method yang digunakan
		// validasi signing method yaitu yang kita gunakan HS256
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != cfg.JwtSigningMethod {
			return nil,fmt.Errorf("Invalid token signing method")
		}
		return cfg.JwtSignatureKey,nil
	})

	if err != nil {
		return nil,fmt.Errorf("Invalid parse token sdf : %s",err.Error())
	}
	// cek claims yang sudah didaftarkan sebelumnya
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil,fmt.Errorf("Invalid token MapClaims")
	}
	return claims,nil
}