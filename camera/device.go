// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraDevice.h>
import "C"

type CameraID string

type Device struct {
	c *C.struct_ACameraDevice
}

func (dev *Device) Close() error {
	if dev.c != nil {
		err := Status(C.ACameraDevice_close(dev.c))
		if err != StatusOk {
			return err
		}
		dev.c = nil
	}
	return nil
}

func (dev *Device) ID() CameraID {
	var (
		cid = C.ACameraDevice_getId(dev.c)
		id  = C.GoString(cid)
	)
	return CameraID(id)
}
