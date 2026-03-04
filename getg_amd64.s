#include "go_asm.h"
#include "textflag.h"

TEXT ·getg(SB),NOSPLIT,$0-8
  MOVQ (TLS), AX
  MOVQ AX, ret+0(FP)
  RET
