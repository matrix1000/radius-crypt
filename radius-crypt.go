/*
 * GPLv3
 * Crypt for Radius
 * https://freeradius.org/radiusd/man/rlm_pap.html
 *
 */

package main

import (
    "crypto/md5"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/base64"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

//VERSION of this script
const VERSION string = "1.0"

func getSalt(length int) string {
    rand.Seed(time.Now().UnixNano())
    chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
        "abcdefghijklmnopqrstuvwxyz" +
        "0123456789" +
        "-_")
    var b strings.Builder
    for i := 0; i < length; i++ {
        b.WriteRune(chars[rand.Intn(len(chars))])
    }
    return string(b.String())
}

func main() {
    salt := getSalt(20)
    if len(os.Args) != 2 {
        fmt.Println("Usage:", os.Args[0], "PASSWORD")
        fmt.Println("Please insert only one password.")
        os.Exit(1)
    }
    pw := os.Args[1]
    md5 := md5.New()
    sha256 := sha256.New()
    sha512 := sha512.New()
    pw = pw + salt
    md5.Write([]byte(string(pw)))
    sha256.Write([]byte(string(pw)))
    sha512.Write([]byte(string(pw)))
    rMd5 := md5.Sum(nil)
    rSha256 := sha256.Sum(nil)
    rSha512 := sha512.Sum(nil)
    fmt.Println("Version:", VERSION)
    fmt.Println("Salt:", salt)
    fmt.Println("{SMD5-Password}:", base64.StdEncoding.EncodeToString(append(rMd5, []byte(salt)...)))
    fmt.Println("{SSHA2-256-Password}:", base64.StdEncoding.EncodeToString(append(rSha256, []byte(salt)...)))
    fmt.Println("{SSHA2-512-Password}:", base64.StdEncoding.EncodeToString(append(rSha512, []byte(salt)...)))
}
