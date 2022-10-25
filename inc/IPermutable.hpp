#ifndef IPERMUTABLE_HPP
#define IFPERMUABLE_HPP

#include <string>

class IPermutable
{
public:
    virtual std::string permute(std::string permutation) = 0;
};

#endif