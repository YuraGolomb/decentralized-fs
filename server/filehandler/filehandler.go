package filehandler

import (
	"io"
	"math/rand"
	"os"

	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

const bufferSize = 256
const keySize = 32

// DecodeFile uploads file
func DecodeFile(path string, keyPath string) (e error) {
	buffer := make([]byte, bufferSize)
	startFile, e := os.Open(path)
	endFile, e := os.Create(path + "_dec")
	keyFile, e := os.Open(keyPath)

	key := make([]byte, keySize)

	_, e = keyFile.Read(key)

	defer startFile.Close()
	defer endFile.Close()
	defer keyFile.Close()

	if e != nil {
		astilog.Debug("Files open failed")
		astilog.Fatal(errors.Wrap(e, "error"))
		return
	}
	for {
		bytesread, err := startFile.Read(buffer)

		if err != nil {
			if err != io.EOF {
				astilog.Debug("Files read failed")
				astilog.Fatal(errors.Wrap(err, "error"))
				return err
			}
			break
		}
		stringread := string(buffer[:bytesread])

		encryptedString, err := decrypt(key, stringread)

		if err != nil {
			astilog.Debug("Files decrypt failed")
			astilog.Fatal(errors.Wrap(e, "error"))
			return err
		}
		_, err = endFile.WriteString(encryptedString)
		if err != nil {
			astilog.Debug("Files wrtite failed")
			astilog.Fatal(errors.Wrap(e, "error"))
			return err
		}
	}
	return nil
}

// EncodeFile uploads file
func EncodeFile(path string) (e error) {
	buffer := make([]byte, bufferSize)
	key := randStringBytesRmndr(keySize)
	startFile, e := os.Open(path)
	endFile, e := os.Create(path + "_sec")
	kf, e := os.Create(path + "_key")

	defer startFile.Close()
	defer endFile.Close()
	defer kf.Close()

	if e != nil {
		astilog.Fatal(errors.Wrap(e, "error"))
		return
	}

	_, e = kf.WriteString(string(key))
	if e != nil {
		astilog.Fatal(errors.Wrap(e, "error"))
		return
	}

	for {
		bytesread, err := startFile.Read(buffer)

		if err != nil {
			if err != io.EOF {
				astilog.Fatal(errors.Wrap(err, "error"))
				return err
			}
			break
		}
		stringread := string(buffer[:bytesread])

		encryptedString, err := encrypt(key, stringread)

		if err != nil {
			astilog.Fatal(errors.Wrap(e, "error"))
			return err
		}
		_, err = endFile.WriteString(encryptedString)
		if err != nil {
			astilog.Fatal(errors.Wrap(e, "error"))
			return err
		}
	}
	return nil
}

// func check(e error) {
// 	if e != nil {
// 		astilog.Fatal(errors.Wrap(e, "error"))
// 	}
// }

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0132456789"

func randStringBytesRmndr(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}
