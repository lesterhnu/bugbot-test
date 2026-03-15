#include "cases/leaks.h"

#include <iostream>

int main() {
    std::cout << "buggy c++ leak demo started\n";

    leaks::raw_new_leak();
    leaks::overwrite_pointer_leak();
    leaks::exception_path_leak();
    leaks::shared_cycle_leak();

    std::cout << "buggy c++ leak demo finished\n";
    return 0;
}
