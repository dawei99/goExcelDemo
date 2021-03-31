package main
import ("fmt")
func main () {
    var a = test()
    fmt.Println(a)
}

func test () (int, int){
    return 1,2
}
