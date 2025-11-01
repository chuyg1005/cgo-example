#include "json_processor.h"
#include <nlohmann/json.hpp>
#include <string>
#include <iostream>
#include <cstring>
#include <cstdlib>

using json = nlohmann::json;

extern "C" {

char* process_json(const char* json_str) {
    try {
        // Parse the JSON string
        json j = json::parse(json_str);
        
        // Add a timestamp field
        j["processed_at"] = std::time(nullptr);
        
        // Add a processed flag
        j["processed"] = true;
        
        // If there's a "name" field, create a greeting
        if (j.contains("name")) {
            j["greeting"] = "Hello, " + j["name"].get<std::string>() + "!";
        }
        
        // Convert back to string
        std::string result = j.dump(2); // Pretty print with 2 spaces indentation
        
        // Allocate memory for the result (will be freed by caller)
        char* result_cstr = (char*)malloc(result.length() + 1);
        std::strcpy(result_cstr, result.c_str());
        
        return result_cstr;
    } catch (const std::exception& e) {
        // Return error message as JSON
        json error = {
            {"error", "Failed to process JSON"},
            {"message", e.what()}
        };
        
        std::string result = error.dump(2);
        char* result_cstr = (char*)malloc(result.length() + 1);
        std::strcpy(result_cstr, result.c_str());
        
        return result_cstr;
    }
}

void free_json_result(char* result) {
    if (result) {
        free(result);
    }
}

} // extern "C"