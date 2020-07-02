package main
import (
  "fmt"
  "os/exec"
//  "io"
  "bytes"
  "log"
  "strings"
)
func main(){
  f,err := exec.LookPath("ls")
  if err != nil{
    fmt.Println(err)
  }
  fmt.Println(f)

  cmd := exec.Command(f)
  cmd.Stdin = strings.NewReader("some input")
  var out bytes.Buffer
  cmd.Stdout = &out
  err =cmd.Run()
  if err != nil{
    log.Fatal(err)  
 }
  fmt.Printf("in all caps:%q\n",out.String())

}

//type cmd struct{
  //  Stdin io.Reader
//}
