package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
    bytes := []byte(plain)
	for i, b := range bytes {
        bytes[i] = process(b, shiftKey)
    }

    return string(bytes)
}

func process(b byte, shift int) byte {
    if b >= 'a' && b <= 'z' {
        return (((b - 'a') + uint8(shift))%26) + uint8('a')
    }
    if b >= 'A' && b <= 'Z' {
        return (((b - 'A') + uint8(shift))%26) + uint8('A')
    }

    return b
}