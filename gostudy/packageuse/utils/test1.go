package utils

import "fmt"

//同包下的访问不用导入
func MyTest2(){
    Count()
}

func init(){
    fmt.Print("这是test1下面的init()")
}