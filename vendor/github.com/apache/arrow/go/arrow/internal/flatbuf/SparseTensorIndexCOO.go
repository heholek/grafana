// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// ----------------------------------------------------------------------
/// EXPERIMENTAL: Data structures for sparse tensors
/// Coodinate format of sparse tensor index.
type SparseTensorIndexCOO struct {
	_tab flatbuffers.Table
}

func GetRootAsSparseTensorIndexCOO(buf []byte, offset flatbuffers.UOffsetT) *SparseTensorIndexCOO {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SparseTensorIndexCOO{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SparseTensorIndexCOO) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SparseTensorIndexCOO) Table() flatbuffers.Table {
	return rcv._tab
}

/// COO's index list are represented as a NxM matrix,
/// where N is the number of non-zero values,
/// and M is the number of dimensions of a sparse tensor.
/// indicesBuffer stores the location and size of this index matrix.
/// The type of index value is long, so the stride for the index matrix is unnecessary.
///
/// For example, let X be a 2x3x4x5 tensor, and it has the following 6 non-zero values:
///
///   X[0, 1, 2, 0] := 1
///   X[1, 1, 2, 3] := 2
///   X[0, 2, 1, 0] := 3
///   X[0, 1, 3, 0] := 4
///   X[0, 1, 2, 1] := 5
///   X[1, 2, 0, 4] := 6
///
/// In COO format, the index matrix of X is the following 4x6 matrix:
///
///   [[0, 0, 0, 0, 1, 1],
///    [1, 1, 1, 2, 1, 2],
///    [2, 2, 3, 1, 2, 0],
///    [0, 1, 0, 0, 3, 4]]
///
/// Note that the indices are sorted in lexicographical order.
func (rcv *SparseTensorIndexCOO) IndicesBuffer(obj *Buffer) *Buffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Buffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// COO's index list are represented as a NxM matrix,
/// where N is the number of non-zero values,
/// and M is the number of dimensions of a sparse tensor.
/// indicesBuffer stores the location and size of this index matrix.
/// The type of index value is long, so the stride for the index matrix is unnecessary.
///
/// For example, let X be a 2x3x4x5 tensor, and it has the following 6 non-zero values:
///
///   X[0, 1, 2, 0] := 1
///   X[1, 1, 2, 3] := 2
///   X[0, 2, 1, 0] := 3
///   X[0, 1, 3, 0] := 4
///   X[0, 1, 2, 1] := 5
///   X[1, 2, 0, 4] := 6
///
/// In COO format, the index matrix of X is the following 4x6 matrix:
///
///   [[0, 0, 0, 0, 1, 1],
///    [1, 1, 1, 2, 1, 2],
///    [2, 2, 3, 1, 2, 0],
///    [0, 1, 0, 0, 3, 4]]
///
/// Note that the indices are sorted in lexicographical order.
func SparseTensorIndexCOOStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SparseTensorIndexCOOAddIndicesBuffer(builder *flatbuffers.Builder, indicesBuffer flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(indicesBuffer), 0)
}
func SparseTensorIndexCOOEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
