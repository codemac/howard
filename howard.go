// howard - a goddamned window manager
// Copyright (C) 2014- Jeff Mickey

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
)

func main() {
	fmt.Println("howard - Copyright (C) 2014 Jeff Mickey")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY; for details type `show w'.")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("under certain conditions; type `show c' for details.")

	flag.Parse()
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
