// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gostring

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
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
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
