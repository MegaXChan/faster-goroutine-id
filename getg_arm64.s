#include "go_asm.h"
#include "textflag.h"

TEXT ·getg(SB),NOSPLIT,$0-8
    MOVD g, R0
    MOVD R0, ret+0(FP)
    RET
