// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraCaptureSession.h>
import "C"

type CaptureSession = C.ACameraCaptureSession

func (sess *CaptureSession) Close() {
	C.ACameraCaptureSession_close(sess)
}

func (sess *CaptureSession) Device() (*Device, error) {
	var (
		dev *Device
		ok  = C.ACameraCaptureSession_getDevice(sess, &dev)
		err = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}
	return dev, nil
}

func (sess *CaptureSession) AbortCaptures() error {
	var (
		ok  = C.ACameraCaptureSession_abortCaptures(sess)
		err = Status(ok)
	)
	if err != StatusOk {
		return err
	}
	return nil
}

type CaptureSessionStateCallbacks struct {
	OnClose  func(sess *CaptureSession)
	OnReady  func(sess *CaptureSession)
	OnActive func(sess *CaptureSession)
}

type CaptureFailure = C.ACameraCaptureFailure
