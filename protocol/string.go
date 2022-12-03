package protocol

// a string is prefixed with its length, and this length is a varint
func EncodeString(value string) (data []byte) {
	data = append(data, EncodeVarInt(int32(len(value)))...)
	data = append(data, []byte(value)...)

	return
}

// decode a minecraft string
func DecodeString(data []byte) (value string, endPosition int32) {
	// strings are prefixed with their length, and this length is a varint
	// so we need to decode the varint first then read next X bytes to get the string
	result, endPosition, _ := DecodeVarInt(data)

	value = string(data[endPosition : int(endPosition+result)])
	endPosition += result + 1

	return
}