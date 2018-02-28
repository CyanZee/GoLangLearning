package main
import ( 
  "fmt"
)
func main(){
  getDemo() 
  fmt.Println("+++ main exit. +++")
}
func getDemo(){
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
