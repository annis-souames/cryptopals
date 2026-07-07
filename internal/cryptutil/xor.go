// internal/cryptutil/xor.go
package cryptutil

// FixedXOR returns a XOR b (same length).
func FixedXOR(a, b []byte) ([]byte, error) { 
	if len(a) != len(b) {
		return nil, errors.New("inputs must be of the same length")
	}
	result := make([]byte, len(a))
	for i:=0; i < len(a); i++ {
		result[i] = a[i] ^ b[i] // xor byte per byte
	}
	return result, nil
 }

// SingleByteXOR applies key to every byte of input data.
func SingleByteXOR(data []byte, key byte) []byte { 
	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
	result[i] = data[i] ^ key
	}
	return result
}

// RepeatingKeyXOR cycles key across input data.
func RepeatingKeyXOR(data, key []byte) []byte { 

	result := make([]byte, len(data))
	keyLen := len(key)
	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ key[i % keyLen]
	}

	return result
}
