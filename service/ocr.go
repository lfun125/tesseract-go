package service

import (
	"context"

	"github.com/lfun125/tesseract-go/ocr"
	gosseract "github.com/otiai10/gosseract/v2"
)

// o *OCRService orc.OCRServiceServer
type OCRService struct {
	client *gosseract.Client
	ocr.UnsafeOCRServiceServer
}

func NewOCRService() *OCRService {
	client := gosseract.NewClient()

	return &OCRService{
		client: client,
	}
}

func (s *OCRService) OCR(ctx context.Context, req *ocr.OCRRequest) (*ocr.OCRResponse, error) {
	if err := s.client.SetLanguage(req.Langs...); err != nil {
		return nil, err
	}

	if err := s.client.SetImageFromBytes(req.Bytes); err != nil {
		return nil, err
	}

	text, err := s.client.Text()
	if err != nil {
		return nil, err
	}

	return &ocr.OCRResponse{
		Text: text,
	}, nil
}
