package ipfs

import (
	"fmt"
	"os/exec"
	"strings"
)

func Add(path string) (string, error) {
	// cmd := "echo %PATH%" //"ipfs add " //+ strings.Replace(path, "\\\\", "\\", -1)
	output, err := exec.Command("cmd", "/C", "ipfs", "add", path).Output()
	fmt.Println(err)
	fmt.Println(output)
	fmt.Println(strings.Fields(string(output)))
	return strings.Fields(string(output))[1], err
}

func Init() error {
	cmd := "ipfs init"
	return exec.Command(cmd).Run()
}

func Cat(key string, path string) error {
	output, err := exec.Command("cmd", "/C", "ipfs", "cat", key, ">", path).Output()
	fmt.Println(err)
	fmt.Println(output)
	return err
}

// import (
// 	"fmt"
// 	"os"

// 	core "github.com/ipfs/go-ipfs/core"

// 	"golang.org/x/net/context"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please give a peer ID as an argument")
// 		return
// 	}
// 	target, err := peer.IDB58Decode(os.Args[1])
// 	if err != nil {
// 		fmt.Print(err)
// 		// panic(err)
// 	}
// 	// config.Init()

// 	// Basic ipfsnode setup
// 	// r, err := fsrepo.Open("~/.ipfs")

// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	cfg := &core.BuildCfg{
// 		// Repo:   r,
// 		NilRepo: true,
// 		Online:  true,
// 	}

// 	nd, err := core.NewNode(ctx, cfg)
// 	nd.Pinning.
// 	if err != nil {
// 		panic(err)
// 	}

// 	// fmt.Printf("I am peer %s dialing %s\n", nd.Identity, target)

// 	// con, err := corenet.Dial(nd, target, "/app/whyrusleeping")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// io.Copy(os.Stdout, con)
// }

// // func initRepo() repo.Repo {
// // 	rootDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
// // 	repoPath := filepayj.Join(rootDir, "../repo")
// // 	if fsrepo.IsInitialized(repoPath) {
// // 		fsrepo.Init(repoPath, nil)
// // 	}
// // 	repo, err := fsrepo.Open(repoPath)
// // 	return repo
// // }
