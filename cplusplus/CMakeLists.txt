cmake_minimum_required(VERSION 3.16)
project(BuggyCppLeaks LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

add_executable(bugbot
    src/main.cpp
    src/cases/raw_new_leak/raw_new_leak.cpp
    src/cases/overwrite_pointer_leak/overwrite_pointer_leak.cpp
    src/cases/exception_path_leak/exception_path_leak.cpp
    src/cases/shared_cycle_leak/shared_cycle_leak.cpp
)

target_include_directories(bugbot PRIVATE src)
