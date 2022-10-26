using System.Text;
namespace Ciphers
{
    namespace Classical
    {
        public class Caesar : ClassicalCipher
        {
            public Caesar(string shift = "13")
            {
                _key = shift;

            }
            public override string Encode(string plain)
            {
                // Getting shift from the key field
                int shift = int.Parse(_key);

                //turning the string uppercase
                plain = plain.ToUpper();

                // Creating a stringbuilder entitiy to host the string
                StringBuilder cipherText = new StringBuilder();

                //looping through the message to get individual characters
                foreach (char character in plain.ToCharArray())
                {
                    if (!char.IsLetter(character))
                    {
                        cipherText.Append(character);
                        continue;
                    }
                    var pos = Array.IndexOf(_alphabet, character);
                    cipherText.Append(_alphabet[(pos - shift + _alphabet.Length) % _alphabet.Length]); ;
                }
                return cipherText.ToString();
            }
            public override string Decode(string cipher)
            {
                // Getting shift from the key field
                int shift = int.Parse(_key);
                //turning the string uppercase for consistency
                cipher = cipher.ToUpper();
                // Decomposing the alphabet into a char array
                // Creating a stringbuilder entitiy to host the string
                StringBuilder plainText = new StringBuilder();

                //looping through the message to get individual characters
                foreach (char character in cipher.ToCharArray())
                {
                    if (!char.IsLetter(character))
                    {
                        plainText.Append(character);
                        continue;
                    }
                    var pos = Array.IndexOf(_alphabet, character);
                    plainText.Append(_alphabet[(pos - shift + _alphabet.Length) % _alphabet.Length]);
                }
                return plainText.ToString().ToLower();
            }
        }
    }
}