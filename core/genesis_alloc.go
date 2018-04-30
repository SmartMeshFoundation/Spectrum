// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xf8\xa0\u0794:_\xba\xc6\u02915\x99\xc5\xfd\xe8\xc1c\x8d\xb5\x8d\x01\u078aH\x88\r\u0db3\xa7d\x00\x00\u07d4\xa9\xa4a\xab\xfa\u01a9\xa0X\xa6\xef!\x98\xfd'\x80\xe7\xd4\xf7\u224f\xa8#?\xb2\xc1(\x00\x00\u0794\xadL\x80\x16@e\xa3\xc3=\xd2\x01I\b\xc7V>\xff\x88\xabI\x88\r\u0db3\xa7d\x00\x00\u0794\xc2-SEj\xbd\x14\xda4u\x17\xa4\xb4~\xa2Hf\xb8\u3b88\r\u0db3\xa7d\x00\x00\xe2\x94\xf2\xf53\x82\xca\x0f\x0f&\xc7ht?\xf2\xe1z7\x94g1r\x8c\n&\xa9\x8f'&\u075fy\xc0\x00\x00"
const testnetAllocData = "\xf8i\xe2\x94A\x10\xbd\x1f\xf0\xb7?\xa1,%\x9a\xcf9\xc9P'\u007f&g\x87\x8c O\xce^>%\x02a\x10\x00\x00\x00\u252d\xb3\xea:\xd3V\x19\x92\x06\u0281{\x04\xfdf\x8c\xc5\x19m\xf2\x8c O\xce^>%\x02a\x10\x00\x00\x00\u2539K:\xa4\x16\t\xe3\xf5\x9c\xba\xff<,)\x8cl\xc4\xc5\v\x81\x8c O\xce^>%\x02a\x10\x00\x00\x00"
const rinkebyAllocData = "\xf8i\xe2\x94A\x10\xbd\x1f\xf0\xb7?\xa1,%\x9a\xcf9\xc9P'\u007f&g\x87\x8c O\xce^>%\x02a\x10\x00\x00\x00\u252d\xb3\xea:\xd3V\x19\x92\x06\u0281{\x04\xfdf\x8c\xc5\x19m\xf2\x8c O\xce^>%\x02a\x10\x00\x00\x00\u2539K:\xa4\x16\t\xe3\xf5\x9c\xba\xff<,)\x8cl\xc4\xc5\v\x81\x8c O\xce^>%\x02a\x10\x00\x00\x00"

