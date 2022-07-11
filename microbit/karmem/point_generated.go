package main

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierPoint = 12809409023221035873
)

type Point struct {
	Description string
	Value       float64
}

func (x *Point) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierPoint
}

func (x *Point) Reset() {
	x.Description = x.Description[:0]
	x.Value = 0
}

func (x *Point) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Point) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(32)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(24))
	__DescriptionSize := uint(1 * len(x.Description))
	__DescriptionOffset, err := writer.Alloc(__DescriptionSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__DescriptionOffset))
	writer.Write4At(offset+4+4, uint32(__DescriptionSize))
	writer.Write4At(offset+4+4+4, 1)
	__DescriptionSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Description)), __DescriptionSize, __DescriptionSize}
	writer.WriteAt(__DescriptionOffset, *(*[]byte)(unsafe.Pointer(&__DescriptionSlice)))
	__ValueOffset := offset + 16
	writer.Write8At(__ValueOffset, *(*uint64)(unsafe.Pointer(&x.Value)))

	return offset, nil
}

func (x *Point) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewPointViewer(reader, 0), reader)
}

func (x *Point) Read(viewer *PointViewer, reader *karmem.Reader) {
	x.Description = string(viewer.Description(reader))
	x.Value = viewer.Value()
}

type PointViewer struct {
	_data [32]byte
}

var _NullPointViewer = PointViewer{}

func NewPointViewer(reader *karmem.Reader, offset uint32) (v *PointViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullPointViewer
	}
	v = (*PointViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullPointViewer
	}
	return v
}

func (x *PointViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *PointViewer) Description(reader *karmem.Reader) (v []byte) {
	if 4+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *PointViewer) Value() (v float64) {
	if 16+8 > x.size() {
		return v
	}
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 16))
}
