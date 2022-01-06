// +build amd64 !noasm !appengine

#include "textflag.h"

TEXT ·add(SB),$88-48
        MOVQ         $0, ret0+32(FP)
        MOVQ         $0, ret0+40(FP)
        MOVQ         $0, t0-16(SP)
        MOVQ         $0, t0-8(SP)
        MOVQ         $0, t1-32(SP)
        MOVQ         $0, t1-24(SP)
block0:
        MOVOU        x+0(FP), X15
        MOVO         X15, X14
        MOVOU        y+16(FP), X13
        MOVO         X13, X12
        MOVO         X14, X11
        MOVO         X12, X10
        MOVOU        X11, t2-48(SP)
        PADDL        X10, X11
        MOVOU        X11, ret0+32(FP)
        RET

