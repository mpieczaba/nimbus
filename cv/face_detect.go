package cv

import (
	"gocv.io/x/gocv"
)

type FaceDetect struct {
	classifier gocv.CascadeClassifier
}

func NewFaceDetect() *FaceDetect {
	faceDetect := &FaceDetect{
		classifier: gocv.NewCascadeClassifier(),
	}

	if !faceDetect.classifier.Load("./classifiers/haarcascade_frontalface_default.xml") {
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
