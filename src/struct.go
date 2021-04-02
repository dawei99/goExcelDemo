// struct 结构体
package main
import (
    "fmt"
    "reflect"
)

func main(){

     var p point
     p.x = 123
     var px  = p
     px.x = 10
     fmt.Println("---------p----------")
     fmt.Println(p)

     var p2 = new(point)
     p2.x = 1
     p2.y = 2
     var pxxx = p2
     pxxx.x = 3200
     fmt.Println("---------p2----------")
     fmt.Println(p2)
     fmt.Println(reflect.TypeOf(p2))

     var p3 = &point{}
     p3.x = 33
     p3.y = 3
     fmt.Println("---------p3----------")
     fmt.Println(p3)

    // 键值对初始化
     var p4 = &point{
        x: 5,
        y: 6,
        s: "init string",
     }
     fmt.Println("---------p4----------")
     var px2 = p4
     px2.x = 100
     fmt.Println(p4)

    // 结构体指针成员
     p5 := &point {
        s: "D1",
        child: &point{
            s: "D2",
            child: &point{
                s: "d3",
            },
        },
     }
     fmt.Println("---------p5----------")
     fmt.Println(p5.child.child.s)

    //  无建名字初始化
    p6 := &point {
        1,
        2,
        "zi fu chuan",
        &point{},
    }
    fmt.Println("---------p6----------")
    zhizhenChild := p6.child
    zhizhenChild.s = "new value"
    fmt.Println(p6.child)

    // 匿名结构体
    msg := &struct {
        name string
        sex int
    } {
        name: "小明",
        sex: 18,
    }
     fmt.Println("---------msg----------")
    fmt.Println(msg)

    // 构造函数
    fmt.Println("-------构造------")
    parentPoint := new(point)
    fmt.Println(pointCreate(10, 20, "z30", parentPoint))

    // 派生构造
    fmt.Println("-------派生构造------")
    fmt.Println(chinaCatNew(map[string]int{"width":10,"height":20},"red"))

    fmt.Println("-------生成方式------")
    fmt.Println(chinaCat{
                                 name: "TEXT",
                             })
}

type point struct {
    x int
    y int
    s string
    child *point
 }

// 构造函数
func pointCreate (x, y int, s string, child *point) *point {
    var result *point = new(point)
    result.x = x
    result.y = y
    result.s = s
    result.child = child
    return result
}

type BaseBase struct {
    baseProperty float64
}
// 派生构造
type BaseCat struct {
    BaseBase
    size map[string]int  //大小
}
type chinaCat struct {
    BaseCat
    name string
}
func chinaCatNew(size map[string]int , name  string)*chinaCat{
    var chinaCat = &chinaCat{}
    chinaCat.size = size
    chinaCat.name = name
    chinaCat.baseProperty = 5.6
    return chinaCat
}
