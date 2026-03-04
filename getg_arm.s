#include "go_asm.h"
#include "textflag.h"

TEXT ·getg(SB),NOSPLIT,$0-4
    MOVW g, R0
    MOVW R0, ret+0(FP)
    RET
