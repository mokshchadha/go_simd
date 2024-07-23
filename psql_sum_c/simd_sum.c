#if defined(__x86_64__) || defined(_M_X64) || defined(__i386) || defined(_M_IX86)

#include <immintrin.h>

void simd_sum(float* a, float* b, float* result, int size) {
    for (int i = 0; i < size; i += 8) {
        __m256 vec_a = _mm256_loadu_ps(&a[i]);
        __m256 vec_b = _mm256_loadu_ps(&b[i]);
        __m256 vec_result = _mm256_add_ps(vec_a, vec_b);
        _mm256_storeu_ps(&result[i], vec_result);
    }
}

#else

void simd_sum(float* a, float* b, float* result, int size) {
    for (int i = 0; i < size; i++) {
        result[i] = a[i] + b[i];
    }
}

#endif
