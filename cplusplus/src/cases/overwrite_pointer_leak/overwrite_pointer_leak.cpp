#include "cases/leaks.h"

#include <iostream>

namespace leaks {

void overwrite_pointer_leak() {
    int* p = new int(7);
    // Bug pattern: pointer overwritten before delete, original allocation leaked.
    p = new int(9);

    std::cout << "[overwrite_pointer_leak] current value = " << *p << '\n';
    delete p;
}

} // namespace leaks
