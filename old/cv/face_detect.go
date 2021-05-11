package cv

import (
	"os"

	"gocv.io/x/gocv"
)

type FaceDetect struct {
	classifier        gocv.CascadeClassifier
	dataDirectoryPath string
}

func NewFaceDetect() *FaceDetect {
	faceDetect := &FaceDetect{
		classifier:        gocv.NewCascadeClassifier(),
		dataDirectoryPath: os.Getenv("DATA_DIRECTORY_PATH"),
	}

	if !faceDetect.classifier.Load(faceDetect.dataDirectoryPath + "/classifiers/haarcascade_frontalface_default.xml") {
		panic("Cannot load classifiers!")
	}

	return faceDetect
}

func (f *FaceDetect) DetectFace(image []byte) (bool, error) {
	img, err := gocv.IMDecode(image, gocv.IMReadColor)

	if err != nil {
		return false, err
	}

	if rects := f.classifier.DetectMultiScale(img); len(rects) <= 0 {
		return false, nil
	}

	return true, nil
}
