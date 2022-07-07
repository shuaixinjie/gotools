package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func MD5String(s string) string {
	md5hash := md5.New()
	md5hash.Write([]byte(s))
	return hex.EncodeToString(md5hash.Sum(nil))
}

func MD5File(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5hash := md5.New()
	_, err = io.Copy(md5hash, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5hash.Sum(nil)), nil
}
