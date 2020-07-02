package utils   //这个包名可以不是utils下，但是utils下的文件package声明要相同（与test1）

import "fmt"

func Count(){   //首字母大写别的包通过import导入此包之后可以使用这个
    fmt.Println("utils下的Count()")
}
func init(){   //init函数和main函数一样是内置函数，没有参数，没有返回值
               //比main函数先执行，用于初始化包所需要的特定资料,main函数的入口
               //只能由go程序自动调用，不能手动调用
    fmt.Println("utils包的初始化函数，用于一些信息的初始化")
}
func init(){   //init函数一个包下可以有多个（按顺序），也可以在不同包下（按import顺序）
    fmt.Println("另一个init()函数")
}