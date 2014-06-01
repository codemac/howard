package main

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
)

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		panic("No display connection can be made!")
	}

	randr.Init(X)

	screen := xproto.Setup(X).DefaultScreen(X)
	screen_info_cookie := randr.GetScreenInfo(X, screen.Root)
	screen_info, err := screen_info_cookie.Reply()
	if err != nil {
		panic("I don't know what errors it could possibly respond with")
	}

	for _, s := range screen_info.Sizes {
		fmt.Printf("%#v\n", s.Height, s.Width)
	}

}
