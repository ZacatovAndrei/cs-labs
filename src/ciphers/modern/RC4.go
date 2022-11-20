package modern

import "encoding/hex"

type RC4 struct {
	key string
	s   []byte
	i   byte
	j   byte
}

func NewRC4(key string) *RC4 {
	S := make([]byte, 256)
	trc4 := RC4{
		key,
		S,
		0,
		0,
	}
	trc4.scheduleKey([]byte(key))
	return &trc4
}

func (R RC4) GetType() string {
	return "RC4, a symmetric stream cipher"
}

func (R RC4) Encode(plain string) string {
	msg := []byte(plain)
	msgLen := len(plain)
	keyBytes := R.getKeyStreamBytes(msgLen)
	res := make([]byte, msgLen)
	for i := 0; i < msgLen; i++ {
		res[i] = keyBytes[i] ^ msg[i]
	}
	return hex.EncodeToString(res)
}

func (R RC4) Decode(cipher string) string {
	R.scheduleKey([]byte(R.key))
	msg, ok := hex.DecodeString(cipher)
	if ok != nil {
		panic(ok)
	}
	msgLen := len(msg)
	keyBytes := R.getKeyStreamBytes(msgLen)
	res := make([]byte, msgLen)
	for i := 0; i < msgLen; i++ {
		res[i] = keyBytes[i] ^ msg[i]
	}
	return string(res)
}

func (R RC4) scheduleKey(key []byte) {
	keyLen := len(key)
	for i := 0; i < 256; i++ {
		R.s[i] = byte(i)
	}
	var j byte = 0
	for i := 0; i < 256; i++ {
		j = j + R.s[i] + key[i%keyLen]
		tmp := R.s[i]
		R.s[i] = R.s[j]
		R.s[j] = tmp
	}
}

func (R RC4) getKeyStreamBytes(n int) []byte {
	res := make([]byte, n)
	index := byte(0)
	for k := 0; k < n; k++ {
		R.i++
		R.j = R.j + R.s[R.i]

		tmp := R.s[R.i]
		R.s[R.i] = R.s[R.j]
		R.s[R.j] = tmp

		index = R.s[R.i] + R.s[R.j]
		res[k] = R.s[index]
	}
	return res
}
