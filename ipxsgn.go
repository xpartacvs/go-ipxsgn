package ipxsgn

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type imgsign struct {
	bytesKey  []byte
	bytesSalt []byte
}

type ImgSign interface {
	GetPath(c *Config, url string) (string, error)
}

var regexURL = regexp.MustCompile(`^(local|s3|gs|abs|https?)://.*`)

func New(key, salt string, keysaltAreEncoded bool) (ImgSign, error) {
	if keysaltAreEncoded {
		bKey, err := hex.DecodeString(key)
		if err != nil {
			return nil, err
		}

		bSalt, err := hex.DecodeString(salt)
		if err != nil {
			return nil, err
		}

		return &imgsign{bytesKey: bKey, bytesSalt: bSalt}, nil
	}

	bKey := []byte(key)
	bSalt := []byte(salt)

	return &imgsign{bytesKey: bKey, bytesSalt: bSalt}, nil
}

func validateConfig(c *Config) error {
	val := validator.New()
	err := val.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

func (i *imgsign) GetPath(c *Config, url string) (string, error) {
	var enlarge uint8 = 1
	if c.Enlarge == 0 {
		enlarge = 0
	}

	err := validateConfig(c)
	if err != nil {
		return "", err
	}

	if !regexURL.MatchString(url) {
		return "", errors.New("invalid url format")
	}
	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(url))

	var extension string
	if len(c.Extension) > 0 {
		extension = fmt.Sprintf(".%s", c.Extension)
	}

	path := fmt.Sprintf(
		"%s/%d/%d/%s/%d/%s%s",
		c.Resize,
		c.Width,
		c.Height,
		c.Gravity,
		enlarge,
		encodedURL,
		extension,
	)

	mac := hmac.New(sha256.New, i.bytesKey)

	_, err = mac.Write(i.bytesSalt)
	if err != nil {
		return "", err
	}

	_, err = mac.Write([]byte(path))
	if err != nil {
		return "", err
	}

	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s/%s", signature, path), nil
}
