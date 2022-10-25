#ifndef CLASSICAL_HPP
#define CLASSICAL_HPP

#include <Cipher.hpp>

class ClassicalCipher : public Cipher
{
    protected:
    std::string m_alphabet="ABCDEFGHIJKLMOPQRSTUVWXYZ";

public:
    ClassicalCipher(std::string& key): Cipher(key){};
    ClassicalCipher()=default;
    ~ClassicalCipher()=default;
};


#endif