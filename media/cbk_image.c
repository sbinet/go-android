// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <media/NdkImageReader.h>

extern
void
goImageReaderCbk(void *context, AImageReader *r);

void
cbkImageReader(void *context, AImageReader *r) {
	goImageReaderCbk(context, r);
}

