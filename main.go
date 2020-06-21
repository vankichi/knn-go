package main

import (
	"fmt"

	"github.com/kpango/glg"
	"github.com/vankichi/knn-go/knn"
	"github.com/vankichi/knn-go/loader"
)

const ratio float64 = 0.1

const file = "assets/iris.data"
const K int32 = 2

func main() {
	d, err := loader.New(file)
	if err != nil {
		glg.Error(err)
		panic(err)
	}
	train, test := loader.Set(d, ratio)
	var precision float64 = 0.0
	for _, t := range test {
		var set = knn.Set{Object: t, Train: train}
		nn, err := set.L2()
		if err != nil {
			panic(err)
		}
		list := knn.Knn(nn, K)
		pc := knn.PreClass(list)
		if set.Class == pc {
			precision++
		}
	}
	precision = precision / float64(len(test))
	fmt.Println(precision)
}
