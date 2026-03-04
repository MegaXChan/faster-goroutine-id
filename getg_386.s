#include "go_asm.h"
#include "textflag.h"

TEXT ·getg(SB),NOSPLIT,$0-4
    MOVL (TLS), AX
    MOVL AX, ret+0(FP)
    RET


