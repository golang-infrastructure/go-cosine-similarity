package cosine_similarity

import (
	"fmt"
	"github.com/golang-infrastructure/go-gtypes"
	"math"
)

func NumberSliceCosineSimilarity[T gtypes.Number](sliceA, sliceB []T) float64 {
	r, _ := NumberSliceCosineSimilarityE(sliceA, sliceB)
	return r
}

// NumberSliceCosineSimilarityE 计算两个向量的余弦相似度
func NumberSliceCosineSimilarityE[T gtypes.Number](sliceA, sliceB []T) (float64, error) {
	if len(sliceA) != len(sliceB) {
		return 0, fmt.Errorf("length must equals")
	}
	var t1 float64
	var t2 float64
	var t3 float64
	for index := range sliceA {
		t1 += float64(sliceA[index]) * float64(sliceB[index])
		t2 += float64(sliceA[index]) * float64(sliceA[index])
		t3 += float64(sliceB[index]) * float64(sliceB[index])
	}
	fmt.Println(math.Sqrt(t2))
	fmt.Println(math.Sqrt(t3))
	return t1 / (math.Sqrt(t2) * math.Sqrt(t3)), nil
}
