package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x int) int {
	return int(pc[byte(x>>0*8)] +
		pc[byte(x>>1*8)] +
		pc[byte(x>>2*8)] +
		pc[byte(x>>3*8)] +
		pc[byte(x>>4*8)] +
		pc[byte(x>>5*8)] +
		pc[byte(x>>6*8)] +
		pc[byte(x>>7*8)])
}
func CountOneByTableCirculation(x int) int {
	total := 0
	for i := 0; i < 8; i++ {
		total += int(pc[byte(x>>i*8)])
	}
	return total
}

func CountOneByShift(x int) int {
	total := 0
	for x != 0 {
		total += x & 1
		x >>= 1
	}
	return total
}

func CountOneByHamming(x int) int {
	total := 0
	for x != 0 {
		total++
		x &= x - 1
	}
	return total
}
