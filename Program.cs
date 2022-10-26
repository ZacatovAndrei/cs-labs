using System.Text;
using Ciphers;

Console.WriteLine("Hello World!");
Console.BackgroundColor=ConsoleColor.Blue;
Console.Clear();
ClassicalCipher testcase = new Ciphers.Classical.Playfair("test");
//Console.WriteLine(testcase.Decode(testcase.Encode("")));