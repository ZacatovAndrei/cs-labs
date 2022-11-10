using System.Text;
using System.Collections;
using Ciphers;
using Hash;
Console.WriteLine("Made by Zacatov Andrei");
Console.WriteLine("CS lab work ciphers implemented:\nCaesar's cipher\tCaesar with permutations\nVigenere\t Atbash");
//TODO: add a name field to every class to allow for choosing a ciphers;

//Here come the tests for each cipher implemented:
List<Cipher> ciphersImplemented = new List<Cipher>();
ciphersImplemented.Add(new Ciphers.Atbash());
ciphersImplemented.Add(new Ciphers.Caesar("2"));
ciphersImplemented.Add(new Ciphers.CaesarP("5", "This is a test"));
ciphersImplemented.Add(new Ciphers.Vigenere("Attack"));
ciphersImplemented.Add(new Ciphers.RC4("Some random key of a decent length"));
ciphersImplemented.Add(new Ciphers.RC5("hello there i hate it here"));

Console.WriteLine($"{BitConverter.IsLittleEndian}");

Hash.SHA1 oof = new SHA1();

oof.processBlock(oof.preprocessMessage("abc"));



string stringToEncode = "Hello there";
string encodedString = "";

foreach (Cipher cipher in ciphersImplemented)
{
    encodedString = cipher.Encode(stringToEncode);
    Console.WriteLine("encoding string...\nResult:\t{0}\nDecoding:\t{1}", encodedString, cipher.Decode(encodedString));

}
