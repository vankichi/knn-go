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
	Distance float64
}

var list []nn

type Set struct {
	*loader.Object
	Train []*loader.Object
}

func (s *Set) L2() (res []*nn, err error) {
	res = make([]*nn, 0, len(s.Train))
	for _, t := range s.Train {
		if len(t.Vector) != len(s.Vector) {
			glg.Error("vector size is not equal")
			return nil, err
		}
		var d float64 = 0
		for i := 0; i < len(s.Vector); i++ {
			d += math.Pow(s.Vector[i]-t.Vector[i], 2.0)
		}
		res = append(res, &nn{
			ID:       t.ID,
			Class:    t.Class,
			Distance: math.Sqrt(d),
		})
	}
	return res, err
}

func Knn(n []*nn, k int32) []*nn {
	sort.Slice(n, func(i, j int) bool {
		return n[i].Distance < n[j].Distance
	})
	return n[:k+1]
}

func PreClass(knn []*nn) string {
	l := make([]string, 0, len(knn))
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
	re := make([]*r, 0, len(l))
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
		return re[i].count > re[j].count
	})
	return re[0].name
}
