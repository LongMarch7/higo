package token

import (
    "crypto/md5"
    "fmt"
    "io"
    "strconv"
    "time"
)
func NewToken(user string) string{
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    io.WriteString(h, user)
    token := fmt.Sprintf("%x", h.Sum(nil))
    return token
}
