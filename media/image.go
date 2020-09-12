// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

//#include <media/NdkImageReader.h>
//
//extern
//void
//cbkImageReader(void *context, AImageReader *r);
//
import "C"

import (
	"unsafe"
)

type ImageFormat int32

const (
	FormatJPEG = ImageFormat(C.AIMAGE_FORMAT_JPEG)
)

type ImageReader struct {
	c *C.AImageReader
}

func NewImageReader(width, height int, fmt ImageFormat, maxImages int32) (ImageReader, error) {
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
		return r, err
	}

	return r, nil
}

func (r *ImageReader) Delete() {
	if r.c == nil {
		return
	}
	C.AImageReader_delete(r.c)
	r.c = nil
}

func (r *ImageReader) SetImageListener(cbk func(r ImageReader)) error {
	id := unsafe.Pointer(&cbk)
	imageCbks[id] = cbk

	var (
		lis = C.AImageReader_ImageListener{
			context:          id,
			onImageAvailable: (C.AImageReader_ImageCallback)(unsafe.Pointer(C.cbkImageReader)),
		}
		ok  = C.AImageReader_setImageListener(r.c, &lis)
		err = Status(ok)
	)
	if err != StatusOk {
		return err
	}
	return nil
}

var (
	imageCbks = make(map[unsafe.Pointer]func(r ImageReader))
)

//export goImageReaderCbk
func goImageReaderCbk(ctx unsafe.Pointer, r *C.AImageReader) {
	cbk := imageCbks[ctx]
	cbk(ImageReader{c: r})
}
