using System.Text;
namespace Ciphers
{
    namespace Modern
    {

        public class RC4 : Cipher
        {
            private byte[] _S;
            private byte _i = 0;
            private byte _j = 0;

            public RC4(string key)
            {
                _key = key;
                _S = new byte[256];
                initialise();
            }

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

        }
    }
}