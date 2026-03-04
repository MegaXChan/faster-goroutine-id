#include "go_asm.h"
#include "textflag.h"

TEXT ·getg(SB),NOSPLIT,$0-8
    MOVV g, R4
    MOVV R4, ret+0(FP)
    RET
