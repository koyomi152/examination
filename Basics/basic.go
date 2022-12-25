package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	username string
	nickname string
	age      uint8
	birthday string
}

/*
将sex改为age；用string表示生日日期；
b题的main函数在第9行返回了，所以后面的输出3未执行。
*/
func main() {
	str := fmt.Sprintf("%d-%d", time.Now().Month(), time.Now().Day())
	u := user{
		username: "坤坤",
		nickname: "阿坤",
		age:      20,
		birthday: str,
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
