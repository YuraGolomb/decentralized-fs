package filehandler

import (
	"fmt"
	"io"
	"os"
)

type readFilResult struct {
	err           error
	isEnd         bool
	filename      string
	bytesRead     int
	lastPosition  int
	resultByteArr []byte
}

// DecodingBuffer sad
type DecodingBuffer struct {
	buffer           []byte
	filenames        [32]string
	bufferSize       int
	isFinished       bool
	countOfFiles     int
	currentFileNum   int
	lastFilePosition int
}

// CreateDecodingBuffer buf
func CreateDecodingBuffer(pathes [32]string) DecodingBuffer {
	return DecodingBuffer{
		buffer:           nil,
		filenames:        pathes,
		bufferSize:       1040,
		isFinished:       false,
		countOfFiles:     32,
		currentFileNum:   0,
		lastFilePosition: 0,
	}
}

// ReadNextBufferPart buf
func (buf *DecodingBuffer) ReadNextBufferPart() ([]byte, error) {
	currentFileName := buf.filenames[buf.currentFileNum]
	bytesToRead := buf.bufferSize - len(buf.buffer)
	result := buf.readFilePart(currentFileName, buf.lastFilePosition, bytesToRead)
	if result.err != nil {
		return nil, result.err
	}
	buf.lastFilePosition = result.lastPosition
	buf.buffer = append(buf.buffer[:], result.resultByteArr[:]...)

	if result.isEnd {
		if buf.currentFileNum == buf.countOfFiles-1 {
			b := buf.buffer
			buf.buffer = nil
			buf.isFinished = true
			return b, nil
		}
		buf.currentFileNum++
		buf.lastFilePosition = 0
	}
	if len(buf.buffer) == buf.bufferSize {
		b := buf.buffer
		buf.buffer = nil
		return b, nil
	}
	return buf.ReadNextBufferPart()
}

func (buf *DecodingBuffer) readFilePart(name string, position int, readCount int) readFilResult {

	result := readFilResult{
		err:           nil,
		isEnd:         false,
		filename:      name,
		bytesRead:     0,
		lastPosition:  position,
		resultByteArr: nil,
	}
	file, e := os.Open(name)
	defer file.Close()
	if e != nil {
		fmt.Println(e)
		result.err = e
		return result
	}
	fmt.Println("reading from ")
	fmt.Println(name)

	filePart := make([]byte, readCount)
	bytesRead, e := file.ReadAt(filePart, int64(position))
	if e != nil {
		if e != io.EOF {
			fmt.Println(e)
			result.err = e
			return result
		}
		result.isEnd = true
	}
	result.bytesRead = bytesRead
	result.lastPosition += bytesRead
	result.resultByteArr = filePart[:bytesRead]

	return result
}
