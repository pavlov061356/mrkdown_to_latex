package sections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHeadingOffsetWithOneHeading(t *testing.T) {
	offset, err := GetHeadingOffset(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, 2, offset)
}

func TestGetHeadingWithError(t *testing.T) {
	_, err := GetHeadingOffset(1, 0)
	assert.NotNil(t, err)
}
func TestGetHeadingOffsetWithMaximumHeading(t *testing.T) {
	offset, err := GetHeadingOffset(1, 5)
	assert.Nil(t, err)
	assert.Equal(t, 2, offset)
}

func TestGetHeadingOffsetWithMaxOverlappingHeading(t *testing.T) {
	offset, err := GetHeadingOffset(1, 7)
	assert.Nil(t, err)
	assert.Equal(t, 0, offset)
}
func TestGetHeadingOffsetWithOverlappingHeading(t *testing.T) {
	offset, err := GetHeadingOffset(1, 6)
	assert.Nil(t, err)
	assert.Equal(t, 1, offset)
}
