package messages

import (
	"encoding/json"
	"fmt"

	"github.com/YuraGolomb/decentralized_fs/server/filehandler"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Message message
type Message struct {
	Path    string
	KeyPath string
}

// HandleMessages handles messages
func HandleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	var message Message
	fmt.Println(m.Name)
	switch m.Name {
	case "init":
		// w.OpenDevTools()

		return
	case "checkFile":
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &message); err != nil {
				astilog.Error(errors.Wrap(err, "error"))
				payload = err.Error()
				return
			}
		}
		descr, err := filehandler.GetFileInfo(message.Path)
		payload = descr
		astilog.Debug("Here")
		return payload, err
	case "uploadFile":
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &message); err != nil {
				astilog.Error(errors.Wrap(err, "error"))
				payload = err.Error()
				return
			}
		}
		filehandler.EncodeFile(message.Path)
		return
	case "downloadFile":
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &message); err != nil {
				astilog.Error(errors.Wrap(err, "error"))
				payload = err.Error()
				return
			}
		}
		filehandler.DecodeFile(message.Path, message.KeyPath)
		return
	}
	return
}

// File type
// type File struct {
// 	Name string `json:"name"`
// 	Path string `json:"path"`
// 	Size int64  `json:"size"`
// }

// explore explores a path.
// If path is empty, it explores the user's home directory
// func explore(path string) (file File, err error) {

// 	astilog.Info(path)
// 	fd, err := os.Open(path)
// 	if err != nil {
// 		return
// 	}

// 	fileStat, err := fd.Stat()
// 	// fmt.Print(string(fileStat))

// 	if err != nil {
// 		return
// 	}

// 	file = File{
// 		Name: fileStat.Name(),
// 		Path: path,
// 		Size: fileStat.Size(),
// 	}

// }
