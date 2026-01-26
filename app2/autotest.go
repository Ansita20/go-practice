package app2

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func IsEmail(s string)(string, error) {
	r, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,}
`)
	if r.MatchString(s){
		return "Valid Email", nil
	}else{
		return "", fmt.Errorf("not a valid email")
	}
}