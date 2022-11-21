package asymmetric

import (
	"hash"
)

type Signature struct {
	digest hash.Hash
	RSA
}
