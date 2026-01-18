package imageProcessors

import (
	"bytes"

	"github.com/disintegration/imaging"
)

type ImageProcessors interface {
	ResizeImage(img []byte) ([]byte, error)
}

type processor struct {
}

func NewImageProcessor() ImageProcessors {
	return &processor{}
}

func (p *processor) ResizeImage(img []byte) ([]byte, error) {
	reader := bytes.NewReader(img)
	decodedImage, err := imaging.Decode(reader)
	if err != nil {
		return nil, err
	}

	resizedImg := imaging.Resize(decodedImage, 150, 150, imaging.Lanczos)

	var buf bytes.Buffer
	err = imaging.Encode(&buf, resizedImg, imaging.PNG)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
