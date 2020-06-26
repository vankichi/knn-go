package loader

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kpango/glg"
	"github.com/vankichi/knn-go/util"
)

type Object struct {
	ID     int32
	Vector []float64
	Class  string
}

func New(path string) (datas []*Object, err error) {
	f, err := os.Open(path)

	if err != nil {
		glg.Error(err)
		return nil, err
	}
	defer f.Close()

	line := bufio.NewScanner(f)
	for i := int32(0); line.Scan(); i++ {
		str := strings.Split(line.Text(), ",")
		// minimun len(str) is 1
		if len(str) <= 1 {
			break
		}
		var data = Object{
			ID:    i,
			Class: str[len(str)-1],
		}
		str = str[:len(str)-1]
		for _, v := range str {
			f64, err := strconv.ParseFloat(v, 64)
			if err != nil {
				glg.Error(err)
				return nil, err
			}
			data.Vector = append(data.Vector, f64)
		}
		datas = append(datas, &data)
	}
	if err = line.Err(); err != nil {
		glg.Error(err)
		return nil, err
	}
	return datas, err
}

func Set(data []*Object, r float64) ([]*Object, []*Object) {
	// generate id list for test
	rand.Seed(time.Now().UnixNano())
	testIds := make([]int32, 0, len(data))
	for c := int32(0); c < int32(float64(len(data))*r); c++ {
		testIds = append(testIds, rand.Int31n(int32(len(data))))
	}
	// devide data into Train and Test
	train := make([]*Object, 0, len(data))
	test := make([]*Object, 0, len(data))
	for _, d := range data {
		if util.IntContains(testIds, d.ID) {
			test = append(test, d)
		} else {
			train = append(train, d)
		}
	}
	return train, test
}
