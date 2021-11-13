package ipxsgn

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"regexp"
	"strconv"
)

type ImgSign struct {
	bytesKey  []byte
	bytesSalt []byte
}

var regexURL *regexp.Regexp = regexp.MustCompile(`^(local|s3|gs|abs|https?)://.*`)

func New(key, salt string, keysaltAreEncoded bool) (*ImgSign, error) {
	if keysaltAreEncoded {
		bKey, err := hex.DecodeString(key)
		if err != nil {
			return nil, err
		}

		bSalt, err := hex.DecodeString(salt)
		if err != nil {
			return nil, err
		}

		return &ImgSign{bytesKey: bKey, bytesSalt: bSalt}, nil
	}

	bKey := []byte(key)
	bSalt := []byte(salt)

	return &ImgSign{bytesKey: bKey, bytesSalt: bSalt}, nil
}

func (i ImgSign) GetPath(c *Config, url string) (string, error) {
	var enlarge uint8 = 1
	if c.enlarge == 0 {
		enlarge = 0
	}

	if !regexURL.MatchString(url) {
		return "", errors.New("invalid url format")
	}

	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(url))

	var extension string
	if len(c.extension) > 0 {
		extension = "." + string(c.extension)
	}

	path := string(c.resize) + "/" + strconv.FormatUint(uint64(c.width), 10) + "/" + strconv.FormatUint(uint64(c.height), 10) + "/" + string(c.gravity) + "/" + strconv.FormatUint(uint64(enlarge), 10) + "/" + encodedURL + extension

	mac := hmac.New(sha256.New, i.bytesKey)
	if _, err := mac.Write(i.bytesSalt); err != nil {
		return "", err
	}
	if _, err := mac.Write([]byte(path)); err != nil {
		return "", err
	}

	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return "/" + signature + "/" + path, nil
}
