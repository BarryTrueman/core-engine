#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Define the version of the core-engine
#define CORE_ENGINE_VERSION "1.2.3"

// Define the author of the core-engine
#define CORE_ENGINE_AUTHOR "John Doe"

// Function to print the version and author of the core-engine
void print_core_engine_info() {
    printf("Core Engine Version: %s\n", CORE_ENGINE_VERSION);
    printf("Author: %s\n", CORE_ENGINE_AUTHOR);
}

int main() {
    // Print the version and author of the core-engine
    print_core_engine_info();

    // Return 0 to indicate successful execution
    return 0;
}