// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

//#include <media/NdkImageReader.h>
import "C"

type ImageFormat int32

const (
	FormatJPEG = ImageFormat(C.AIMAGE_FORMAT_JPEG)
)

type ImageReader struct {
	c *C.AImageReader
}

func NewImageReader(width, height int, fmt ImageFormat, maxImages int32) (*ImageReader, error) {
	var (
		r  ImageReader
		ok = C.AImageReader_new(
			C.int32_t(width), C.int32_t(height),
			C.int32_t(fmt),
			C.int32_t(maxImages),
			&r.c,
		)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}

	return &r, nil
}

func (r *ImageReader) Delete() {
	if r.c != nil {
		C.AImageReader_delete(r.c)
		r.c = nil
	}
}
