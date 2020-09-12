// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraWindowType.h>
import "C"

import (
	"github.com/sbinet/go-android"
)

type WindowType = android.NativeWindow
