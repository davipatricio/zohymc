package protocol

// transforms a unsigned short (2 bytes) to an int
func DecodeUnsignedShort(data []byte) uint16 {
	return uint16(data[0])<<8 | uint16(data[1])
}
