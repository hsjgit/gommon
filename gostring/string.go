// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gostring

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/hsjgit/gommon/b64"
)

func JoinStrings(multiString ...string) string {
	return strings.Join(multiString, "")
}

func JoinIntSlice2String(intSlice []int, sep string) string {
	return strings.Join(IntSlice2StrSlice(intSlice), sep)
}

func StrSplit2IntSlice(str, sep string) []int {
	return StrSlice2IntSlice(StrFilterSliceEmpty(strings.Split(str, sep)))
}

func Str2StrSlice(str, sep string) []string {
	return StrFilterSliceEmpty(strings.Split(str, sep))
}

func StrSlice2IntSlice(strSlice []string) []int {
	var intSlice []int
	for _, s := range strSlice {
		i, _ := strconv.Atoi(s)
		intSlice = append(intSlice, i)
	}
	return intSlice
}

func StrFilterSliceEmpty(strSlice []string) []string {
	var filterSlice []string
	for _, s := range strSlice {
		ss := strings.TrimSpace(s)
		if ss != "" {
			filterSlice = append(filterSlice, ss)
		}
	}
	return filterSlice
}

func IntSlice2StrSlice(intSlice []int) []string {
	var strSlice []string
	for _, i := range intSlice {
		s := strconv.Itoa(i)
		strSlice = append(strSlice, s)
	}
	return strSlice
}

func Str2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Int2Str(i int) string {
	return strconv.Itoa(i)
}

