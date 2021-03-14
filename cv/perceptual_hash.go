package cv

import (
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

type PerceptualHash struct {
}

func NewPerceptualHash() *PerceptualHash {
	return &PerceptualHash{}
}

func (p *PerceptualHash) GetHashFromImage(image []byte) ([]byte, error) {
	var hash contrib.PHash

	result := gocv.NewMat()

	img, err := gocv.IMDecode(image, gocv.IMReadColor)

	if err != nil {
		return nil, err
	}

	hash.Compute(img, &result)

	return result.ToBytes(), nil
}

func (p *PerceptualHash) CompareHashes(hash1 []byte, hash2 []byte) float64 {
	var hash contrib.PHash

	hash1mat, _ := gocv.NewMatFromBytes(1, 8, gocv.MatTypeCV8U, hash1)
	hash2mat, _ := gocv.NewMatFromBytes(1, 8, gocv.MatTypeCV8U, hash2)

	return hash.Compare(hash1mat, hash2mat)
}
