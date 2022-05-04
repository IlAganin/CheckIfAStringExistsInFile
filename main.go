package main

import ( "fmt"
          "io/ioutil"
          "regexp"
          "os"
)

func IsExist(str, filepath string) bool {
   b ,err := ioutil.ReadFile(filepath)
   if err != nil {
      panic(err)
   }

   isExist, err := regexp.Match(str, b)
   if err != nil{
      panic(err)
   }
   return isExist
}

func main(){
   var str string
   var file_path string
   fmt.Println("Enter your file_path")
   fmt.Fscan(os.Stdin, &file_path)
   fmt.Println("Enter string for existing:")
   fmt.Fscan(os.Stdin, &str)
   fmt.Println(IsExist(str, file_path))   
}