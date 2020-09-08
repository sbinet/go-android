// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraDevice.h>
//#include <camera/NdkCameraManager.h>
//#include <camera/NdkCameraMetadataTags.h>
//#include <camera/NdkCameraMetadata.h>
//
//#include <android/input.h>
//
import "C"

type Device struct {
	device *C.struct_ACameraDevice
}

type Manager struct {
	mgr *C.struct_ACameraManager
}
