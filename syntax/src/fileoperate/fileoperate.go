package main

import (
	"io/ioutil"
	"fmt"
)

func main(){
	SearchFileInDir("E:/",".java")
}

/*
*  特定格式文件搜素 
*
*/

func SearchFileInDir(dir string,suffix string){
	
	files,err := ioutil.ReadDir(dir)
	if (err!=nil){
		fmt.Println(err)
	}
	fmt.Println(files)
}
