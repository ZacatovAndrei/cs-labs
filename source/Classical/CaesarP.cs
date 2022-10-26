using System.Text;

namespace Ciphers
{
    namespace Classical
    {

        public class CaesarP : Caesar
        {
            public CaesarP(string shift, string? permutation = null)
            {
                _key = shift;
                if (permutation is not null)
                {
                    _alphabet = PermuteAlphabet(permutation.ToUpper());
                }
            }
            private char[] PermuteAlphabet(string Permutation)
            {
                StringBuilder newalphabet = new StringBuilder();
                SortedSet<char> checkedLetters = new SortedSet<char>(_alphabet);

                foreach (var character in Permutation)
                {
                    if (checkedLetters.Remove(character)) newalphabet.Append(character);
                }
                Console.Write(newalphabet.ToString());

                foreach (var character in checkedLetters)
                {
                    newalphabet.Append(character);
                }
                return newalphabet.ToString().ToCharArray();
            }
        }
    }
}