package user

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"github.com/god-jason/bucket/config"
)

func passwordHash(password string) string {
	hash := config.GetString(MODULE, "password_hash")
	suffix := config.GetString(MODULE, "password_suffix")

	switch hash {
	case "md5":
		h := md5.New()
		h.Write([]byte(password + suffix))
		sum := h.Sum(nil)
		return hex.EncodeToString(sum)
	case "sha1":
		h := sha1.New()
		h.Write([]byte(password + suffix))
		sum := h.Sum(nil)
		return hex.EncodeToString(sum)
	case "sha256":
		h := sha256.New()
		h.Write([]byte(password + suffix))
		sum := h.Sum(nil)
		return hex.EncodeToString(sum)
	}
	return password
}
