// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package android

//#include <android/native_window.h>
import "C"

type NativeWindow = C.ANativeWindow

func (win *NativeWindow) Release() {
	if win == nil {
		return
	}
	C.ANativeWindow_release(win)
	win = nil
}

func (win *NativeWindow) Width() int {
	return int(C.ANativeWindow_getWidth(win))
}

func (win *NativeWindow) Height() int {
	return int(C.ANativeWindow_getHeight(win))
}
