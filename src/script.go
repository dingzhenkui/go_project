package main

import (
	"fmt"
)

f, err := os.OpenFile("D:\Go_package\src\slice_02.go", os.O_RDONLY,0600)
    defer f.Close()
    if err !=nil {
        fmt.Println(err.Error())
    } else {
     contentByte,err=ioutil.ReadAll(f)
        checkErr(err)
fmt.Println(string(contentByte))
    }