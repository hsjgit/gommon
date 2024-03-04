package main

import (
	"fmt"

	"github.com/hsjgit/gommon/exec"
	"github.com/hsjgit/gommon/google"
)

func main() {
	auth := google.NewGoogleAuth()
	secret := auth.GetSecret()
	fmt.Println(secret)
	name := fmt.Sprintf("test-%s", "hsj")
	url := auth.GetQrcodeUrl(name, secret)
	fmt.Println(url)
	exec.Shell(fmt.Sprintf("open %s", url))
	for {
		code := ""
		_, err := fmt.Scan(&code)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		verifyCode, err := auth.VerifyCode(secret, code)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if verifyCode {
			fmt.Println("success")
		} else {
			fmt.Println("fail")
		}
	}

}
