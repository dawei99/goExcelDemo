package main
import (
    "fmt"
//    "reflect"
    "strings"
)
func main() {
    var day, hours, minute = toThreeDate(846000)
    fmt.Printf("天:%v 时:%v 分:%v \n", day, hours, minute)

    fmt.Println(test())
    fmt.Println("-------------")

    var pi = myInt{
        1,
        2,
        new(myString),
    }
    fmt.Printf("%v \n", &pi)
    zhuanFunc(pi)

    var _ func(int, string)int = abc

    //var x [1]func(int, string)int = [1]func(int, string)int{abc}
    var qx []func(int, string)int = []func(int, string)int{abc}
    var new = qx
    new[0] = func (int, string) int {
        return 2
    }
    fmt.Println(qx[0](1, "b"))

    var textString string = "table_abc"
    fmt.Println(strings.TrimPrefix(textString, "table_"))
}

func abc (a int, b string) int {
    return 1
}

func zhuanFunc (myint myInt) int {
    fmt.Printf("%v \n", &myint)
    return 1
}

const (
    MinuteSecond = 60
    HoursSecond = MinuteSecond * 60
    DayHours = HoursSecond * 24
)
func toThreeDate (second int)  (day, hours, minute int) {
    day = second / DayHours
    hours = second / HoursSecond
    minute = second / MinuteSecond
    return
}

func test () myInt {
    var myint = myInt{
        10,
        20,
        new(myString),
    }
    var myint2 = myint
    myint2.max = 1
    fmt.Println("myint",myint)
    fmt.Println("myint2",myint2)
    return myint
}

type myInt struct {
    max int
    min int
    str *myString
}

type myString struct {
    value string
}


