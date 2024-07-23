#include "textflag.h"

// func simdSumArrayAsm(arr []float32) float32
TEXT ·simdSumArrayAsm(SB), NOSPLIT, $0-24
    MOVQ    arr+0(FP), SI       // Load address of arr into SI
    MOVQ    arr_len+8(FP), CX   // Load length of arr into CX
    XORPS   X0, X0              // Clear X0 (accumulator)
    XORQ    AX, AX              // Clear AX (loop counter)

loop:
    CMPQ    AX, CX
    JGE     done
    MOVUPS  (SI)(AX*4), X1      // Load 4 floats from arr into X1
    ADDPS   X1, X0              // Accumulate X1 into X0
    ADDQ    $4, AX              // Increment loop counter by 4
    JMP     loop

done:
    // Horizontal sum of X0
    MOVAPS  X0, X1
    SHUFPS  $0x1B, X0, X0       // Reverse the vector
    ADDPS   X1, X0              // Add reversed vector to original
    MOVAPS  X0, X1
    SHUFPS  $0x01, X0, X0       // Swap the two middle elements
    ADDPS   X1, X0              // Add again
    MOVSS   X0, ret+16(FP)      // Store the result in return variable
    RET
