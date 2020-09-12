// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraError.h>
import "C"

import (
	"fmt"
)

type Status C.camera_status_t

const (
	StatusOk                      = Status(C.ACAMERA_OK)
	StatusErrorBase               = Status(C.ACAMERA_ERROR_BASE)
	StatusErrorInvalidParameter   = Status(C.ACAMERA_ERROR_INVALID_PARAMETER)
	StatusErrorCameraDisconnected = Status(C.ACAMERA_ERROR_CAMERA_DISCONNECTED)
	StatusErrorNotEnoughMemory    = Status(C.ACAMERA_ERROR_NOT_ENOUGH_MEMORY)
)

func (status Status) Error() string {
	switch status {
	case StatusOk:
		return "Ok"
	case StatusErrorBase:
		return "ErrorBase"
	case StatusErrorInvalidParameter:
		return "ErrorInvalidParameter"
	case StatusErrorCameraDisconnected:
		return "ErrorCameraDisconnected"
	case StatusErrorNotEnoughMemory:
		return "ErrorNotEnoughMemory"
	}
	return fmt.Sprintf("camera-error=%d", status)
}
