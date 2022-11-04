// Useless class made for the simple purspose of differentiating the ciphers
// TODO: refactor if needed

namespace Ciphers;
abstract public class ClassicalCipher : Cipher
{
    protected char[] _alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".ToCharArray();
}