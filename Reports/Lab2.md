# Laboratory work #2. Symmetric ciphers. Block And Stream ciphers

## Course: Cryptography & Security

## Author: Zacatov Andrei

----

### Theory

Stream ciphers are symmetric key ciphers where plaintext digits are combined with the output of a pseudorandom key steam generator. In this sort of ciphers each unit of plaintext is encoded with one unit of the key stream. In practice the bitwise XOR (eXclusive OR) operation is used to get the ciphertext.  
Stream ciphers heavliy rely on the keystream generation methods.

Block Ciphers, on the other hand, operate on larger blocks of data and rely not so much on key generation, but on the process of encryption itself, like, for example, using multiple confusions and diffusions.

### Objectives

* Implement a stream cipher (RC4 in this case)
* Implement a block cipher (Blowfish) (work in progress)...

### Implementation description

Both `BLowfish` and `RC4` are classes inheriting from the base class `Cipher`. The difference in structure is prominent enough to not group them under a common `ModernCipher` abstraction, and the number of ciphers to implement doesn't call for separate abstractions, such as a `StringCipher`, as well as block ciphes being different enough between themselves that the hierarchy will be unneedingly complex.

#### **Implementation proper**

1. RC4

```c#
 private void initialise()
            {
                _i = _j = 0;
                List<byte> byteKey = new List<byte>();

                //converting a string into a byte stream
                foreach (char symb in _key)
                {
                    byteKey.Add((byte)symb);
                }

                var len = byteKey.Count;
                //initialising the identity permutation
                for (int i = 0; i < 256; i++)
                {
                    _S[i] = (byte)i;
                }
                //initialising the permutation
                byte j = 0;
                for (int i = 0; i < 256; i++)
                {
                    j = (byte)(((j + _S[i] + _key[i % len])) % 256);

                    //swapping the values
                    //TODO: refactor into a method
                    var temp = _S[i];
                    _S[i] = _S[j];
                    _S[j] = temp;

                }
            }

```

The key-scheduling algorithm of the RC4 cipher, as seen in [here](https://en.wikipedia.org/wiki/RC4#Key-scheduling_algorithm_(KSA))

```c#
private byte[] getKeyStreamBytes(int n)
            {
                byte[] result = new byte[n];
                byte index = 0;
                for (var k = 0; k < n; k++)
                {
                    _i = (byte)((_i + 1) % 256);
                    _j = (byte)((_j + _S[_i]) % 256);

                    var temp = _S[_i];
                    _S[_i] = _S[_j];
                    _S[_j] = temp;

                    index = (byte)((_S[_i] + _S[_j]) % 256);
                    result[k] = _S[index];
                }
                return result;
            }
```

function for generating the cipherstream as seen [here](https://en.wikipedia.org/wiki/RC4#Pseudo-random_generation_algorithm_(PRGA))

```c#
public override string Encode(string plain)
            {
                int len = plain.Length;
                byte[] msgBytes = new byte[len];
                byte[] keyBytes = getKeyStreamBytes(len);
                byte[] cipherBytes = new byte[len];
                for (int i = 0; i < len; i++)
                {
                    msgBytes[i] = ((byte)plain[i]);
                }

                for (int i = 0; i < len; i++)
                {
                    cipherBytes[i] = (byte)((keyBytes[i] ^ msgBytes[i]) % 256);
                }
                return BitConverter.ToString(cipherBytes);
            }

```

function for encoding the palintext. String is transformed into a sequence of bytes of length N, then N bits of keystream are generated, the two byte arrays are XOR-ed together and the result is outputted as a string of hex values.

```c#
public override string Decode(string cipher)
            {
                //reinitialise the PRNG
                initialise();
                List<byte> msgBytes = new List<byte>();
                StringBuilder res = new StringBuilder();
                foreach (var hex in cipher.Split("-"))
                {
                    msgBytes.Add(byte.Parse(hex, System.Globalization.NumberStyles.HexNumber));
                }
                byte[] keyBytes = getKeyStreamBytes(msgBytes.Count);
                for (int i = 0; i < msgBytes.Count; i++)
                {
                    res.Append((char)(msgBytes[i] ^ keyBytes[i]));
                }
                return res.ToString();
            }
```

2. Blowfish  **\<Still in progress>**

Function that decodes the hex-string back into regular text. The generator is reset to the initial state so that correct key is generated. Strings get separated into N pairs of hex literals, forming one byte each. N bytes of cipher stream are generated and XOR-ed with  the ciphertext. bytes are cast into char types and outputted back as a string.

### Conclusions / Screenshots / Results

By going through this laboratory work one can see the transition from simple "classical" ciphers that used methods, easily understandable for humans, to computer-related operations that are harder to compute and keep track of.  
One has seen the difference between the 2 main types of modern symmetric ciphers: the stream ciphers, that are faster, but are less secure, and block ciphers, which have a more complex structure and are slower, but with that are also more secure.
