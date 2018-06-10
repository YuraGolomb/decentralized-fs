package messages

import (
	"encoding/json"
	// "io/ioutil"
	"os"
	// "os/user"
	// "path/filepath"
	// "sort"
	"github.com/asticode/go-astilog"

	// "strconv"
	"github.com/YuraGolomb/decentralized_fs/server/filehandler"
	// "github.com/asticode/go-astichartjs"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// Message message
type Message struct {
	Path    string
	KeyPath string
}

// HandleMessages handles messages
func HandleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	var message Message
	switch m.Name {
	case "init":
		w.OpenDevTools()

		return
	case "uploadFile":
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &message); err != nil {
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
				payload = err.Error()
				return
			}
		}
		filehandler.DecodeFile(message.Path, message.KeyPath)
		return
	}
	return
}

// // Exploration represents the results of an exploration
// type Exploration struct {
// 	Dirs       []Dir              `json:"dirs"`
// 	Files      *astichartjs.Chart `json:"files,omitempty"`
// 	FilesCount int                `json:"files_count"`
// 	FilesSize  string             `json:"files_size"`
// 	Path       string             `json:"path"`
// }

// File type
type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
}

// explore explores a path.
// If path is empty, it explores the user's home directory
func explore(path string) (file File, err error) {
	// If no path is provided, use the user's home dir
	// if len(path) == 0 {
	// 	var u *user.User
	// 	if u, err = user.Current(); err != nil {
	// // 		return
	// 	}
	// 	path = u.HomeDir
	// }
	//
	// Read dir
	// fmt.Println(path)
	astilog.Info(path)
	fd, err := os.Open(path)
	if err != nil {
		return
	}
	// Init exploration

	// Add previous dir
	// if filepath.Dir(path) != path {
	// 	e.Dirs = append(e.Dirs, Dir{
	// 		Name: "..",
	// 		Path: filepath.Dir(path),
	// 	})
	// }
	// var fileStat os.FileInfo
	fileStat, err := fd.Stat()
	// fmt.Print(string(fileStat))

	if err != nil {
		return
	}

	file = File{
		Name: fileStat.Name(),
		Path: path,
		Size: fileStat.Size(),
	}

	// Loop through files
	// var sizes []int
	// var sizesMap = make(map[int][]string)
	// var filesSize int64
	// for _, f := range files {
	// 	if f.IsDir() {
	// 		e.Dirs = append(e.Dirs, Dir{
	// 			Name: f.Name(),
	// 			Path: filepath.Join(path, f.Name()),
	// 		})
	// 	} else {
	// 		var s = int(f.Size())
	// 		sizes = append(sizes, s)
	// 		sizesMap[s] = append(sizesMap[s], f.Name())
	// 		e.FilesCount++
	// 		filesSize += f.Size()
	// 	}
	// }

	// Prepare files size
	// if filesSize < 1e3 {
	// 	e.FilesSize = strconv.Itoa(int(filesSize)) + "b"
	// } else if filesSize < 1e6 {
	// 	e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024), 'f', 0, 64) + "kb"
	// } else if filesSize < 1e9 {
	// 	e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024*1024), 'f', 0, 64) + "Mb"
	// } else {
	// 	e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024*1024*1024), 'f', 0, 64) + "Gb"
	// }

	// Prepare files chart
	// sort.Ints(sizes)
	// if len(sizes) > 0 {
	// 	e.Files = &astichartjs.Chart{
	// 		Data: &astichartjs.Data{Datasets: []astichartjs.Dataset{{
	// 			BackgroundColor: []string{
	// 				astichartjs.ChartBackgroundColorYellow,
	// 				astichartjs.ChartBackgroundColorGreen,
	// 				astichartjs.ChartBackgroundColorRed,
	// 				astichartjs.ChartBackgroundColorBlue,
	// 				astichartjs.ChartBackgroundColorPurple,
	// 			},
	// 			BorderColor: []string{
	// 				astichartjs.ChartBorderColorYellow,
	// 				astichartjs.ChartBorderColorGreen,
	// 				astichartjs.ChartBorderColorRed,
	// 				astichartjs.ChartBorderColorBlue,
	// 				astichartjs.ChartBorderColorPurple,
	// 			},
	// 		}}},
	// 		Type: astichartjs.ChartTypePie,
	// 	}
	// 	var sizeOther int
	// 	for i := len(sizes) - 1; i >= 0; i-- {
	// 		for _, l := range sizesMap[sizes[i]] {
	// 			if len(e.Files.Data.Labels) < 4 {
	// 				e.Files.Data.Datasets[0].Data = append(e.Files.Data.Datasets[0].Data, sizes[i])
	// 				e.Files.Data.Labels = append(e.Files.Data.Labels, l)
	// 			} else {
	// 				sizeOther += sizes[i]
	// 			}
	// 		}
	// 	}
	// 	if sizeOther > 0 {
	// 		e.Files.Data.Datasets[0].Data = append(e.Files.Data.Datasets[0].Data, sizeOther)
	// 		e.Files.Data.Labels = append(e.Files.Data.Labels, "other")
	// 	}
	// }
	return
}
