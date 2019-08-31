package monty

const SIZE uint32 = 1

func monty(a, b, rInv, N []uint32) []uint32 {
	var result [SIZE + 2]uint32
	var temp uint64
	for i := 0; i < len(a); i++ {
		var carry uint32
		for j := 0; j < len(a); j++ {
			temp = uint64(result[j])
			temp += uint64(a[i]) * uint64(b[j])
			temp += uint64(carry)
			result[j] = uint32(temp)
			carry = uint32(temp >> 32)
		}
		//temp = uint64(result[len(a)]) + uint64(carry)
		result[len(a)] = uint32(temp)
		result[len(a)+1] = uint32(temp >> 32)

		m := uint64(result[0]) * uint64(rInv[0]) % 32
		temp = uint64(result[0]) + m*uint64(N[0])
		carry = uint32(temp >> 32)
		for k := 1; k < len(a); k++ {
			temp = uint64(result[k])
			temp += m * uint64(N[k])
			temp += uint64(carry)
			result[k-1] = uint32(temp)
			carry = uint32(temp >> 32)
		}
		temp = uint64(result[len(a)]) + uint64(carry)
		result[len(a)-1] = uint32(temp)
		result[len(a)] = result[len(a)+1] + uint32(temp>>32)
	}
	return result[:]
}
