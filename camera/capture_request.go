// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCaptureRequest.h>
import "C"

type CaptureRequest = C.ACaptureRequest

func (req *CaptureRequest) Delete() {
	C.ACaptureRequest_free(req)
}

func (req *CaptureRequest) Copy() *CaptureRequest {
	return C.ACaptureRequest_copy(req)
}
