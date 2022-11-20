package modern

import (
	"encoding/binary"
	"encoding/hex"
	"math"
	"math/bits"
)

type RC5 struct {
	key       string
	wordWidth int
	blockSize int
	rounds    int
	s         []uint64
}

func NewRC5(key string, rounds int) *RC5 {
	S := make([]uint64, 2*(rounds+1))
	ret := RC5{
		key,
		64,
		4,
		rounds,
		S,
	}
	ret.scheduleKey([]byte(key))
	return &ret

}

func (R RC5) GetType() string {
	return "RC5 symmetric block Cipher"
}

func (R RC5) Encode(plain string) string {
	for
	blockRet := R.encodeBlock([]byte(plain))
	return hex.EncodeToString(blockRet)
}

func (R RC5) Decode(cipher string) string {
	a, _ := hex.DecodeString(cipher)
	blockRet := R.decodeBlock(a)
	return string(blockRet)
}

func (R RC5) scheduleKey(keyBytes []byte) {
	const (
		Pw = 0xB7E151628AED2A6B
		Qw = 0x9E3779B97F4A7C15
	)

	var (
		b = len(keyBytes)
		u = R.wordWidth / 8
		t = 2 * (R.rounds + 1)
		c = int(math.Max(math.Ceil(8*float64(b)/float64(u)), 1))
		A = uint64(0)
		B = uint64(0)
		L = make([]uint64, c)
	)

	for i := b - 1; i >= 0; i-- {
		L[i/u] = (L[i/u] << 8) + uint64(keyBytes[i])
	}
	R.s[0] = Pw
	for i := 1; i < t-1; i++ {
		R.s[i] = R.s[i-1] + Qw
	}
	i, j := 0, 0
	for k := 0; k < 3*t; k++ {
		tmp := (R.s[i] + A + B) << 3
		R.s[i] = tmp
		A = tmp
		tmp = (L[j] + A + B) << (A + B)
		L[j] = tmp
		B = tmp
		i = (i + 1) % t
		j = (j + 1) % c
	}

}

func (R RC5) encodeBlock(block []byte) []byte {
	var (
		A = binary.LittleEndian.Uint64(block[:8]) + R.s[0]
		B = binary.LittleEndian.Uint64(block[8:16]) + R.s[1]
	)
	for i := 1; i <= R.rounds; i++ {
		A = bits.RotateLeft64(A^B, int(B%64)) + R.s[2*i]
		B = bits.RotateLeft64(B^A, int(A%64)) + R.s[2*i+1]
	}
	var res []byte
	res = binary.LittleEndian.AppendUint64(res, A)
	res = binary.LittleEndian.AppendUint64(res, B)
	return res
}

func (R RC5) decodeBlock(block []byte) []byte {
	var (
		A = binary.LittleEndian.Uint64(block[:8])
		B = binary.LittleEndian.Uint64(block[8:16])
	)
	for i := R.rounds; i > 0; i-- {
		B = bits.RotateLeft64(B-R.s[2*i+1], -int(A%64)) ^ A
		A = bits.RotateLeft64(A-R.s[2*i], -int(B%64)) ^ B
	}
	var res []byte
	res = binary.LittleEndian.AppendUint64(res, A-R.s[0])
	res = binary.LittleEndian.AppendUint64(res, B-R.s[1])
	return res
}
