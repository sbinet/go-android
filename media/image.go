// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

//#include <media/NdkImage.h>
import "C"

import (
	"reflect"
	"unsafe"
)

type ImageFormat int32

const (
	FormatRGBA8888        = ImageFormat(C.AIMAGE_FORMAT_RGBA_8888)
	FormatRGBX8888        = ImageFormat(C.AIMAGE_FORMAT_RGBX_8888)
	FormatRGB888          = ImageFormat(C.AIMAGE_FORMAT_RGB_888)
	FormatRGB565          = ImageFormat(C.AIMAGE_FORMAT_RGB_565)
	FormatRGBAFP16        = ImageFormat(C.AIMAGE_FORMAT_RGBA_FP16)
	FormatYUV420888       = ImageFormat(C.AIMAGE_FORMAT_YUV_420_888)
	FormatJPEG            = ImageFormat(C.AIMAGE_FORMAT_JPEG)
	FormatRAW16           = ImageFormat(C.AIMAGE_FORMAT_RAW16)
	FormatRAWPrivate      = ImageFormat(C.AIMAGE_FORMAT_RAW_PRIVATE)
	FormatRAW10           = ImageFormat(C.AIMAGE_FORMAT_RAW10)
	FormatRAW12           = ImageFormat(C.AIMAGE_FORMAT_RAW12)
	FormatDepth16         = ImageFormat(C.AIMAGE_FORMAT_DEPTH16)
	FormatDepthPointCloud = ImageFormat(C.AIMAGE_FORMAT_DEPTH_POINT_CLOUD)
	FormatPrivate         = ImageFormat(C.AIMAGE_FORMAT_PRIVATE)
	FormatY8              = ImageFormat(C.AIMAGE_FORMAT_Y8)
	FormatHEIC            = ImageFormat(C.AIMAGE_FORMAT_HEIC)
	FormatDepthJPEG       = ImageFormat(C.AIMAGE_FORMAT_DEPTH_JPEG)
)

type Image = C.AImage

func (img *Image) Delete() {
	C.AImage_delete(img)
}

func (img *Image) Width() (int32, error) {
	var (
		v   C.int32_t
		ok  = C.AImage_getWidth(img, &v)
		err = Status(ok)
	)
	if err != StatusOk {
		return 0, err
	}
	return int32(v), nil
}

func (img *Image) Height() (int32, error) {
	var (
		v   C.int32_t
		ok  = C.AImage_getHeight(img, &v)
		err = Status(ok)
	)
	if err != StatusOk {
		return 0, err
	}
	return int32(v), nil
}

func (img *Image) Format() (ImageFormat, error) {
	var (
		v   C.int32_t
		ok  = C.AImage_getFormat(img, &v)
		err = Status(ok)
	)
	if err != StatusOk {
		return 0, err
	}
	return ImageFormat(v), nil
}

func (img *Image) Timestamp() (int64, error) {
	var (
		v   C.int64_t
		ok  = C.AImage_getTimestamp(img, &v)
		err = Status(ok)
	)
	if err != StatusOk {
		return 0, err
	}
	return int64(v), nil
}

func (img *Image) NumberOfPlanes() (int32, error) {
	var (
		v   C.int32_t
		ok  = C.AImage_getNumberOfPlanes(img, &v)
		err = Status(ok)
	)
	if err != StatusOk {
		return 0, err
	}
	return int32(v), nil
}

func (img *Image) PlanePixelStride(plane int) (int32, error) {
	var (
		v   C.int32_t
		ok  = C.AImage_getPlanePixelStride(img, C.int(plane), &v)
		err = Status(ok)
	)
	if err != StatusOk {
		return 0, err
	}
	return int32(v), nil
}

func (img *Image) PlaneRowStride(plane int) (int32, error) {
	var (
		v   C.int32_t
		ok  = C.AImage_getPlaneRowStride(img, C.int(plane), &v)
		err = Status(ok)
	)
	if err != StatusOk {
		return 0, err
	}
	return int32(v), nil
}

func (img *Image) PlaneData(plane int) ([]byte, error) {
	var (
		n   C.int32_t
		v   *C.uint8_t
		ok  = C.AImage_getPlaneData(img, C.int(plane), &v, &n)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}
	var (
		out = make([]byte, n)
		raw []byte
		hdr = (*reflect.SliceHeader)((unsafe.Pointer(&raw)))
	)
	hdr.Len = int(n)
	hdr.Cap = int(n)
	hdr.Data = uintptr(unsafe.Pointer(v))

	copy(out, raw)
	return out, nil
}
