package common

func IsSigned8(val uint8) bool {
	return val&0b10000000 != 0
}

func MSByte(val uint16) uint8 {
	return uint8(val >> 8)
}

func LSByte(val uint16) uint8 {
	return uint8(val)
}

func U8Tou16(highByte, lowByte uint8) uint16 {
	return uint16(highByte)<<8 | uint16(lowByte)
}

/*
// NOTE: Meant only as reference
func AddWithCarryOut(a, b uint8) (uint8, uint8) {
	var carry uint8 = 0
	var sum uint8 = 0
	for i := range 8 {
		aBit := (a >> i) & 1
		bBit := (b >> i) & 1

		var sumBit uint8 = (aBit) ^ (bBit) ^ carry
		sum |= (sumBit << i)
		carry = (aBit & bBit) | (aBit & carry) | (bBit & carry)
	}

	return sum, carry
}
*/
