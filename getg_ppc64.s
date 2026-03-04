#include "go_asm.h"
#include "textflag.h"

TEXT ·getg(SB),NOSPLIT,$0-8
    MOVD g, R3
    MOVD R3, ret+0(FP)
    RET
