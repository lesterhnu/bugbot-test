# Buggy C++ Leak Demo

This project contains intentionally buggy C++ code for code-review testing.

## Structure

- `src/cases/raw_new_leak/raw_new_leak.cpp`: `new[]` allocated memory is not released.
- `src/cases/overwrite_pointer_leak/overwrite_pointer_leak.cpp`: pointer is overwritten, losing original allocation.
- `src/cases/exception_path_leak/exception_path_leak.cpp`: memory leak on exception/early return path.
- `src/cases/shared_cycle_leak/shared_cycle_leak.cpp`: `shared_ptr` cyclic reference leak.
- `src/main.cpp`: calls all leak cases.

## Build and Run (CMake)

```bash
cd cplusplus
cmake -S . -B build
cmake --build build
./build/bugbot
```

## Alternative One-Command Compile (g++)

```bash
cd cplusplus
g++ -std=c++17 -Wall -Wextra -I./src \
  src/main.cpp \
  src/cases/raw_new_leak/raw_new_leak.cpp \
  src/cases/overwrite_pointer_leak/overwrite_pointer_leak.cpp \
  src/cases/exception_path_leak/exception_path_leak.cpp \
  src/cases/shared_cycle_leak/shared_cycle_leak.cpp \
  -o bugbot
./bugbot
```
