package sections

import (
	"fmt"

	"github.com/pavlov061356/markdown_to_latex/internal/constants"
)

// func GetSectionName() string {

// }

const defaultSectionOffset = 2

func GetHeadingCount(heading string) int {
	var headingCount int

	for _, section := range heading {
		if section == '#' {
			headingCount++
		}
	}

	return headingCount
}

func GetHeadingOffset(minHeadingCount, maxHeadingCount int) (int, error) {

	if maxHeadingCount < minHeadingCount {
		return 0, fmt.Errorf("max heading count cannot be less than min heading count")
	}

	return max(min(len(constants.Sections)-(maxHeadingCount-minHeadingCount)-1, defaultSectionOffset), 0), nil

}
