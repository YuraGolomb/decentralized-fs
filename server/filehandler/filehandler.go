package filehandler

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

const bufferSize = 1024
const keySize = 32
const blockSize = 16 // 16bytes
// FileDescriptor struct
type FileDescriptor struct {
	Size    int64  `json:"Size"`
	Name    string `json:"Name"`
	IsDir   bool   `json:"IsDir"`
	ModTime string `json:"ModTime"`
	Preview string `json:"Preview"`
}

// GetFileInfo fileinfo
func GetFileInfo(path string) (fileDescription FileDescriptor, e error) {
	buffer := make([]byte, bufferSize)
	startFile, e := os.Open(path)

	defer startFile.Close()

	fileStat, e := startFile.Stat()
	bytesread, e := startFile.Read(buffer)

	if e != nil {
		astilog.Error(errors.Wrap(e, "error"))
		return
	}

	filePreview := string(buffer[:bytesread])
	fileDescription = FileDescriptor{
		Size:    fileStat.Size(),
		Name:    fileStat.Name(),
		IsDir:   fileStat.IsDir(),
		ModTime: fileStat.ModTime().String(),
		Preview: filePreview,
	}
	return
}

// DecodeFile uploads file
func DecodeFile(path string, keyPath string) (e error) {
	fmt.Println("Decoding file")
	decryptedFile, e := os.Create(keyPath + ".dec")
	if e != nil {
		astilog.Error("error while crating .dec file")
		astilog.Error(errors.Wrap(e, "error"))
		return
	}
	defer decryptedFile.Close()

	keyFile, e := os.Open(keyPath)
	if e != nil {
		astilog.Error("error while opening .key file")
		astilog.Error(errors.Wrap(e, "error"))
		return
	}
	defer keyFile.Close()

	reader := bufio.NewReader(keyFile)
	keyS, e := reader.ReadString('\n')
	if e != nil {
		astilog.Error("error while reading .key file")
		astilog.Error(errors.Wrap(e, "error"))
		return
	}

	key := []byte(keyS[:len(keyS)-2])

	var pathes [32]string
	for i := 0; i < 32; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			astilog.Error("error while reading from .key filenames")
			astilog.Error(errors.Wrap(e, "error"))
			return err
		}
		pathes[i] = line[:len(line)-2]
	}

	decodingBuffer := CreateDecodingBuffer(pathes)

	for !decodingBuffer.isFinished {
		stringToDecrypt, err := decodingBuffer.ReadNextBufferPart()
		if err != nil {
			astilog.Error("error while reading buffer")
			astilog.Error(errors.Wrap(e, "error"))
			return err
		}
		fmt.Println("string to decode")
		fmt.Println(stringToDecrypt)
		fmt.Println(len(stringToDecrypt))
		decryptedString, err := decrypt(key, stringToDecrypt)
		if err != nil {
			astilog.Error("error while decrypting buffer")
			fmt.Println(err)
			return err
		}
		fmt.Println("descrypted string")
		fmt.Println(decryptedString)
		decryptedFile.WriteString(decryptedString)
	}
	return nil
}

// EncodeFile uploads file
func EncodeFile(path string) (e error) {
	fmt.Println("Encoding file")
	key := randStringBytesRmndr(keySize)
	buffer := make([]byte, bufferSize)
	keyFile, e := os.Create(path + ".key")
	initialFile, e := os.Open(path)
	encodedFile, e := os.Create(path + ".sec")

	defer keyFile.Close()
	defer initialFile.Close()

	if e != nil {
		astilog.Error(errors.Wrap(e, "error"))
		return
	}
	// write key to file
	_, e = keyFile.WriteString(string(key) + "\r\n")
	if e != nil {
		astilog.Error(errors.Wrap(e, "error"))
		return
	}

	for {
		// read file by buffer parts
		bytesread, err := initialFile.Read(buffer)

		if err != nil {
			if err != io.EOF {
				astilog.Error(errors.Wrap(err, "error"))
				return err
			}
			break
		}

		stringread := string(buffer[:bytesread])
		encryptedString, err := encrypt(key, stringread)
		fmt.Println(encryptedString)
		fmt.Println(len(encryptedString))

		if err != nil {
			astilog.Error(errors.Wrap(e, "error"))
			return err
		}
		_, err = encodedFile.Write(encryptedString)
		if err != nil {
			astilog.Error(errors.Wrap(e, "error"))
			return err
		}
	}
	e = encodedFile.Sync()
	e = encodedFile.Close()
	if e != nil {
		fmt.Println(e)
		return
	}
	filePathes, e := fileSpliter(path + ".sec")
	if e != nil {
		fmt.Println(e)
		return
	}
	for i := 0; i < 32; i++ {
		_, e = keyFile.WriteString(filePathes[i] + "\r\n")
		if e != nil {
			fmt.Println(e)
			return
		}
	}
	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0132456789"

func randStringBytesRmndr(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}

// func fileSpliter(filepath string) (filenames [32]string, err error) {
// 	const totalPartsNum int64 = 32
// 	file, err := os.Open(filepath)
// 	defer file.Close()

// 	if err != nil {
// 		return
// 	}

// 	fileInfo, _ := file.Stat()
// 	var fileSize = fileInfo.Size()
// 	var fileChunk = fileSize / totalPartsNum

// 	for i := int64(0); i < totalPartsNum; i++ {

// 		partSize := int(math.Min(float64(fileChunk), float64(fileSize-int64(i*fileChunk))))
// 		partBuffer := make([]byte, partSize)

// 		file.Read(partBuffer)

// 		fileName := filepath + "_" + strconv.FormatUint(uint64(i), 10) + ".part"
// 		_, err := os.Create(fileName)

// 		if err != nil {
// 			return filenames, err
// 		}

// 		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
// 		filenames[i] = fileName
// 	}
// 	return
// }
