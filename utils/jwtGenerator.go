package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/satori/go.uuid"
	"go_test_project/models"
	"time"
)

var hs *jwt.ECDSASHA //每次服务器启动时重新生成

func init() {
	var privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	var publicKey = &privateKey.PublicKey
	hs = jwt.NewES256(
		jwt.ECDSAPublicKey(publicKey),
		jwt.ECDSAPrivateKey(privateKey),
	)
}

// Sign 签名
func Sign(id string, username string) (string, error) {
	now := time.Now()
	pl := models.LoginToken{
		Payload: jwt.Payload{
			Issuer:         "utf8coding",
			Audience:       jwt.Audience{},
			ExpirationTime: jwt.NumericDate(now.Add(7 * 24 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.NewV4().String(),
		},
		ID:       id,
		Username: username,
	}
	token, err := jwt.Sign(pl, hs)
	return string(token), err
}

// Verify 验证
func Verify(token []byte) (*models.LoginToken, error) {
	pl := &models.LoginToken{}
	_, err := jwt.Verify(token, hs, pl)
	return pl, err
}
