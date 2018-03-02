package main
import ( 
  "fmt"
)

func main(){
  explicCallPanicTest()  //explicitly call panic() test
  //systemPanicTest()    //system panic test
  fmt.Println("+++ main exit. +++")
}

func systemPanicTest(){
  defer func(){ 
    if v := recover();v != nil { 
      fmt.Printf("--- Recovered a panic:%v\n",v)
    } 
    fmt.Printf("+++ test() defer.+++\n") 
  }() 
  myIndex := 4 
  ia := [3]int{1,2,3}
  _=ia[myIndex]
  fmt.Printf("+++ test exit. +++\n")
}

func explicCallPanicTest(){
  defer func(){
    if v := recover();v != nil {
      fmt.Printf("--- Recovered a panic. [index=%d]\n",v)  
    }
    fmt.Println("+++ getDemo() defer. +++") 
  }()
  str := []string{"a","b","c"} 
  fmt.Printf("+++ Get the elements in %v one by one.\n",str) 
  getElement(str,0) 
  fmt.Println("+++ getDemo() exit. +++")
}

func getElement(str []string,index int)(element string){
  defer func(){
  fmt.Printf("+++ getElement() defer. +++\n")
  }() 
  if index >= len(str) {
    fmt.Printf("--- There is a panic! [index=%d]\n",index)  
    panic(index) 
  }
  fmt.Printf("+++ Searching the element ... [index=%d]\n",index) 
  element = str[index] 
  defer fmt.Printf("+++ The element is %s. [index=%d]\n",element,index) 
  getElement(str,index+1) 
  return
}
