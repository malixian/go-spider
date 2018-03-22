package main

import "github.com/pkg/errors"

func main(){
	const DEFINEERROR error = errors.New("test")
	if(true){
		defer println("1")
	}else{
		defer println("2")
	}
	println("3")
}
