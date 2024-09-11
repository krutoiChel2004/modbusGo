package modbusGo

import (
	"encoding/binary"
	"math"
)

func ByteToUint16(bytes []byte, id int) uint16 {
	bytes[0], bytes[1] = bytes[1], bytes[0]

	reg := binary.LittleEndian.Uint16(bytes)

	l := [16]uint16{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8196, 16364, 32768}

	if reg&l[id] == l[id] {
		return 1
	}

	return 0
}

func ByteToFloat32(bytes []byte) float32 {
	bs1 := [2]byte{bytes[0], bytes[1]}
	reg1 := binary.BigEndian.Uint16(bs1[:])

	bs2 := [2]byte{bytes[2], bytes[3]}
	reg2 := binary.BigEndian.Uint16(bs2[:])

	var combined uint32 = (uint32(reg2) << 16) | uint32(reg1)

	return math.Float32frombits(combined)
}

// func Float32ToByte(float float32) []byte {
// 	bits := math.Float32bits(float)
// 	bytes := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(bytes, bits)

// 	return bytes
// }
