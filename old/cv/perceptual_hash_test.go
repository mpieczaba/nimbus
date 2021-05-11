package cv

import (
	"fmt"
	"testing"

	"github.com/mpieczaba/nimbus/filesystem"
)

func TestGetHashFromImage(t *testing.T) {
	cv := New()
	fs := filesystem.NewFilesystem()

	img1, err := fs.ReadFile("test1")

	if err != nil {
		t.Errorf("Cannot get image from file!")
	}

	img2, err := fs.ReadFile("test2")

	if err != nil {
		t.Errorf("Cannot get image from file!")
	}

	hash1, err := cv.PerceptualHash.GetHashFromImage(img1)

	if err != nil {
		t.Errorf("Cannot get hash!")
	}

	hash2, err := cv.PerceptualHash.GetHashFromImage(img2)

	if err != nil {
		t.Errorf("Cannot get hash!")
	}

	fmt.Println(hash1)
	fmt.Println(hash2)
}

func TestCompareHashes(t *testing.T) {
	cv := New()

	similarity := cv.PerceptualHash.CompareHashes([]byte{51, 204, 204, 51, 51, 204, 206, 51}, []byte{246, 57, 67, 18, 177, 44, 15, 196})

	fmt.Println(similarity)
}
