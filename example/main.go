package main

import (
	"fmt"
	cosine_similarity "github.com/golang-infrastructure/go-cosine-similarity"
)

func main() {
	sliceA := []int{1, 1, 2, 1, 1, 1, 0, 0, 0}
	sliceB := []int{1, 1, 1, 0, 1, 1, 1, 1, 1}
	r, err := cosine_similarity.NumberSliceCosineSimilarityE(sliceA, sliceB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output:
	// 0.7071067811865475
}
