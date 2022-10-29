package rdir

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//ScanDir takes the path of the directory it should scan,
//first it prints the current dir, so we know what dir we are working
//then it calls ReadDir method on that path to get the dir contents
// loops over the slice of FileInfo values ReadDir returns,
//processing each one.
//it calls filepath.Join to join the current dir path and current filename together with slashes
//(so "go" and "src" are joined to become "go/src")
//if the current file isn't directory, ScanDir just prints its full path
//and moves on to the next file(if there are anymore in the current directory)
//if the current file is dir, the recursion kicks in: ScanDir calls itself
//with the subdirectories' path
//if that subdirectory has any subdirectories, ScanDir will call itself
func ScanDir(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	//join dir path and file name with a slash
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			ScanDir(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

//ReportPanic if recover return nothing there is no panic
//otherwise, get the underlying type if  error value and print it
// if not error then call the panic again
func ReportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err) //when the panic takes an error value
	} else {
		panic(p)
	}
}
