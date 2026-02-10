package utility

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"image/png"
	"io"
	"math/big"
	"regexp"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(bytes), err
}

func GenerateUserID() (string, error) {
	const digits = 10

	result := make([]byte, digits)
	for i := 0; i < digits; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		result[i] = byte('0' + n.Int64())
	}

	if result[0] == '0' {
		result[0] = '1'
	}

	return string(result), nil
}

func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	) == nil
}

func HashPath(parts ...string) string {
	data := strings.Join(parts, "/")
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func GenerateBarcodeImage(
	value string,
	width int,
	height int,
	writer io.Writer,
) error {

	bar, err := code128.Encode(value)
	if err != nil {
		return err
	}

	barScaled, err := barcode.Scale(bar, width, height)
	if err != nil {
		return err
	}

	return png.Encode(writer, barScaled)
}

func NormalizeCategoryKey(input string) string {
	key := strings.ToLower(strings.TrimSpace(input))
	key = strings.ReplaceAll(key, " ", "-")
	key = regexp.MustCompile(`[^a-z0-9\-]`).ReplaceAllString(key, "")
	return key
}
