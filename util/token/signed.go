package token

import (
    "encoding/base64"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/lifenod/secure"
)

var (
    // Generate by RandomBase64StdString(512 / 8).
    hs512SigningKey = []byte("aWYlMjB5b3UlMjBpbnZhZGUlMjBteSUyMHN5c3RlbSV1RkYwQ2klMjB3aWxsJTIwZnVjayUyMHlvdXIlMjBtb3RoZXI=")
    ExpireTime = 12
)

type DefaultClaims struct {
    Pwd      string `json:"pwd"`
}

func (c DefaultClaims) Valid() error {
    if len(c.Pwd) == 0 {
        return fmt.Errorf("pwd is nil")
    }


    return nil
}

func NewDefaultClaims() jwt.Claims{
    customClaims := &DefaultClaims{
        Pwd: "default",
    }
    return customClaims
}

// * Return errors:
// * token.SignedString
func Sign(standardClaims jwt.Claims) (signedString string, err error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardClaims)
    signedString, err = token.SignedString(hs512SigningKey)
    if err != nil {
        // TODO define a special error.
        return "", err
    }
    signedString = base64.StdEncoding.EncodeToString([]byte(signedString))
    return signedString, nil
}

// * Return errors:
// * secure.ErrTokenInvalid
func Parse(signedString string,cl jwt.Claims) (interface{}, error) {
    signedByte, _ := base64.StdEncoding.DecodeString(signedString)
    signedString = string(signedByte)
    token, err := jwt.ParseWithClaims(signedString, cl, func(innerToken *jwt.Token) (interface{}, error) {
        if _, ok := innerToken.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, secure.ErrTokenInvalid
        }
        return hs512SigningKey, nil
    })
    if err != nil {
        return nil, secure.ErrTokenInvalid
    }

    if !token.Valid {
        return nil, secure.ErrTokenInvalid
    }

    return token.Claims, nil
}

