#include "cases/leaks.h"

#include <iostream>
#include <stdexcept>

namespace leaks {

void exception_path_leak() {
    int* data = new int[256];
    data[0] = 100;

    try {
        throw std::runtime_error("simulated failure");
    } catch (const std::exception& ex) {
        // Bug pattern: exception path exits without delete[] data.
        std::cout << "[exception_path_leak] caught: " << ex.what() << '\n';
        return;
    }
}

} // namespace leaks
