# Laboratory work #1. Classical ciphers

## Course: Cryptography & Security

## Author: Zacatov Andrei

----

### Theory

This laboratory work mostly focuses on the ciphers that have been used in the past and which, in a way, might be one of the first ciphers used. They were mostly simple, like the Caesar's cipher or Atbash, and could be broken easily if the method of encryption was known, so the most of the security came from the secrecy of the algorithm. however, with the passing of time,  more cryptographically secure ciphers, like the Vigenere's started appearing.  

Now, i will give a shor summary of each cipher:

1. Atbash  
Atbash is a simple monoalphabetic substitution cipher that was originaly used to encrypt the hebrew alphabet, however it might apply to any other writing system where the symbol order is established. the cipher lacks the key, and,hence provides no cryptographical security.

2. Caesar's cipher  
Caesar's cipher is a monoalphabetic substitution cipher used by Julius Caesar. The encryption process is simple, as it substitutes a letter of the plaintext with a letter **shift** positions to the right. For the modern Latin alphabet the cipher only has a keyspace of 25 keys, giving it no cryptographical security as well, however the keyspace can be improved with a simple addition of a new step as we would see in the 3

3. Caesar's cipher with alphabet permutation  
this is a variation on a classical Caesar's cipher, however the second key string is also given that would indicate the alphabet permutation. After the permutation the Encoding and Decoding processes are identical to the Original cipher.
The keyspace is significantly larger at P(26)*26 keys;

4. Vigenere's cipher  
A more modern cipher, at least compared to the previous ones, it uses a table of all the possible alphabet shifts for a given alphabet. A key is given as a sequence of letters. Then, for each character in the plaintext the offset is based on the letter of the key. Since the key can be of an arbitrary length, one can use Vigenere for a OTP(One-Time Pad) implementation, by having the key be as long as the message.

### Objectives

* Implement 4 classical ciphers.

### Implementation description

* A base absract class Cipher is created with the following structure,

```c#
public abstract class Cipher
{
    protected string _key = "";
    abstract public string Encode(string plain);
    abstract public string Decode(string cipher);

}
```

which is inherited by another abstract class ClassicalCipher, which adds a new protected `_alphabet` field of type `char[]`. This class is mostly inserted for grouping purposes as the field `type` will be implemented at some point to differentiate between the classes during runtime.

Here one will be absle to see the specific code for Encoding/Decoding algorithm for the implemented ciphers:  

0. General specifics  

    * A `string.toUpper()` method is called in all the methods to remove the unnecessary complications of having  both cases of letters, hence all encoded text is fully uppercase and all decoded text is fully lowercased. Internaly, however, everything is primarily uppercase.
    * A `StringBuilder` type object exists in every method as a container for the result, as C# string class is immutable.
    * the following block

    ```c#
    if (!char.IsLetter(character))
    {
        cipherText.Append(character);
        continue;
    }
    ```  

    is generally used to let all of the punctuation through, however it is usually ignored in a practical application for it would be able to give more insight on the ciphertext's structure

1. Caesar's cipher  

    ```c#
    public override string Encode(string plain)
    {
        int shift = int.Parse(_key);
        foreach (char character in plain.ToCharArray())
        {
            if (!char.IsLetter(character))
            {
                cipherText.Append(character);
                continue;
            }
            var pos = Array.IndexOf(_alphabet, character);
            /*difference*/  
            cipherText.Append(
                _alphabet[(pos + shift + _alphabet.Length)% _alphabet.Length]
                );
        }
    }
    ```

    Here we  can see that the encoding method simply adds the shift to the position of the letter in the alphabet and then gets the modulo of the sum by the lenght of the alphabet.

    The only difference for the decode method is in the fact that in the decode method the evidentiated line of code looks like this

    ```c#
    cipherText.Append(_alphabet[(pos - shift + _alphabet.Length)% _alphabet.Length]);
    ```

    in both cases `+ _alphabet.Length` is added before the modulo for the consistency, as well as to guard the lookup of negative indices while decoding

2. Caesar's with permutation  
This cipher is identical to the previous one in encoding/decoding, however the following method has also been added

    ```c#
    private char[] PermuteAlphabet(string Permutation)
                {
                    SortedSet<char> checkedLetters = new SortedSet<char>(_alphabet);
                    foreach (var character in Permutation)
                    {
                        if (checkedLetters.Remove(character)) newalphabet.Append(character);
                    }
                    foreach (var character in checkedLetters)
                    {
                        newalphabet.Append(character);
                    } 
    ```

    here the algorithm adds the regular alphabet to a sorted set, adds all the letters from the given permutation to the beginning, then appending all of the letters that were unused in the alphabetical order.

3. Atbash  
A maximally simple cipher, whose functions are technically symmetrical, for the algorithm is exactly the same.

    ```c#
    public override string Encode(string plain)
    {
        plain=plain.ToUpper();
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
    }    
    ```

    the -1 is added due to the 0-based adressing.

4. Vigenere's cipher
 Here the function is a bit more complicated, yet similar to the caesar's ciphers.

    ```c#
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
    ```

The method loops through the message shifting the character by alphabetical index of the current key letter, acting like **N** caesar's ciphers with the key lenght of **N**

### Conclusions / Screenshots / Results

Concluding this laboratory work one can see and learn how   cryptography has started and developen in the ages before computers were introduced to the general public. We saw the progression from a cipher with no key, to a cipher that can be used to create virtually unbreakable ciphertext messages(granted the length of the key is sufficient).
As a result one now can use 4 classical ciphers of varying levels of cryptographical security.
