// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraDevice.h>
import "C"

import (
	"unsafe"
)

type TemplateKind = C.ACameraDevice_request_template

const (
	TemplatePreview        = TemplateKind(C.TEMPLATE_PREVIEW)
	TemplateStillCapture   = TemplateKind(C.TEMPLATE_STILL_CAPTURE)
	TemplateRecord         = TemplateKind(C.TEMPLATE_RECORD)
	TemplateVideoSnapshot  = TemplateKind(C.TEMPLATE_VIDEO_SNAPSHOT)
	TemplateZeroShutterLag = TemplateKind(C.TEMPLATE_ZERO_SHUTTER_LAG)
	TemplateManual         = TemplateKind(C.TEMPLATE_MANUAL)
)

type CameraID string

type Device = C.struct_ACameraDevice
type CaptureSessionOutput = C.ACaptureSessionOutput
type CaptureSessionOutputContainer = C.ACaptureSessionOutputContainer

func (dev *Device) Close() error {
	if dev == nil {
		return nil
	}
	err := Status(C.ACameraDevice_close(dev))
	if err != StatusOk {
		return err
	}
	dev = nil
	return nil
}

func (dev *Device) ID() CameraID {
	var (
		cid = C.ACameraDevice_getId(dev)
		id  = C.GoString(cid)
	)
	return CameraID(id)
}

func (dev *Device) CreateCaptureRequest(id TemplateKind) (*CaptureRequest, error) {
	var (
		req *CaptureRequest
		ok  = C.ACameraDevice_createCaptureRequest(dev, id, &req)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}
	return req, nil
}

func (dev *Device) CreateCaptureSession(outputs *CaptureSessionOutputContainer, cbks CaptureSessionCaptureCallbacks) (*CaptureSession, error) {
	// FIXME(sbinet): properly propagate callbacks
	var (
		sess *CaptureSession
		ok   = C.ACameraDevice_createCaptureSession(dev, outputs, nil, &sess)
		err  = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}
	return sess, nil
}

func NewCaptureSessionOutput(win *WindowType) (*CaptureSessionOutput, error) {
	var (
		out *CaptureSessionOutput
		ok  = C.ACaptureSessionOutput_create(
			(*C.ACameraWindowType)(unsafe.Pointer(win)),
			&out,
		)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}
	return out, nil
}

func (sess *CaptureSessionOutput) Delete() {
	C.ACaptureSessionOutput_free(sess)
}

func NewCaptureSessionOutputContainer() (*CaptureSessionOutputContainer, error) {
	var (
		out *CaptureSessionOutputContainer
		ok  = C.ACaptureSessionOutputContainer_create(&out)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}
	return out, nil
}

func (sess *CaptureSessionOutputContainer) Delete() {
	C.ACaptureSessionOutputContainer_free(sess)
}

func (sess *CaptureSessionOutputContainer) Add(out *CaptureSessionOutput) error {
	var (
		ok  = C.ACaptureSessionOutputContainer_add(sess, out)
		err = Status(ok)
	)
	if err != StatusOk {
		return err
	}
	return nil
}

func (sess *CaptureSessionOutputContainer) Remove(out *CaptureSessionOutput) error {
	var (
		ok  = C.ACaptureSessionOutputContainer_remove(sess, out)
		err = Status(ok)
	)
	if err != StatusOk {
		return err
	}
	return nil
}
