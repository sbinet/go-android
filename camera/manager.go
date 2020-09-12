// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraDevice.h>
//#include <camera/NdkCameraManager.h>
//#include <camera/NdkCameraMetadataTags.h>
//#include <camera/NdkCameraMetadata.h>
//
//#include <stdlib.h>
//#include <string.h>
//
//static const char*
//chars_at(const char** arr, int i) {
//	return arr[i];
//}
//
//void
//onDisconnected(void* context, ACameraDevice* device) {
//}
//
//void
//onError(void* context, ACameraDevice* device, int error) {
//}
//
import "C"

import (
	"unsafe"
)

// Manager provides access to the camera service.
type Manager = C.ACameraManager

// NewManager creates a new camera manager.
func NewManager() *Manager {
	return C.ACameraManager_create()
}

func (mgr *Manager) Delete() {
	if mgr != nil {
		C.ACameraManager_delete(mgr)
	}
	mgr = nil
}

func (mgr *Manager) CameraIDs() ([]CameraID, error) {
	var (
		lst *C.ACameraIdList
		ok  = C.ACameraManager_getCameraIdList(mgr, &lst)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}
	defer C.ACameraManager_deleteCameraIdList(lst)

	ids := make([]CameraID, lst.numCameras)
	for i := range ids {
		str := C.GoString(C.chars_at(lst.cameraIds, C.int(i)))
		ids[i] = CameraID(str)
	}
	return ids, nil
}

func (mgr *Manager) CameraCharacteristics(id CameraID) (Metadata, error) {
	var (
		md  Metadata
		cid = C.CString(string(id))
		ok  = C.ACameraManager_getCameraCharacteristics(mgr, cid, &md.c)
		err = Status(ok)
	)
	defer C.free(unsafe.Pointer(cid))

	if err != StatusOk {
		return md, err
	}

	return md, nil
}

func (mgr *Manager) Open(id CameraID) (*Device, error) {

	var (
		dev *Device
		cbk = C.ACameraDevice_StateCallbacks{
			context:        nil,
			onDisconnected: (C.ACameraDevice_StateCallback)(unsafe.Pointer(C.onDisconnected)),
			onError:        (C.ACameraDevice_ErrorStateCallback)(unsafe.Pointer(C.onError)),
		}
		cid = C.CString(string(id))
		ok  = C.ACameraManager_openCamera(mgr, cid, &cbk, &dev)
		err = Status(ok)
	)
	defer C.free(unsafe.Pointer(cid))

	if err != StatusOk {
		return nil, err
	}

	return dev, nil
}
