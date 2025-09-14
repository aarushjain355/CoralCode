cmake_minimum_required(VERSION 3.10)
project(CoralCompiler)

set(CMAKE_CXX_STANDARD 17)


include_directories(${CMAKE_SOURCE_DIR}/antlr4-cpp-runtime-install/include/antlr4-runtime)
link_directories(${CMAKE_SOURCE_DIR}/antlr4-cpp-runtime-install/lib)

add_executable(coral_compiler
    main.cpp
    coralLexer.cpp
    coralParser.cpp   
    coralBaseListener.cpp   
    CoralListener.cpp
)

# Link against ANTLR4 runtime

target_include_directories(coral_compiler PRIVATE
    ${CMAKE_SOURCE_DIR}/antlr4-cpp-runtime-install/include/antlr4-runtime
)

target_link_directories(coral_compiler PRIVATE
    ${CMAKE_SOURCE_DIR}/antlr4-cpp-runtime-install/lib
)

target_link_libraries(coral_compiler PRIVATE antlr4-runtime)