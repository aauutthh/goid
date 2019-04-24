#include "textflag.h"
#include "go_tls.h"

GLOBL ·goidOffset(SB), $8

// func getg() interface{}
TEXT ·getg(SB), NOSPLIT, $32-16
	// get runtime.g
	get_tls(CX)
	MOVQ  g(CX), AX
	// get runtime.g type
	MOVQ $type·runtime·g(SB), BX
	
	// convert (*g) to interface{}
	MOVQ AX, 8(SP)
	MOVQ BX, 0(SP)
	CALL runtime·convT2E(SB)
	MOVQ 16(SP), AX
	MOVQ 24(SP), BX
	
	// return interface{}
	MOVQ AX, ret+0(FP)
	MOVQ BX, ret+8(FP)
	RET

TEXT ·Getg(SB), NOSPLIT, $0-8
	// get runtime.g
	get_tls(CX)
	MOVQ  g(CX), AX
	// get runtime.g type
	MOVQ AX, ret+0(FP)
	RET

TEXT ·GetGoIDAsm(SB),NOSPLIT,$0-8
	MOVQ ·goidOffset(SB), BX
	TESTQ BX, BX
	JBE nothas_GOIDOFFSET
	get_tls(CX)
	MOVQ  g(CX), AX
	MOVQ  0(AX)(BX*1), AX
	MOVQ  AX, ret+0(FP)
	RET
	
	nothas_GOIDOFFSET:
	MOVQ $-0x1, ret+0(FP)
	RET
