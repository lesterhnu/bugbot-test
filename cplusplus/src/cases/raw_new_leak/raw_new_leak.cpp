#include "cases/leaks.h"

#include <iostream>

namespace leaks {

void raw_new_leak() {
    int* nums = new int[1024];
    nums[0] = 42;
    // Bug pattern: missing delete[] nums;
    std::cout << "[raw_new_leak] first value = " << nums[0] << '\n';
}

} // namespace leaks
