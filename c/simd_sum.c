// #include <immintrin.h>

// int simd_sum(int *arr, int length) {
//     __m128i sum_vec = _mm_setzero_si128();
//     int i;

//     for (i = 0; i <= length - 4; i += 4) {
//         __m128i data_vec = _mm_loadu_si128((__m128i*)&arr[i]);
//         sum_vec = _mm_add_epi32(sum_vec, data_vec);
//     }

//     int sum[4];
//     _mm_storeu_si128((__m128i*)sum, sum_vec);

//     int result = sum[0] + sum[1] + sum[2] + sum[3];

//     // Handle remaining elements
//     for (; i < length; ++i) {
//         result += arr[i];
//     }

//     return result;
// }
