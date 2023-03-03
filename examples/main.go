package main

import (
    "encoding/json"
    "fmt"
)

type Results struct {
    Code int `json:"code,omitempty" xml:"code,omitempty"` //转换为json时 键 Code 变为 code 如果为 0 则忽略  转换为xml时 同样的规则
    Method string `json:"method"` //转换为json时 键 Method 变为 method
    Msg string `json:"msg,omitempty"` //转换为json时 键 Msg 变为 msg 如果为 空字符串 则忽略
}

func main() {
    results := Results{
        Method:"POST",
        Msg:"",
    }
    jsonStr,_:=json.Marshal(results)
    fmt.Println(string(jsonStr))
    // 输出结果为 {"method":"POST"}  Code没赋值默认为0
    // omitempty表省略，如果code没有加omitempty，输出结果为{"method":"POST", "code": 0}
}

