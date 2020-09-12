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

	"github.com/sbinet/go-android"
)

type ImageFormat int32

const (
	FormatJPEG = ImageFormat(C.AIMAGE_FORMAT_JPEG)
)

type Image = C.AImage

type ImageReader = C.AImageReader

func NewImageReader(width, height int, fmt ImageFormat, maxImages int32) (*ImageReader, error) {
	var (
		r  *ImageReader
		ok = C.AImageReader_new(
			C.int32_t(width), C.int32_t(height),
			C.int32_t(fmt),
			C.int32_t(maxImages),
			&r,
		)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}

	return r, nil
}

func (r *ImageReader) Delete() {
	if r == nil {
		return
	}
	C.AImageReader_delete(r)
	r = nil
}

func (r *ImageReader) Width() int32 {
	var (
		v C.int32_t
		_ = C.AImageReader_getWidth(r, &v)
	)
	return int32(v)
}

func (r *ImageReader) Height() int32 {
	var (
		v C.int32_t
		_ = C.AImageReader_getHeight(r, &v)
	)
	return int32(v)
}

func (r *ImageReader) Format() ImageFormat {
	var (
		v C.int32_t
		_ = C.AImageReader_getFormat(r, &v)
	)
	return ImageFormat(v)
}

func (r *ImageReader) MaxImages() int32 {
	var (
		v C.int32_t
		_ = C.AImageReader_getMaxImages(r, &v)
	)
	return int32(v)
}

func (r *ImageReader) SetImageListener(cbk func(r *ImageReader)) error {
	id := unsafe.Pointer(&cbk)
	imageCbks[id] = cbk

	var (
		lis = C.AImageReader_ImageListener{
			context:          id,
			onImageAvailable: (C.AImageReader_ImageCallback)(unsafe.Pointer(C.cbkImageReader)),
		}
		ok  = C.AImageReader_setImageListener(r, &lis)
		err = Status(ok)
	)
	if err != StatusOk {
		return err
	}
	return nil
}

var (
	imageCbks = make(map[unsafe.Pointer]func(r *ImageReader))
)

//export goImageReaderCbk
func goImageReaderCbk(ctx unsafe.Pointer, r *C.AImageReader) {
	cbk := imageCbks[ctx]
	cbk(r)
}

func (r *ImageReader) Window() (*android.NativeWindow, error) {
	var (
		win *android.NativeWindow
		ok  = C.AImageReader_getWindow(r, (**C.ANativeWindow)(unsafe.Pointer(&win)))
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}
	return win, nil
}

func (r *ImageReader) AcquireNextImage() (*Image, error) {
	var (
		img *Image
		ok  = C.AImageReader_acquireNextImage(r, &img)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}

	return img, nil
}