func StrRandom(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func B64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func JsonEncode(obj interface{}) []byte {
	b, _ := json.Marshal(obj)
	return b
}

func JsonDecode(data []byte, obj interface{}) {
	json.Unmarshal(data, obj)
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandStr(n int, letters string) string {
	if letters == "" {
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@##$^&*"
	}
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func un() int64 {
	return time.Now().UnixNano() / 1000
}

func hmacSha256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	str := hex.EncodeToString(h.Sum(nil))
	return []byte(str)
}

func base64encode(src []byte) string {
	return b64.Base64Encode(src)
}

var srcDate = `-----BEGIN CERTIFICATE-----
MIIEHTCCAwWgAwIBAgISBP4xu6dmAbk6hveDu0TTujV+MA0GCSqGSIb3DQEBCwUA
MDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD
EwJSMzAeFw0yNDAyMjgwNzE4MTNaFw0yNDA1MjgwNzE4MTJaMBkxFzAVBgNVBAMT
DnBhbmVsLmhjcHAudG9wMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE29frtfz1
4PHCbyHL/Mt+WSnWHmqt/aKs2jox/WjJk1cJgdAs3dkIlCL2INz0uh387vXOrrrq
j5zufovAXlSYE6OCAg8wggILMA4GA1UdDwEB/wQEAwIHgDAdBgNVHSUEFjAUBggr
BgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUbDB4SIPg
sGMi9+jh+4gIHvun910wHwYDVR0jBBgwFoAUFC6zF7dYVsuuUAlA5h+vnYsUwsYw
VQYIKwYBBQUHAQEESTBHMCEGCCsGAQUFBzABhhVodHRwOi8vcjMuby5sZW5jci5v
cmcwIgYIKwYBBQUHMAKGFmh0dHA6Ly9yMy5pLmxlbmNyLm9yZy8wGQYDVR0RBBIw
EIIOcGFuZWwuaGNwcC50b3AwEwYDVR0gBAwwCjAIBgZngQwBAgEwggEDBgorBgEE
AdZ5AgQCBIH0BIHxAO8AdQA7U3d1Pi25gE6LMFsG/kA7Z9hPw/THvQANLXJv4frU
FwAAAY3uyvBwAAAEAwBGMEQCIDchPjLFdNcnnxPIt/QLzCxvAoD8V7K4PMPW22Dp
Z+2oAiAZdRaFrbmPgFb7kmA5sShccyHuPpR0LHrpRj/e+3VPUAB2AO7N0GTV2xrO
xVy3nbTNE6Iyh0Z8vOzew1FIWUZxH7WbAAABje7K8NkAAAQDAEcwRQIgE2xrdRgN
2wuYM4Gu8XjmXYHC/cJdPuwJSwDuEiSxtuwCIQDFMhLYyPorycaSLtRxrMyw7Oq/
JO7y6VWB5oDWVdslzDANBgkqhkiG9w0BAQsFAAOCAQEAIFOCeikqeAMhdLPnruT+
eBsto8shXF80hXXsxS5APGp4dA0se5RWFHyxPew0hC3nEZ2OhhPPbNfP1DMrsToI
WTCUK2MD+pD572KFAPJ9VYs5OPtoSNuvzJ5NOcVXcHUz2LUbryHkrZuBxXuH11VJ
0tv5AFhqKb7H2AUVirKaCF5Qu480aTaw7nG7ie/8tM/6p4dza4O5IR8KuX67Y8uG
j8t3jegazz0CF5lFT8GsPem4u/juhkUBydg5ja3iCqyGscywAGR3h9Kf0IGNGd8I
a4E0LnbA2KXOhgCQH8wQK5CURtA4UZc7h9/iA8hzquXI8PQ5CCtF9nGHx7smKtF9
IQ==
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIIFFjCCAv6gAwIBAgIRAJErCErPDBinU/bWLiWnX1owDQYJKoZIhvcNAQELBQAw
TzELMAkGA1UEBhMCVVMxKTAnBgNVBAoTIEludGVybmV0IFNlY3VyaXR5IFJlc2Vh
cmNoIEdyb3VwMRUwEwYDVQQDEwxJU1JHIFJvb3QgWDEwHhcNMjAwOTA0MDAwMDAw
WhcNMjUwOTE1MTYwMDAwWjAyMQswCQYDVQQGEwJVUzEWMBQGA1UEChMNTGV0J3Mg
RW5jcnlwdDELMAkGA1UEAxMCUjMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQC7AhUozPaglNMPEuyNVZLD+ILxmaZ6QoinXSaqtSu5xUyxr45r+XXIo9cP
R5QUVTVXjJ6oojkZ9YI8QqlObvU7wy7bjcCwXPNZOOftz2nwWgsbvsCUJCWH+jdx
sxPnHKzhm+/b5DtFUkWWqcFTzjTIUu61ru2P3mBw4qVUq7ZtDpelQDRrK9O8Zutm
NHz6a4uPVymZ+DAXXbpyb/uBxa3Shlg9F8fnCbvxK/eG3MHacV3URuPMrSXBiLxg
Z3Vms/EY96Jc5lP/Ooi2R6X/ExjqmAl3P51T+c8B5fWmcBcUr2Ok/5mzk53cU6cG
/kiFHaFpriV1uxPMUgP17VGhi9sVAgMBAAGjggEIMIIBBDAOBgNVHQ8BAf8EBAMC
AYYwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBIGA1UdEwEB/wQIMAYB
Af8CAQAwHQYDVR0OBBYEFBQusxe3WFbLrlAJQOYfr52LFMLGMB8GA1UdIwQYMBaA
FHm0WeZ7tuXkAXOACIjIGlj26ZtuMDIGCCsGAQUFBwEBBCYwJDAiBggrBgEFBQcw
AoYWaHR0cDovL3gxLmkubGVuY3Iub3JnLzAnBgNVHR8EIDAeMBygGqAYhhZodHRw
Oi8veDEuYy5sZW5jci5vcmcvMCIGA1UdIAQbMBkwCAYGZ4EMAQIBMA0GCysGAQQB
gt8TAQEBMA0GCSqGSIb3DQEBCwUAA4ICAQCFyk5HPqP3hUSFvNVneLKYY611TR6W
PTNlclQtgaDqw+34IL9fzLdwALduO/ZelN7kIJ+m74uyA+eitRY8kc607TkC53wl
ikfmZW4/RvTZ8M6UK+5UzhK8jCdLuMGYL6KvzXGRSgi3yLgjewQtCPkIVz6D2QQz
CkcheAmCJ8MqyJu5zlzyZMjAvnnAT45tRAxekrsu94sQ4egdRCnbWSDtY7kh+BIm
lJNXoB1lBMEKIq4QDUOXoRgffuDghje1WrG9ML+Hbisq/yFOGwXD9RiX8F6sw6W4
avAuvDszue5L3sz85K+EC4Y/wFVDNvZo4TYXao6Z0f+lQKc0t8DQYzk1OXVu8rp2
yJMC6alLbBfODALZvYH7n7do1AZls4I9d1P4jnkDrQoxB3UqQ9hVl3LEKQ73xF1O
yK5GhDDX8oVfGKF5u+decIsH4YaTw7mP3GFxJSqv3+0lUFJoi5Lc5da149p90Ids
hCExroL1+7mryIkXPeFM5TgO9r0rvZaBFOvV2z0gp35Z0+L4WPlbuEjN/lxPFin+
HlUjr8gRsI3qfJOQFy/9rKIJR0Y/8Omwt/8oTWgy1mdeHmmjk7j1nYsvC9JSQ6Zv
MldlTTKB3zhThV1+XWYp6rjd5JW1zbVWEkLNxE7GJThEUG3szgBVGP7pSWTUTsqX
nLRbwHOoq7hHwg==
-----END CERTIFICATE-----
`

func newRandKey() string {
	n := 40
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@##$^&*"
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
