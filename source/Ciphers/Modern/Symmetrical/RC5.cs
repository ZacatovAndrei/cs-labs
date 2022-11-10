using System.Text;
using System.Collections;
using System.Numerics;
namespace Ciphers
{
    class RC5 : Cipher
    {
        private int _w = 64;
        private int _blocksize;
        private int _r;
        private UInt64[] _S;

        public RC5(string key, int rounds = 12)
        {
            _key = key;
            _r = rounds;
            _blocksize = _w / 16;
            _S = new UInt64[2 * (rounds + 1)];
            setup(Encoding.ASCII.GetBytes(_key), _key.Length);
            Console.WriteLine("henlo, cipher initialised");
        }
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


        }

        private byte[] encodeBlock(byte[] block)
        {
            UInt64 A = BitConverter.ToUInt64(block, 0) + _S[0];
            UInt64 B = BitConverter.ToUInt64(block, 8) + _S[1];
            byte[] Result = new byte[16];
            for (int i = 1; i <= _r; i++)
            {
                A = BitOperations.RotateLeft(A ^ B, (int)(B % 64)) + _S[2 * i];
                B = BitOperations.RotateLeft(B ^ A, (int)(A % 64)) + _S[2 * i + 1];
            }
            BitConverter.GetBytes(A).CopyTo(Result, 0);
            BitConverter.GetBytes(B).CopyTo(Result, 8);
            return Result;
        }

        private byte[] decodeBlock(byte[] block)
        {
            UInt64 A = BitConverter.ToUInt64(block, 0);
            UInt64 B = BitConverter.ToUInt64(block, 8);
            byte[] Result = new byte[16];
            for (int i = _r; i > 0; i--)
            {
                B = BitOperations.RotateRight(B - _S[2 * i + 1], (int)(A % 64)) ^ A;
                A = BitOperations.RotateRight(A - _S[2 * i], (int)(B % 64)) ^ B;
            }
            A -= _S[0];
            B -= _S[1];
            BitConverter.GetBytes(A).CopyTo(Result, 0);
            BitConverter.GetBytes(B).CopyTo(Result, 8);
            return Result;
        }

        public override string Encode(string plain)
        {
            //adjusting the length of the ciphertext
            if (plain.Length % _blocksize != 0)
            {
                var add = (_blocksize - plain.Length % _blocksize);
                plain += new String('\0', add);
            }
            // Ecnoding the blocks
            var byteString = Encoding.ASCII.GetBytes(plain);
            var encodedString = new byte[byteString.Length];
            var strlen = byteString.Length;
            for (int i = 0; i < strlen; i += _blocksize)
            {

                encodeBlock(byteString[i..(i + _blocksize)]).CopyTo(encodedString, i);
            }
            return Convert.ToHexString(encodedString);
        }

        public override string Decode(string cipher)
        {
            //getting bytes from the string
            var bytes = Convert.FromHexString(cipher);
            var strlen = bytes.Length;
            var decodedString = new byte[strlen];
            for (int i = 0; i < strlen; i += _blocksize)
            {

                decodeBlock(bytes[i..(i + _blocksize)]).CopyTo(decodedString, i);
            }
            var cleaning = Array.IndexOf<byte>(decodedString, 0);
            Array.Resize<byte>(ref decodedString, cleaning);
            return Encoding.ASCII.GetString(decodedString);
        }
    }
}