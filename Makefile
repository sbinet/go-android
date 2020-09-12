## Copyright ©2020 The go-android Authors. All rights reserved.
## Use of this source code is governed by a BSD-style
## license that can be found in the LICENSE file.

.PHONY: all

all: build

build:
	gogio -target android -x -minsdk 28 ./cmd/camera
