using System.Text;
using System.Collections;
using Ciphers;

Console.WriteLine("Made by Zacatov Andrei");
Console.WriteLine("CS lab work ciphers implemented:\nCaesar's cipher\tCaesar with permutations\nVigenere\t Atbash");
//TODO: add a name field to every class to allow for choosing a ciphers;

//Here come the tests for each cipher implemented:
List<Cipher> ciphersImplemented = new List<Cipher>();
ciphersImplemented.Add(new Ciphers.Classical.Atbash());
ciphersImplemented.Add(new Ciphers.Classical.Caesar("13"));
ciphersImplemented.Add(new Ciphers.Classical.CaesarP("5", "This is a test"));
ciphersImplemented.Add(new Ciphers.Classical.Vigenere("Attack"));

string stringToEncode = "This is a simple test of the cipher. Capitalisation will be lost in the process of encoding and decoding";
string encodedString = "";

foreach (Cipher cipher in ciphersImplemented)
{
    encodedString = cipher.Encode(stringToEncode);
    Console.WriteLine("encoding string...\nResult:\t{0}\nDecoding:\t{1}", encodedString, cipher.Decode(encodedString));

}