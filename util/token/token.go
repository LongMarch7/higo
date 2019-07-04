package token

import (
    "crypto/md5"
    "fmt"
    "io"
    "strconv"
    "time"
)

func NewTokenWithTime(user string) string{
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    io.WriteString(h, user)
    token := fmt.Sprintf("%x", h.Sum(nil))
    return token
}

func NewTokenWithSalt(value string) string{
    return NewTokenWithSalt2(NewTokenWithSalt1(value))
}

func NewTokenWithSalt1(value string) string{
    h := md5.New()
    salt := "91ea4fe9267bcfb42bf0c27211507390"
    io.WriteString(h, value)
    io.WriteString(h, salt)
    token := fmt.Sprintf("%x", h.Sum(nil))
    return token
}

func NewTokenWithSalt2(value string) string{
    h := md5.New()
    salt := "ac23312dd195185d73f469562f50f892"
    io.WriteString(h, value)
    io.WriteString(h, salt)
    token := fmt.Sprintf("%x", h.Sum(nil))
    return token
}
