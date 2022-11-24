package better

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	filter := &filter{}
	filter.Add("Thomas Shen")
	filter.Add("Rick Wu")
	filter.Add("Johnny Zhou")
	filter.Add("Shengfa Wang")

	assert.True(t, filter.Exist("Thomas Shen"))
	assert.True(t, filter.Exist("Rick Wu"))
	assert.True(t, filter.Exist("Johnny Zhou"))
	assert.True(t, filter.Exist("Shengfa Wang"))

	assert.False(t, filter.Exist("Elon Musk"))

}
