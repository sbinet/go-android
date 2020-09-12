// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	_ "gioui.org/app/permission/camera"
	_ "gioui.org/app/permission/storage"
	"github.com/sbinet/go-android/camera"
)

func main() {
	mgr := camera.NewManager()
	defer mgr.Delete()

	ids, err := mgr.CameraIDs()
	if err != nil {
		log.Fatalf("could not get camera IDs: %+v", err)
	}

	for _, id := range ids {
		log.Printf("camera id=%v", id)
		md, err := mgr.CameraCharacteristics(id)
		if err != nil {
			log.Fatalf("could not retrieve camera-id=%v characteristics: %+v", id, err)
		}
		defer md.Delete()

		entries, err := md.Entries()
		if err != nil {
			log.Fatalf("could not retrieve camera-id=%v entries: %+v", id, err)
		}

		for _, entry := range entries {
			log.Printf("entry: %#v", entry)
		}
	}
}

func backFacingCameraID(mgr *camera.Manager) (camera.CameraID, error) {
	var (
		back camera.CameraID
		err  error
	)

	ids, err := mgr.CameraIDs()
	if err != nil {
		return back, err
	}

	for _, id := range ids {
		md, err := mgr.CameraCharacteristics(id)
		if err != nil {
			return back, err
		}
		defer md.Delete()

		entry, err := md.Entry(camera.LensFacing)
		if err != nil {
			return back, err
		}

		if entry.LensFacing() == camera.LensFacingBack {
			back = id
			return back, nil
		}
	}

	return back, fmt.Errorf("could not find a camera facing back")
}
