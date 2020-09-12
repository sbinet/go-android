// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package android

//#cgo CFLAGS: -I/opt/android-ndk/sysroot/usr/include
//
//#cgo CFLAGS: -Werror
//#cgo LDFLAGS: -landroid
//
//#cgo arm   LDFLAGS: -L/opt/android-ndk/platforms/android-28/arch-arm/usr/lib
//#cgo arm64 LDFLAGS: -L/opt/android-ndk/platforms/android-28/arch-arm64/usr/lib
//#cgo 386   LDFLAGS: -L/opt/android-ndk/platforms/android-28/arch-x86/usr/lib
//#cgo amd64 LDFLAGS: -L/opt/android-ndk/platforms/android-28/arch-x86_64/usr/lib64
//
import "C"
