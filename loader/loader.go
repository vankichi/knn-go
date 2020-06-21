package loader

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/kpango/glg"
	"github.com/vankichi/knn-go/util"
)

type Object struct {
	ID     int32
	Vector []float64
	Class  string
}

func New(s string) ([]*Object, error) {
	var data []*Object
	var err error = nil

	f, e := os.Open(s)
	defer f.Close()

	if e != nil {
		glg.Error(e)
		panic(e)
	}

	var id int32 = 0

	line := bufio.NewScanner(f)
	for line.Scan() {
		str := strings.Split(line.Text(), ",")
		// minimun len(str) is 1
		if len(str) <= 1 {
			break
		}

		var tmp = Object{
			ID:    id,
			Class: str[len(str)-1],
		}
		str = str[:len(str)-1]
		for _, v := range str {
			f64, err := strconv.ParseFloat(v, 64)
			if err != nil {
				glg.Error(err)
				panic(err)
			}
			tmp.Vector = append(tmp.Vector, f64)
		}
		data = append(data, &tmp)
		var oldId = id
		id = atomic.AddInt32(&oldId, 1)
	}
	if err = line.Err(); err != nil {
		glg.Error(err)
		panic(err)
	}
	return data, err
}

func Set(data []*Object, r float64) ([]*Object, []*Object) {
	// generate id list for test
	rand.Seed(time.Now().UnixNano())
	var testIds []int32
	var c int32 = 0
	for c < int32(float64(len(data))*r) {
		n := rand.Int31n(int32(len(data)))
		testIds = append(testIds, n)
		var tmp = c
		c = atomic.AddInt32(&tmp, 1)
	}
	// devide data into Train and Test
	var train []*Object
	var test []*Object
	for _, d := range data {
		if util.IntContains(testIds, d.ID) {
			test = append(test, d)
		} else {
			train = append(train, d)
		}
	}
	return train, test
}
