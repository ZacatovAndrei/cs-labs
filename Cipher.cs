namespace Ciphers;
public abstract class Cipher
{
    protected string _key = "";
    abstract public string Encode(string plain);
    abstract public string Decode(string cipher);

}
