package monty

import "os"

const SIZE int = 8

func monty(x, y, montyPrime, modulus []uint32) []uint32 {
	var N = len(x)
	if SIZE != N {
		os.Exit(1)
	}
	var t [SIZE + 2]uint32
	var temp uint64
	for i := 0; i < N; i++ {
		var carry uint32
		for j := 0; j < N; j++ {
			temp = uint64(t[j])
			temp += uint64(x[i]) * uint64(y[j])
			temp += uint64(carry)
			t[j] = uint32(temp)
			carry = uint32(temp >> 32)
		}
		temp = uint64(t[N]) + uint64(carry)
		t[N] = uint32(temp)
		t[N+1] = uint32(temp >> 32)

		m := uint32(uint64(t[0]) * uint64(montyPrime[0]))
		temp = uint64(t[0]) + uint64(m)*uint64(modulus[0])
		carry = uint32(temp >> 32)
		for k := 1; k < N; k++ {
			temp = uint64(t[k])
			temp += uint64(m) * uint64(modulus[k])
			temp += uint64(carry)
			t[k-1] = uint32(temp)
			carry = uint32(temp >> 32)
		}
		temp = uint64(t[N]) + uint64(carry)
		t[N-1] = uint32(temp)
		t[N] = uint32(temp >> 32)
	}
	return normalize(t[:N], modulus)
}

func normalize(t, modulus []uint32) []uint32 {
	if less(modulus, t) {
		t = sub(t, modulus)
	}
	return t
}

// less checks if a is strictly less than b
func less(a, b []uint32) bool {
	for k := len(a) - 1; k > 0; k-- {
		if a[k] < b[k] {
			return true
		} else if a[k] > b[k] {
			return false
		}
	}
	return false
}

// sub substracts b from a
func sub(a, b []uint32) []uint32 {
	borrow := false
	for i := 0; i < len(a); i++ {
		tmp := uint64(a[i])
		underflow := (tmp == 0) && (b[i] > 0 || borrow)
		if borrow {
			tmp--
		}
		borrow = underflow || tmp < uint64(b[i])
		if borrow {
			tmp += 1 << 33
		}
		a[i] = uint32(tmp - uint64(b[i]))
	}
	return a
}
