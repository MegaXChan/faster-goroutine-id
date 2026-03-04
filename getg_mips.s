#include "go_asm.h"
#include "textflag.h"

TEXT ·getg(SB),NOSPLIT,$0-4
    MOVW g, R4
    MOVW R4, ret+0(FP)
    RET
