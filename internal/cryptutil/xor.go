// internal/cryptutil/xor.go
package cryptutil

// FixedXOR returns a XOR b (same length).
func FixedXOR(a, b []byte) ([]byte, error) { ... }

// SingleByteXOR applies key to every byte of input in.
func SingleByteXOR(in []byte, key byte) []byte { ... }

// RepeatingKeyXOR cycles key across input in.
func RepeatingKeyXOR(in, key []byte) []byte { ... }
