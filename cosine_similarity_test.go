package cosine_similarity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntE(t *testing.T) {
	sliceA := []int{1, 1, 2, 1, 1, 1, 0, 0, 0}
	sliceB := []int{1, 1, 1, 0, 1, 1, 1, 1, 1}
	e, err := NumberSliceCosineSimilarityE(sliceA, sliceB)
	assert.Nil(t, err)
	t.Log(e)
}
