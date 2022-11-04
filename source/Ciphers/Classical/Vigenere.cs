using System.Text;

namespace Ciphers
{
    public class Vigenere : ClassicalCipher
    {
        private int[] keyIndices;
        public Vigenere(string keyword)
        {
            _key = keyword.ToUpper();
            if (keyword != string.Empty)
            {
                keyIndices = generateIndices();
            }
            else throw new ArgumentException("empty keywords not allowed");

        }
        public override string Encode(string plain)
        {
            //turning a string to uppercase for consistency;
            plain = plain.ToUpper();
            StringBuilder result = new StringBuilder();
            int keypos = 0;
            int newindex = 0;
            foreach (var character in plain)
            {
                if (!char.IsLetter(character))
                {
                    result.Append(character);
                    continue;
                }
                newindex = Array.IndexOf<char>(_alphabet, character) + keyIndices[keypos] + _alphabet.Length;
                newindex %= _alphabet.Length;
                result.Append(_alphabet[newindex]);
                keypos++;
                keypos %= _key.Length;
            }
            return result.ToString();
        }
        public override string Decode(string cipher)
        {
            cipher = cipher.ToUpper();
            StringBuilder result = new StringBuilder();
            int keypos = 0;
            int newindex = 0;
            foreach (var character in cipher)
            {
                if (!char.IsLetter(character))
                {
                    result.Append(character);
                    continue;
                }
                newindex = Array.IndexOf<char>(_alphabet, character) - keyIndices[keypos] + _alphabet.Length;
                newindex %= _alphabet.Length;
                result.Append(_alphabet[newindex]);
                keypos++;
                keypos %= _key.Length;
            }
            return result.ToString().ToLower();
        }

        private int[] generateIndices()
        {
            //generating alphabet char array
            var length = _key.Length;
            int[] result = new int[length];
            for (int i = 0; i < length; i++)
            {
                result[i] = Array.IndexOf<char>(_alphabet, _key[i]);
            }
            return result;
        }
    }

}