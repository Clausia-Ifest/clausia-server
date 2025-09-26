package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"math/rand"
	"mime/multipart"
	"time"
)

func hashSHA256(fh *multipart.FileHeader) (string, error) {
	f, err := fh.Open()
	if err != nil {
		return "", errors.New("faled to open file")
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		f.Close()
		return "", errors.New("hash failed")
	}
	f.Close()
	sumHex := hex.EncodeToString(h.Sum(nil))

	return sumHex, nil
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
