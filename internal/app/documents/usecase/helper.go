package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
)

func assertPDF(fh *multipart.FileHeader) error {
	if mt, _, err := mime.ParseMediaType(fh.Header.Get("Content-Type")); err != nil || mt != "application/pdf" {
		return errors.New("only PDF files are allowed")
	}

	f, err := fh.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return err
	}
	if ct := http.DetectContentType(buf[:n]); ct != "application/pdf" {
		return errors.New("file content is not a PDF")
	}

	if s, ok := f.(io.Seeker); ok {
		_, _ = s.Seek(0, io.SeekStart)
	}
	return nil
}

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
