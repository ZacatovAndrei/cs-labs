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

Both `RC5` and `RC4` are classes inheriting from the base class `Cipher`. The difference in structure is prominent enough to not group them under a common `ModernCipher` abstraction, and the number of ciphers to implement doesn't call for separate abstractions, such as a `StringCipher` and/or `BlockCipher`.

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

2. RC5
 RC5 is one of the easier block ciphers to implement both in software and hardware, however it is still sufficiently secure if provided enough rounds. Here are the main functions implemented for this cipher:

```c#
private void setup(byte[] K, int b)
            {
                // magic constants

                UInt64 Pw = 0xB7E151628AED2A6B;
                UInt64 Qw = 0x9E3779B97F4A7C15;
                //things that don't unnecessary have to be uint64
                int u = _w / 8;
                int t = 2 * (_r + 1);
                int c = Math.Max(1, (int)Math.Ceiling(8 * (double)(b / u)));
                int i, j, k;

                //Things that have to be UInt64
                UInt64 A, B;
                UInt64[] L = new UInt64[c];

                // Key scheduling
                for (i = b - 1, L[c - 1] = 0; i != -1; i--)
                {
                    L[i / u] = (L[i / u] << 8) + K[i];
                }

                for (_S[0] = Pw, i = 1; i < t; i++)
                {
                    _S[i] = _S[i - 1] + Qw;
                }

                for (A = B = 0, i = j = k = 0; k < 3 * t; k++, i = (i + 1) % t, j = (j + 1) % c)
                {
                    A = _S[i] = BitOperations.RotateLeft(_S[i] + (A + B), 3);
                    B = L[j] = BitOperations.RotateLeft(L[j] + (A + B), (int)((A + B) % 64));
                }
```

The key scheduling function consists of generating a pseudortandom array S, which is private to the class, and, at first, initialised with the use of 2 Magic numbers P and Q  based on the golden ratio and the number e as the "nothing-up-my-sleeve" numbers. Then with the provided formulas and circular shifting the key is generated based on the secret key provided.

```c#
private byte[] encodeBlock(byte[] block)
            {
                UInt64 A = BitConverter.ToUInt64(block, 0) + _S[0];
                UInt64 B = BitConverter.ToUInt64(block, 8) + _S[1];
                byte[] Result = new byte[16];
                for (int i = 1; i <= _r; i++)
                {
                    A = BitOperations.RotateLeft(A ^ B,/*!!*/ (int)(B % 64)) + _S[2 * i];
                    B = BitOperations.RotateLeft(B ^ A,/*!!*/ (int)(A % 64)) + _S[2 * i + 1];
                }
                BitConverter.GetBytes(A).CopyTo(Result, 0);
                BitConverter.GetBytes(B).CopyTo(Result, 8);
                return Result;
            }
```

The encoding function looks the following way. A block here is considered to be twice the length of the word size chosen (64 in this case). The bytes are XORed and shifted accoring to the formulas for as many rounds (_r) as chosen (12 in this case).  
For the sake of simplicity when using bit shifts the byte array has been broken up into 2 words and then back to byte arrays.  
The modular operations and casting in the evidentiated areas (/*!!*/) are needed due to the method signature being `BinOperations.RotateLeft(UInt64,int)`
The Decode() method is nothing more than the reverse of the Encode() method, hence only one is provided

```c#
public override string Encode(string plain)
            {
                //adjusting the length of the ciphertext
                var bs = 16;
                if (plain.Length % bs != 0)
                {
                    var add = (bs - plain.Length % bs);
                    plain += new String('\0', add);
                }
                // Ecnoding the blocks
                var byteString = Encoding.ASCII.GetBytes(plain);
                var encodedString = new byte[byteString.Length];
                var strlen = byteString.Length;
                for (int i = 0; i < strlen; i += bs)
                {

                    encodeBlock(byteString[i..(i + bs)]).CopyTo(encodedString, i);
                }
                return Convert.ToHexString(encodedString);
            }

            public override string Decode(string cipher)
            {
                //getting bytes from the string
                var bytes = Convert.FromHexString(cipher);
                var strlen = bytes.Length;
                var bs = 16;
                var decodedString = new byte[strlen];
                for (int i = 0; i < strlen; i += bs)
                {

                    decodeBlock(bytes[i..(i + bs)]).CopyTo(decodedString, i);
                }
                var cleaning = Array.IndexOf<byte>(decodedString, 0);
                Array.Resize<byte>(ref decodedString, cleaning);
                return Encoding.ASCII.GetString(decodedString);
            }

```

Encoding and Decoding functions are quite simple. In the encoding function the message is divided into individual bytes, padded with zeroes accodrdingly to the blocksize and then each block is enctypted individually.
In the decoding function the hexString is decoded back into the sequence of bytes, then divided in blocks, each block is decoded and then the padding zeroes are removed.

### Conclusions / Screenshots / Results

By going through this laboratory work one can see the transition from simple "classical" ciphers that used methods, easily understandable for humans, to computer-related operations that are harder to compute and keep track of.  
One has seen the difference between the 2 main types of modern symmetric ciphers: the stream ciphers, that are faster, but are less secure, and block ciphers, which have a more complex structure and are slower, but with that are also more secure.
