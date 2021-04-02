package main
import (
    "fmt"
)
func main() {

    var x[4]int = [4]int{1,2,3,4}
    var y = x
    y[3] = 100
    fmt.Println(x, y)
    fmt.Println("---------------")

    var x2[]int = []int{1,2,3,4}
    var y2 = x2
    y2[3] = 100
    fmt.Println(x2, y2)
    fmt.Println("---------------")

    func (v Value) Interface() interface {}
    var test int
    fmt.Println(test.interface().Interface())
}


