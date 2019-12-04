package main

import (
	"fmt"
	"export/excel"
	"export/model"
	"sync"
)

func main () {
	fmt.Println("RUN GO")
	w := sync.WaitGroup{}
	for i :=0;i < 42;i++ {
		w.Add(1)
		go func(page int) {
			defer w.Done()
			data := model.GetData(page,1000)
			excel.Export(data,page)
		}(i)
	}
	w.Wait()
}
