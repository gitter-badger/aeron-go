/*
Copyright 2016 Stanislav Liberman

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logbuffer

import (
	"github.com/lirm/aeron-go/aeron/atomic"
	"unsafe"
)

type Claim struct {
	buffer atomic.Buffer
}

func (c *Claim) Wrap(buf *atomic.Buffer, offset, length int32) {
	atomic.BoundsCheck(offset, length, buf.Capacity())
	ptr := unsafe.Pointer(uintptr(buf.Ptr()) + uintptr(offset))
	c.buffer.Wrap(ptr, length)
}
