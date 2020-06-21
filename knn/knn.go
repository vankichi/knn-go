package knn


import (
	"math"
	"sort"

	"github.com/kpango/glg"
	"github.com/vankichi/knn-go/loader"
	"github.com/vankichi/knn-go/util"
)

type nn struct {
	ID       int32
	Class    string
	distance float64
}

var list []nn

type Set struct {
	*loader.Object
	Train []*loader.Object
}

func (s *Set) L2() ([]*nn, error) {
	var res []*nn
	var err error = nil
	for _, t := range s.Train {
		if len(t.Vector) != len(s.Vector) {
			glg.Error("vector size is not equal")
			panic("vector size is not equal")
		}
		var d float64 = 0
		for i := 0; i < len(s.Vector); i++ {
			d += math.Pow(s.Vector[i]-t.Vector[i], 2.0)
		}
		res = append(res, &nn{
			ID:       t.ID,
			Class:    t.Class,
			distance: math.Sqrt(d),
		})
	}
	return res, err
}

func Knn(n []*nn, k int32) []*nn {
	sort.Slice(n, func(i, j int) bool {
		if n[i].distance < n[j].distance {
			return true
		}
		return false
	})
	knn := n[:k+1]
	return knn
}

func PreClass(knn []*nn) string {
	var l []string
	// pick up class name
	for _, n := range knn {
		if !util.StrContains(l, n.Class) {
			l = append(l, n.Class)
		}
	}
	// count up class name
	type r struct {
		name  string
		count int
	}
	var re []*r
	for _, c := range l {
		var count int = 0
		for _, n := range knn {
			if n.Class == c {
				count++
			}
		}
		re = append(re, &r{
			name:  c,
			count: count,
		})
	}

	sort.Slice(re, func(i, j int) bool {
		if re[i].count > re[j].count {
			return true
		}
		return false
	})
	return re[0].name
}
