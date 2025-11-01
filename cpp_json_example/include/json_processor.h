#ifndef JSON_PROCESSOR_H
#define JSON_PROCESSOR_H

#ifdef __cplusplus
extern "C" {
#endif

// Process a JSON string and return a modified version
char* process_json(const char* json_str);

// Free memory allocated by process_json
void free_json_result(char* result);

#ifdef __cplusplus
}
#endif

#endif // JSON_PROCESSOR_H