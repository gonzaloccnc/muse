package utils_test

import (
	"muse/utils"
	"testing"
)

func TestValidatePaths(t *testing.T) {
	paths := []string{
		"as93-.@/",
		"/c/route",
		"c/PATH123@/",
		"c/c123",
		"c\\s/fg/",
		"c-/route",
		"c/route",
		"c/sap/pel/ads\\AS\\sa",
		"c\\route\\",
		"c\\s//pe\\",
		"\\c\\a\\s",
		"\\\\c\\\\\\\\\\\\a\\s",
		"c:\\as\\spe\\sa",
		"c:route",
		"/a-\\k",
	}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			want := utils.ValidatePath(path)
			if !want {
				t.Errorf("path is not valid: %s", path)
			}
		})
	}
}
