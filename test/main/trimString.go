package main
import (
"fmt"
"strings"
)
func main (){
    var list []string = []string {"pre_user", "pre_from"}
    var hands []func(string) string = []func(string)string {
        trim,
        append,
    }
    newList := actionForHand(list, hands)
    fmt.Println(list , hands)
    fmt.Println(newList)
}

// 去掉前缀
func trim (str string) string {
    return strings.TrimPrefix(str, "pre_")
}

// 添加后缀
func append (str string) string {
    return str + "_end"
}

// 链式处理
func actionForHand (list []string, handles[]func(string)string) []string {
    for key, readyString := range list {
        var resultString string = readyString
        for _, handFunc := range handles {
            resultString = handFunc(resultString)
        }
        list[key] = resultString
    }
    return list
}
