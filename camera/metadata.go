// Copyright ©2020 The go-android Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package camera

//#include <camera/NdkCameraMetadata.h>
//
//static uint32_t
//u32s_at(const uint32_t* arr, int i) {
//	return arr[i];
//}
import "C"

type Metadata struct {
	c *C.ACameraMetadata
}

func (md *Metadata) Delete() {
	if md.c != nil {
		C.ACameraMetadata_free(md.c)
	}
	md.c = nil
}

func (md *Metadata) AllTags() ([]uint32, error) {
	var (
		count C.int32_t
		ctags *C.uint32_t
		ok    = C.ACameraMetadata_getAllTags(md.c, &count, &ctags)
		err   = Status(ok)
	)
	if err != StatusOk {
		return nil, err
	}

	tags := make([]uint32, count)
	for i := range tags {
		tags[i] = uint32(C.u32s_at(ctags, C.int(i)))
	}

	return tags, nil
}

func (md *Metadata) Entries() ([]MetadataEntry, error) {
	tags, err := md.AllTags()
	if err != nil {
		return nil, err
	}
	entries := make([]MetadataEntry, 0, len(tags))
	for _, tag := range tags {
		var entry MetadataEntry
		err := Status(C.ACameraMetadata_getConstEntry(md.c, C.uint32_t(tag), &entry.c))
		if err != StatusOk {
			continue
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

type MetadataEntry struct {
	c C.ACameraMetadata_const_entry
}
