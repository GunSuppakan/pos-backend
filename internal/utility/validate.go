package utility

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"pos-backend/internal/infrastructure/logs"
	"strings"

	"github.com/spf13/viper"
)

type PaymentType string

const (
	PaymentCreditCard PaymentType = "credit_card"
	PaymentQR         PaymentType = "qr"
	PaymentCOD        PaymentType = "cod"
)

func IsValidPaymentType(p PaymentType) bool {
	switch p {
	case PaymentCreditCard, PaymentQR, PaymentCOD:
		return true
	default:
		return false
	}
}

func IsValidURL(url string) bool {
	if strings.Contains(url, "http") || strings.Contains(url, "https") {
		return true
	}
	return false
}

func IsImage(fileHeader *multipart.FileHeader, filename string) bool {
	maxMB := viper.GetInt("upload.image_max_mb")

	if maxMB <= 0 {
		maxMB = 5
	}

	maxSize := int64(maxMB) << 20
	if fileHeader != nil && fileHeader.Size > maxSize {
		logs.Error("File too large.")
		return false
	}
	allowedExt := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := strings.ToLower(filepath.Ext(filename))
	if !allowedExt[ext] {
		return false
	}

	f, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil {
		return false
	}

	mimeType := http.DetectContentType(buf[:n])
	allowedMime := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	return allowedMime[mimeType]
}

func IsPaymentType(typePayment string) bool {
	key := strings.ToLower(strings.TrimSpace(typePayment))

	const (
		PaymentCOD    = "cod"
		PaymentQR     = "qr"
		PaymentCredit = "credit"
	)

	if key == "qr" || key == "cod" || key == "credit" {
		return true
	} else {
		return false
	}
}
