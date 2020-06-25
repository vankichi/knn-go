package main

import (
	"github.com/kpango/glg"
	"github.com/vankichi/knn-go/knn"
	"github.com/vankichi/knn-go/loader"
)

const ratio float64 = 0.1

const file = "assets/iris.data"
const K int32 = 3

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

		glg.Infof("test data vector: %v", set.Vector)
		for i := 0 ; i < int(K) ; i++ {
			glg.Infof("%dnn vector : { class: %s, distance: %.5f }", i+1, list[i].Class, list[i].Distance)
		}
		pc := knn.PreClass(list)
		if set.Class == pc {
			precision++
		}
		glg.Infof("{correct ClaasName: %s, predicted ClassName: %s}", set.Class, pc)
	}
	precision = precision / float64(len(test))
	glg.Infof("accuracy: %.2f", precision)
}
