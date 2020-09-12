// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraMetadataTags.h>
import "C"

const (
	LensFacing = C.ACAMERA_LENS_FACING

	LensFacingFront    = C.ACAMERA_LENS_FACING_FRONT
	LensFacingBack     = C.ACAMERA_LENS_FACING_BACK
	LensFacingExternal = C.ACAMERA_LENS_FACING_EXTERNAL
)
