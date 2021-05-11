package cv

type CV struct {
	PerceptualHash *PerceptualHash
	FaceDetect     *FaceDetect
}

func New() *CV {
	return &CV{
		PerceptualHash: NewPerceptualHash(),
		FaceDetect:     NewFaceDetect(),
	}
}
