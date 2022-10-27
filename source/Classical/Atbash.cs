using System.Text;
namespace Ciphers
{
    namespace Classical
    {
        public class Atbash : ClassicalCipher
        {
            public override string Encode(string plain)
            {
                plain=plain.ToUpper();
                StringBuilder res = new StringBuilder();
                int newindex = 0;
                foreach (var character in plain)
                {
                    if (!char.IsLetter(character))
                    {
                        res.Append(character);
                        continue;
                    }
                    newindex = _alphabet.Length - Array.IndexOf<char>(_alphabet, character) - 1;
                    res.Append(_alphabet[newindex]);
                }
                return res.ToString();
            }
            public override string Decode(string cipher)
            {
                StringBuilder res = new StringBuilder();
                int newindex = 0;
                foreach (var character in cipher)
                {
                    if (!char.IsLetter(character))
                    {
                        res.Append(character);
                        continue;
                    }
                    newindex = _alphabet.Length - Array.IndexOf<char>(_alphabet, character) - 1;
                    res.Append(_alphabet[newindex]);
                }
                return res.ToString().ToLower();
            }
        }

    }
}