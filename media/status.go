// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

//#include <media/NdkMediaError.h>
import "C"

import (
	"fmt"
)

type Status C.media_status_t

const (
	StatusOk                    = Status(C.AMEDIA_OK)
	StatusErrorInvalidParameter = Status(C.AMEDIA_ERROR_INVALID_PARAMETER)
	StatusErrorUnknown          = Status(C.AMEDIA_ERROR_UNKNOWN)
)

func (status Status) Error() string {
	switch status {
	case StatusOk:
		return "Ok"
	case StatusErrorInvalidParameter:
		return "ErrorInvalidParameter"
	case StatusErrorUnknown:
		return "ErrorUnknown"
	}
	return fmt.Sprintf("media-error=%d", status)
}
