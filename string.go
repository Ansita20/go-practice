package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println
func main() {
// 	sV1 := "A word"
// 	replacer := strings.NewReplacer("A", "Another")
// 	sV2 := replacer.Replace(sV1)
// 	pl(sV2)
// 	pl(len(sV1))
// 	pl(strings.Contains(sV2, "Another"))
	sV1 := "A Word"
	sV2 := strings.ToLower(sV1)
	pl(sV2)
	replacer := strings.NewReplacer("A", "Another")
	sV3 := strings.ToUpper(sV2);
	sV4 := replacer.Replace(sV1);
	pl(replacer.Replace(sV1))
	pl(sV3);
	pl("length: ", len(sV4));
	pl("length: ", len(sV1));
	pl("contains Another? ", strings.Contains(sV2, "Another"))
	pl("o Index? ", strings.Index(sV3,"o"))
	pl("Replace: ",strings.Replace(sV4,"o","0",-1))
	sV0 := "\n Some Place \n"
	sV0 = strings.TrimSpace(sV0)
	pl("Trimmed: ",sV0)
	pl("split: ",strings.Split("a,b,b,d,r", ","))
	pl("Prefix: ",strings.HasPrefix("totatoma","tota"))
	pl("Suffix: ",strings.HasSuffix("tomato","to"))
}

