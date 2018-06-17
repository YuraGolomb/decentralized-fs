package main

import (
	"encoding/json"
	"flag"
	"time"

	"github.com/YuraGolomb/decentralized_fs/server/messages"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Constants
const htmlAbout = `Ласаво просимо`
var (
	BuiltAt string
	w       *astilectron.Window
	AppName string
	debug   = flag.Bool("d", false, "enables the debug mode")
)

func main() {
	astilog.FlagInit()
	flag.Parse()
	
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDefaultPath: "resources/icon.png",
			AppIconDarwinPath:  "resources/icon.icns",
		},
		Debug:    true,
		AssetDir: AssetDir,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astilectron.PtrStr("Abut"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(w, "about", htmlAbout, func(m *bootstrap.MessageIn) {
							var ss string
							if err := json.Unmarshal(m.Payload, &ss); err != nil {
								astilog.Error(errors.Wrap(err, "split"))
								return
							}
							astilog.Infof("Window - %s!", ss)
						}); err != nil {
							astilog.Error(errors.Wrap(err, "Err"))
						}
						return
					},
				},
				{
					Label: astilectron.PtrStr("ToggleDevTools"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						w.OpenDevTools()
						return
					},
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, wsІ []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = wsІ[0]
			go func() {
				time.Sleep(4 * time.Second)
				if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
					astilog.Error(errors.Wrap(err, "sending check.out.menu event failed"))
				}
			}()
			return nil
		},
		Windows: []*bootstrap.Window{{
			MessageHandler: messages.HandleMessages,
			Homepage:       "index.html",
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#323"),
				Height:          astilectron.PtrInt(800),
				Center:          astilectron.PtrBool(true),
				Width:           astilectron.PtrInt(800),
			},
		}},
		RestoreAssets:  RestoreAssets,
		
	}); err != nil {
		astilog.Fatal("пОМИЛКА")
	}
}
