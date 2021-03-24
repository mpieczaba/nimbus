package cv

import (
	"fmt"
	"testing"

	"github.com/mpieczaba/nimbus/filesystem"
)

func TestDetectFace(t *testing.T) {
	cv := New()
	fs := filesystem.NewFilesystem()

	img, err := fs.ReadFile("test3")

	if err != nil {
		t.Errorf("Cannot get image from file!")
	}

	result, err := cv.FaceDetect.DetectFace(img)

	if err != nil {
		t.Errorf("Cannot detect face!")
	}

	fmt.Println(result)
}
