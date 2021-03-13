package cv

type CV struct {
	PerceptualHash *PerceptualHash
}

func New() *CV {
	return &CV{
		PerceptualHash: NewPerceptualHash(),
	}
}
