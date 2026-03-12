#ifndef LEAKS_H
#define LEAKS_H

namespace leaks {
void raw_new_leak();
void overwrite_pointer_leak();
void exception_path_leak();
void shared_cycle_leak();
} // namespace leaks

#endif
