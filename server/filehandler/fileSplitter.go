package filehandler

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func fileSpliter(encFilepath string, rootDir string) (filenames [32]string, err error) {
	const totalPartsNum int64 = 32
	var repoPath = filepath.Join(rootDir, "../repo")
	var encFileName = filepath.Base(encFilepath)
	file, err := os.Open(encFilepath)
	defer file.Close()

	if err != nil {
		return
	}

	fileInfo, _ := file.Stat()
	var fileSize = fileInfo.Size()
	var fileChunk = fileSize / totalPartsNum
	var lastFileChunk = (fileSize / totalPartsNum) + (fileSize % totalPartsNum)

	for i := int64(0); i < totalPartsNum; i++ {
		var partSize int64
		if partSize = fileChunk; i == totalPartsNum-1 {
			partSize = lastFileChunk
		}

		partBuffer := make([]byte, partSize)

		bytes, e := file.ReadAt(partBuffer, i*fileChunk)
		fmt.Println(bytes)
		if e != nil {
			fmt.Println(e)
		}
		fileName := filepath.Join(repoPath, encFileName+"_"+strconv.FormatUint(uint64(i), 10)+".part")
		endFile, err := os.Create(fileName)

		if err != nil {
			return filenames, err
		}
		endFile.Write(partBuffer)
		endFile.Sync()
		endFile.Close()
		// ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		filenames[i] = fileName
	}
	return
}
