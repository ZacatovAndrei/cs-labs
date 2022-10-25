#ifndef CIPHER_HPP
#define CIPHER_HPP
#include <string>

class Cipher
{
protected:
    std::string m_key = "";

public:
    explicit Cipher(std::string &key) : m_key(key){};
    Cipher() = default;
    ~Cipher() = default;
    virtual std::string Encode(std::string plainText) = 0;
    virtual std::string Decode(std::string cipherText) = 0;
};

#endif