"".Add STEXT nosplit size=15 args=0x10 locals=0x0
    0x0000 00000 (method_cases_1.go:4)  TEXT    "".Add(SB), NOSPLIT|ABIInternal, $0-16
    0x0000 00000 (method_cases_1.go:4)  PCDATA  $0, $-2
    0x0000 00000 (method_cases_1.go:4)  PCDATA  $1, $-2
    0x0000 00000 (method_cases_1.go:4)  FUNCDATA    $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x0000 00000 (method_cases_1.go:4)  FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x0000 00000 (method_cases_1.go:4)  FUNCDATA    $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x0000 00000 (method_cases_1.go:5)  PCDATA  $0, $0
    0x0000 00000 (method_cases_1.go:5)  PCDATA  $1, $0
    0x0000 00000 (method_cases_1.go:5)  MOVL    "".b+12(SP), AX
    0x0004 00004 (method_cases_1.go:5)  MOVL    "".a+8(SP), CX
    0x0008 00008 (method_cases_1.go:5)  ADDL    CX, AX
    0x000a 00010 (method_cases_1.go:5)  MOVL    AX, "".~r2+16(SP)
    0x000e 00014 (method_cases_1.go:5)  RET
    0x0000 8b 44 24 0c 8b 4c 24 08 01 c8 89 44 24 10 c3     .D$..L$....D$..
"".Calculator.Value STEXT nosplit size=9 args=0x10 locals=0x0
    0x0000 00000 (method_cases_1.go:13) TEXT    "".Calculator.Value(SB), NOSPLIT|ABIInternal, $0-16
    0x0000 00000 (method_cases_1.go:13) PCDATA  $0, $-2
    0x0000 00000 (method_cases_1.go:13) PCDATA  $1, $-2
    0x0000 00000 (method_cases_1.go:13) FUNCDATA    $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x0000 00000 (method_cases_1.go:13) FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x0000 00000 (method_cases_1.go:13) FUNCDATA    $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x0000 00000 (method_cases_1.go:14) PCDATA  $0, $0
    0x0000 00000 (method_cases_1.go:14) PCDATA  $1, $0
    0x0000 00000 (method_cases_1.go:14) MOVL    "".c+8(SP), AX
    0x0004 00004 (method_cases_1.go:14) MOVL    AX, "".~r0+16(SP)
    0x0008 00008 (method_cases_1.go:14) RET
    0x0000 8b 44 24 08 89 44 24 10 c3                       .D$..D$..
"".(*Calculator).SetValue STEXT nosplit size=12 args=0x10 locals=0x0
    0x0000 00000 (method_cases_1.go:18) TEXT    "".(*Calculator).SetValue(SB), NOSPLIT|ABIInternal, $0-16
    0x0000 00000 (method_cases_1.go:18) PCDATA  $0, $-2
    0x0000 00000 (method_cases_1.go:18) PCDATA  $1, $-2
    0x0000 00000 (method_cases_1.go:18) FUNCDATA    $0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
    0x0000 00000 (method_cases_1.go:18) FUNCDATA    $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
    0x0000 00000 (method_cases_1.go:18) FUNCDATA    $2, gclocals·568470801006e5c0dc3947ea998fe279(SB)
    0x0000 00000 (method_cases_1.go:19) PCDATA  $0, $0
    0x0000 00000 (method_cases_1.go:19) PCDATA  $1, $0
    0x0000 00000 (method_cases_1.go:19) MOVL    "".val+16(SP), AX
    0x0004 00004 (method_cases_1.go:19) PCDATA  $0, $1
    0x0004 00004 (method_cases_1.go:19) PCDATA  $1, $1
    0x0004 00004 (method_cases_1.go:19) MOVQ    "".c+8(SP), CX
    0x0009 00009 (method_cases_1.go:19) PCDATA  $0, $0
    0x0009 00009 (method_cases_1.go:19) MOVL    AX, (CX)
    0x000b 00011 (method_cases_1.go:20) RET
    0x0000 8b 44 24 10 48 8b 4c 24 08 89 01 c3              .D$.H.L$....
"".(*Calculator).Add STEXT nosplit size=12 args=0x10 locals=0x0
    0x0000 00000 (method_cases_1.go:23) TEXT    "".(*Calculator).Add(SB), NOSPLIT|ABIInternal, $0-16
    0x0000 00000 (method_cases_1.go:23) PCDATA  $0, $-2
    0x0000 00000 (method_cases_1.go:23) PCDATA  $1, $-2
    0x0000 00000 (method_cases_1.go:23) FUNCDATA    $0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
    0x0000 00000 (method_cases_1.go:23) FUNCDATA    $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
    0x0000 00000 (method_cases_1.go:23) FUNCDATA    $2, gclocals·568470801006e5c0dc3947ea998fe279(SB)
    0x0000 00000 (method_cases_1.go:24) PCDATA  $0, $0
    0x0000 00000 (method_cases_1.go:24) PCDATA  $1, $0
    0x0000 00000 (method_cases_1.go:24) MOVL    "".incr+16(SP), AX
    0x0004 00004 (method_cases_1.go:24) PCDATA  $0, $1
    0x0004 00004 (method_cases_1.go:24) PCDATA  $1, $1
    0x0004 00004 (method_cases_1.go:24) MOVQ    "".c+8(SP), CX
    0x0009 00009 (method_cases_1.go:24) PCDATA  $0, $0
    0x0009 00009 (method_cases_1.go:24) ADDL    AX, (CX)
    0x000b 00011 (method_cases_1.go:25) RET
    0x0000 8b 44 24 10 48 8b 4c 24 08 01 01 c3              .D$.H.L$....
"".main STEXT size=105 args=0x0 locals=0x20
    0x0000 00000 (method_cases_1.go:27) TEXT    "".main(SB), ABIInternal, $32-0
    0x0000 00000 (method_cases_1.go:27) MOVQ    (TLS), CX
    0x0009 00009 (method_cases_1.go:27) CMPQ    SP, 16(CX)
    0x000d 00013 (method_cases_1.go:27) PCDATA  $0, $-2
    0x000d 00013 (method_cases_1.go:27) JLS 98
    0x000f 00015 (method_cases_1.go:27) PCDATA  $0, $-1
    0x000f 00015 (method_cases_1.go:27) SUBQ    $32, SP
    0x0013 00019 (method_cases_1.go:27) MOVQ    BP, 24(SP)
    0x0018 00024 (method_cases_1.go:27) LEAQ    24(SP), BP
    0x001d 00029 (method_cases_1.go:27) PCDATA  $0, $-2
    0x001d 00029 (method_cases_1.go:27) PCDATA  $1, $-2
    0x001d 00029 (method_cases_1.go:27) FUNCDATA    $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x001d 00029 (method_cases_1.go:27) FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x001d 00029 (method_cases_1.go:27) FUNCDATA    $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    0x001d 00029 (method_cases_1.go:28) PCDATA  $0, $0
    0x001d 00029 (method_cases_1.go:28) PCDATA  $1, $0

    0x001d 00029 (method_cases_1.go:28) MOVQ    $90194313226, AX
    0x0027 00039 (method_cases_1.go:28) MOVQ    AX, (SP)
    0x002b 00043 (method_cases_1.go:28) CALL    "".Add(SB)

    ; 这里没有创建 Calculator 结构，相反，其在栈上创建一个新的和 cal 变量相等的值
    0x0030 00048 (method_cases_1.go:31) MOVL    $0, "".cal+20(SP)
    0x0038 00056 (method_cases_1.go:31) MOVL    $1, "".cal+20(SP)
    ;; 这里编译器做一些优化，生成的汇编没有什么地方有对 type."".Calculator(SB) 的引用，尽管这个地址是我们 receiver 所在的地址位置。
    ;; 这是因为 receiver 是值类型，且编译器能够通过静态分析推测出其值，这种情况下编译器认为不需要对值从它原来的位置进行拷贝了。 
    ;; 相应的，只要简单的在栈上创建一个新的和 cal 相等的值即可。

    ;; 直接新建了 cal 变量的拷贝到栈上，以调用方法 main.Calculator.Value
    0x0040 00064 (method_cases_1.go:32) MOVL    $1, (SP)
    0x0047 00071 (method_cases_1.go:32) CALL    "".Calculator.Value(SB)

    ;; 这里没有新建拷贝，但其直接拷贝值类型的 receiver 变量，来调用方法 main.Calculator.Value,
    ;; 而未调用方法 main.(*Calculator).Value
    0x004c 00076 (method_cases_1.go:33) MOVL    "".cal+20(SP), AX
    0x0050 00080 (method_cases_1.go:33) MOVL    AX, (SP)
    0x0053 00083 (method_cases_1.go:33) CALL    "".Calculator.Value(SB)


    0x0058 00088 (method_cases_1.go:34) MOVQ    24(SP), BP
    0x005d 00093 (method_cases_1.go:34) ADDQ    $32, SP
    0x0061 00097 (method_cases_1.go:34) RET
    0x0062 00098 (method_cases_1.go:34) NOP
    0x0062 00098 (method_cases_1.go:27) PCDATA  $1, $-1
    0x0062 00098 (method_cases_1.go:27) PCDATA  $0, $-2
    0x0062 00098 (method_cases_1.go:27) CALL    runtime.morestack_noctxt(SB)
    0x0067 00103 (method_cases_1.go:27) PCDATA  $0, $-1
    0x0067 00103 (method_cases_1.go:27) JMP 0
    0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 53 48  dH..%....H;a.vSH
    0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 b8 0a  .. H.l$.H.l$.H..
    0x0020 00 00 00 15 00 00 00 48 89 04 24 e8 00 00 00 00  .......H..$.....
    0x0030 c7 44 24 14 00 00 00 00 c7 44 24 14 00 00 00 00  .D$......D$.....
    0x0040 c7 04 24 00 00 00 00 e8 00 00 00 00 8b 44 24 14  ..$..........D$.
    0x0050 89 04 24 e8 00 00 00 00 48 8b 6c 24 18 48 83 c4  ..$.....H.l$.H..
    0x0060 20 c3 e8 00 00 00 00 eb 97                        ........
    rel 5+4 t=17 TLS+0
    rel 44+4 t=8 "".Add+0
    rel 72+4 t=8 "".Calculator.Value+0
    rel 84+4 t=8 "".Calculator.Value+0
    rel 99+4 t=8 runtime.morestack_noctxt+0
"".(*Calculator).Value STEXT dupok size=104 args=0x10 locals=0x18
    0x0000 00000 (<autogenerated>:1)    TEXT    "".(*Calculator).Value(SB), DUPOK|WRAPPER|ABIInternal, $24-16
    0x0000 00000 (<autogenerated>:1)    MOVQ    (TLS), CX
    0x0009 00009 (<autogenerated>:1)    CMPQ    SP, 16(CX)
    0x000d 00013 (<autogenerated>:1)    PCDATA  $0, $-2
    0x000d 00013 (<autogenerated>:1)    JLS 82
    0x000f 00015 (<autogenerated>:1)    PCDATA  $0, $-1
    0x000f 00015 (<autogenerated>:1)    SUBQ    $24, SP
    0x0013 00019 (<autogenerated>:1)    MOVQ    BP, 16(SP)
    0x0018 00024 (<autogenerated>:1)    LEAQ    16(SP), BP
    0x001d 00029 (<autogenerated>:1)    MOVQ    32(CX), BX
    0x0021 00033 (<autogenerated>:1)    TESTQ   BX, BX
    0x0024 00036 (<autogenerated>:1)    JNE 89
    0x0026 00038 (<autogenerated>:1)    NOP
    0x0026 00038 (<autogenerated>:1)    PCDATA  $0, $-2
    0x0026 00038 (<autogenerated>:1)    PCDATA  $1, $-2
    0x0026 00038 (<autogenerated>:1)    FUNCDATA    $0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
    0x0026 00038 (<autogenerated>:1)    FUNCDATA    $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
    0x0026 00038 (<autogenerated>:1)    FUNCDATA    $2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
    0x0026 00038 (<autogenerated>:1)    PCDATA  $0, $1
    0x0026 00038 (<autogenerated>:1)    PCDATA  $1, $1
    0x0026 00038 (<autogenerated>:1)    MOVQ    ""..this+32(SP), AX
    0x002b 00043 (<autogenerated>:1)    TESTQ   AX, AX
    0x002e 00046 (<autogenerated>:1)    JEQ 76
    0x0030 00048 (<autogenerated>:1)    PCDATA  $0, $0
    0x0030 00048 (<autogenerated>:1)    MOVL    (AX), AX
    0x0032 00050 (<autogenerated>:1)    MOVL    AX, (SP)
    0x0035 00053 (<autogenerated>:1)    CALL    "".Calculator.Value(SB)
    0x003a 00058 (<autogenerated>:1)    MOVL    8(SP), AX
    0x003e 00062 (<autogenerated>:1)    MOVL    AX, "".~r0+40(SP)
    0x0042 00066 (<autogenerated>:1)    MOVQ    16(SP), BP
    0x0047 00071 (<autogenerated>:1)    ADDQ    $24, SP
    0x004b 00075 (<autogenerated>:1)    RET
    0x004c 00076 (<autogenerated>:1)    CALL    runtime.panicwrap(SB)
    0x0051 00081 (<autogenerated>:1)    XCHGL   AX, AX
    0x0052 00082 (<autogenerated>:1)    NOP
    0x0052 00082 (<autogenerated>:1)    PCDATA  $1, $-1
    0x0052 00082 (<autogenerated>:1)    PCDATA  $0, $-2
    0x0052 00082 (<autogenerated>:1)    CALL    runtime.morestack_noctxt(SB)
    0x0057 00087 (<autogenerated>:1)    PCDATA  $0, $-1
    0x0057 00087 (<autogenerated>:1)    JMP 0
    0x0059 00089 (<autogenerated>:1)    LEAQ    32(SP), DI
    0x005e 00094 (<autogenerated>:1)    CMPQ    (BX), DI
    0x0061 00097 (<autogenerated>:1)    JNE 38
    0x0063 00099 (<autogenerated>:1)    MOVQ    SP, (BX)
    0x0066 00102 (<autogenerated>:1)    JMP 38
    0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 43 48  dH..%....H;a.vCH
    0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 8b 59  ...H.l$.H.l$.H.Y
    0x0020 20 48 85 db 75 33 48 8b 44 24 20 48 85 c0 74 1c   H..u3H.D$ H..t.
    0x0030 8b 00 89 04 24 e8 00 00 00 00 8b 44 24 08 89 44  ....$......D$..D
    0x0040 24 28 48 8b 6c 24 10 48 83 c4 18 c3 e8 00 00 00  $(H.l$.H........
    0x0050 00 90 e8 00 00 00 00 eb a7 48 8d 7c 24 20 48 39  .........H.|$ H9
    0x0060 3b 75 c3 48 89 23 eb be                          ;u.H.#..
    rel 5+4 t=17 TLS+0
    rel 54+4 t=8 "".Calculator.Value+0
    rel 77+4 t=8 runtime.panicwrap+0
    rel 83+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
    0x0000 6d 61 69 6e                                      main
go.loc."".Add SDWARFLOC size=103
    0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
    0x0010 00 00 00 00 00 00 00 00 0f 00 00 00 00 00 00 00  ................
    0x0020 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
    0x0040 00 00 00 00 00 00 00 00 00 00 00 0f 00 00 00 00  ................
    0x0050 00 00 00 02 00 91 04 00 00 00 00 00 00 00 00 00  ................
    0x0060 00 00 00 00 00 00 00                             .......
    rel 8+8 t=1 "".Add+0
    rel 59+8 t=1 "".Add+0
go.info."".Add SDWARFINFO size=70
    0x0000 03 22 22 2e 41 64 64 00 00 00 00 00 00 00 00 00  ."".Add.........
    0x0010 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 10  ................
    0x0020 61 00 00 04 00 00 00 00 00 00 00 00 10 62 00 00  a............b..
    0x0030 04 00 00 00 00 00 00 00 00 0f 7e 72 32 00 01 04  ..........~r2...
    0x0040 00 00 00 00 00 00                                ......
    rel 0+0 t=24 type.int32+0
    rel 8+8 t=1 "".Add+0
    rel 16+8 t=1 "".Add+15
    rel 26+4 t=30 gofile../home/ubuntu/workSpaces/go/src/github.com/qqzeng/go-code-snippet/interface/method_cases_1.go+0
    rel 36+4 t=29 go.info.int32+0
    rel 40+4 t=29 go.loc."".Add+0
    rel 49+4 t=29 go.info.int32+0
    rel 53+4 t=29 go.loc."".Add+51
    rel 64+4 t=29 go.info.int32+0
go.range."".Add SDWARFRANGE size=0
go.debuglines."".Add SDWARFMISC size=11
    0x0000 04 02 13 06 37 04 01 03 7c 06 01                 ....7...|..
go.loc."".Calculator.Value SDWARFLOC size=51
    0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
    0x0010 00 00 00 00 00 00 00 00 09 00 00 00 00 00 00 00  ................
    0x0020 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00                                         ...
    rel 8+8 t=1 "".Calculator.Value+0
go.info."".Calculator.Value SDWARFINFO size=70
    0x0000 03 22 22 2e 43 61 6c 63 75 6c 61 74 6f 72 2e 56  ."".Calculator.V
    0x0010 61 6c 75 65 00 00 00 00 00 00 00 00 00 00 00 00  alue............
    0x0020 00 00 00 00 00 01 9c 00 00 00 00 01 10 63 00 00  .............c..
    0x0030 0d 00 00 00 00 00 00 00 00 0f 7e 72 30 00 01 0d  ..........~r0...
    0x0040 00 00 00 00 00 00                                ......
    rel 0+0 t=24 type."".Calculator+0
    rel 0+0 t=24 type.int32+0
    rel 21+8 t=1 "".Calculator.Value+0
    rel 29+8 t=1 "".Calculator.Value+9
    rel 39+4 t=30 gofile../home/ubuntu/workSpaces/go/src/github.com/qqzeng/go-code-snippet/interface/method_cases_1.go+0
    rel 49+4 t=29 go.info."".Calculator+0
    rel 53+4 t=29 go.loc."".Calculator.Value+0
    rel 64+4 t=29 go.info.int32+0
go.range."".Calculator.Value SDWARFRANGE size=0
go.debuglines."".Calculator.Value SDWARFMISC size=13
    0x0000 04 02 03 08 14 06 37 04 01 03 73 06 01           ......7...s..
go.loc."".(*Calculator).SetValue SDWARFLOC size=103
    0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
    0x0010 00 00 00 00 00 00 00 00 0c 00 00 00 00 00 00 00  ................
    0x0020 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
    0x0040 00 00 00 00 00 00 00 00 00 00 00 0c 00 00 00 00  ................
    0x0050 00 00 00 02 00 91 08 00 00 00 00 00 00 00 00 00  ................
    0x0060 00 00 00 00 00 00 00                             .......
    rel 8+8 t=1 "".(*Calculator).SetValue+0
    rel 59+8 t=1 "".(*Calculator).SetValue+0
go.info."".(*Calculator).SetValue SDWARFINFO size=79
    0x0000 03 22 22 2e 28 2a 43 61 6c 63 75 6c 61 74 6f 72  ."".(*Calculator
    0x0010 29 2e 53 65 74 56 61 6c 75 65 00 00 00 00 00 00  ).SetValue......
    0x0020 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
    0x0030 00 01 10 63 00 00 12 00 00 00 00 00 00 00 00 10  ...c............
    0x0040 76 61 6c 00 00 12 00 00 00 00 00 00 00 00 00     val............
    rel 0+0 t=24 type.*"".Calculator+0
    rel 0+0 t=24 type.int32+0
    rel 27+8 t=1 "".(*Calculator).SetValue+0
    rel 35+8 t=1 "".(*Calculator).SetValue+12
    rel 45+4 t=30 gofile../home/ubuntu/workSpaces/go/src/github.com/qqzeng/go-code-snippet/interface/method_cases_1.go+0
    rel 55+4 t=29 go.info.*"".Calculator+0
    rel 59+4 t=29 go.loc."".(*Calculator).SetValue+0
    rel 70+4 t=29 go.info.int32+0
    rel 74+4 t=29 go.loc."".(*Calculator).SetValue+51
go.range."".(*Calculator).SetValue SDWARFRANGE size=0
go.debuglines."".(*Calculator).SetValue SDWARFMISC size=14
    0x0000 04 02 03 0d 14 06 37 06 56 04 01 03 6d 01        ......7.V...m.
go.loc."".(*Calculator).Add SDWARFLOC size=103
    0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
    0x0010 00 00 00 00 00 00 00 00 0c 00 00 00 00 00 00 00  ................
    0x0020 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
    0x0040 00 00 00 00 00 00 00 00 00 00 00 0c 00 00 00 00  ................
    0x0050 00 00 00 02 00 91 08 00 00 00 00 00 00 00 00 00  ................
    0x0060 00 00 00 00 00 00 00                             .......
    rel 8+8 t=1 "".(*Calculator).Add+0
    rel 59+8 t=1 "".(*Calculator).Add+0
go.info."".(*Calculator).Add SDWARFINFO size=75
    0x0000 03 22 22 2e 28 2a 43 61 6c 63 75 6c 61 74 6f 72  ."".(*Calculator
    0x0010 29 2e 41 64 64 00 00 00 00 00 00 00 00 00 00 00  ).Add...........
    0x0020 00 00 00 00 00 00 01 9c 00 00 00 00 01 10 63 00  ..............c.
    0x0030 00 17 00 00 00 00 00 00 00 00 10 69 6e 63 72 00  ...........incr.
    0x0040 00 17 00 00 00 00 00 00 00 00 00                 ...........
    rel 0+0 t=24 type.*"".Calculator+0
    rel 0+0 t=24 type.int32+0
    rel 22+8 t=1 "".(*Calculator).Add+0
    rel 30+8 t=1 "".(*Calculator).Add+12
    rel 40+4 t=30 gofile../home/ubuntu/workSpaces/go/src/github.com/qqzeng/go-code-snippet/interface/method_cases_1.go+0
    rel 50+4 t=29 go.info.*"".Calculator+0
    rel 54+4 t=29 go.loc."".(*Calculator).Add+0
    rel 66+4 t=29 go.info.int32+0
    rel 70+4 t=29 go.loc."".(*Calculator).Add+51
go.range."".(*Calculator).Add SDWARFRANGE size=0
go.debuglines."".(*Calculator).Add SDWARFMISC size=14
    0x0000 04 02 03 12 14 06 37 06 56 04 01 03 68 01        ......7.V...h.
go.loc."".main SDWARFLOC size=52
    0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
    0x0010 38 00 00 00 00 00 00 00 69 00 00 00 00 00 00 00  8.......i.......
    0x0020 02 00 91 6c 00 00 00 00 00 00 00 00 00 00 00 00  ...l............
    0x0030 00 00 00 00                                      ....
    rel 8+8 t=1 "".main+0
go.info."".main SDWARFINFO size=47
    0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
    0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
    0x0020 0b 63 61 6c 00 1f 00 00 00 00 00 00 00 00 00     .cal...........
    rel 0+0 t=24 type."".Calculator+0
    rel 9+8 t=1 "".main+0
    rel 17+8 t=1 "".main+105
    rel 27+4 t=30 gofile../home/ubuntu/workSpaces/go/src/github.com/qqzeng/go-code-snippet/interface/method_cases_1.go+0
    rel 38+4 t=29 go.info."".Calculator+0
    rel 42+4 t=29 go.loc."".main+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=29
    0x0000 04 02 03 15 14 0a a5 9c 06 73 06 6c 06 5f 06 60  .........s.l._.`
    0x0010 88 06 37 06 60 03 7d 6f 04 01 03 66 01           ..7.`.}o...f.
""..inittask SNOPTRDATA size=24
    0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0010 00 00 00 00 00 00 00 00                          ........
runtime.memequal32·f SRODATA dupok size=8
    0x0000 00 00 00 00 00 00 00 00                          ........
    rel 0+8 t=1 runtime.memequal32+0
runtime.memequal64·f SRODATA dupok size=8
    0x0000 00 00 00 00 00 00 00 00                          ........
    rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
    0x0000 01                                               .
go.loc."".(*Calculator).Value SDWARFLOC dupok size=0
go.info."".(*Calculator).Value SDWARFINFO dupok size=60
    0x0000 03 22 22 2e 28 2a 43 61 6c 63 75 6c 61 74 6f 72  ."".(*Calculator
    0x0010 29 2e 56 61 6c 75 65 00 00 00 00 00 00 00 00 00  ).Value.........
    0x0020 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f  ................
    0x0030 7e 72 30 00 01 0d 00 00 00 00 00 00              ~r0.........
    rel 0+0 t=24 type.*"".Calculator+0
    rel 0+0 t=24 type.int32+0
    rel 24+8 t=1 "".(*Calculator).Value+0
    rel 32+8 t=1 "".(*Calculator).Value+104
    rel 42+4 t=30 gofile..<autogenerated>+0
    rel 54+4 t=29 go.info.int32+0
go.range."".(*Calculator).Value SDWARFRANGE dupok size=0
go.debuglines."".(*Calculator).Value SDWARFMISC dupok size=20
    0x0000 04 01 0f 0a a5 06 08 37 06 41 06 23 06 08 55 04  .......7.A.#..U.
    0x0010 01 03 00 01                                      ....
type..namedata.*main.Calculator. SRODATA dupok size=19
    0x0000 01 00 10 2a 6d 61 69 6e 2e 43 61 6c 63 75 6c 61  ...*main.Calcula
    0x0010 74 6f 72                                         tor
type..namedata.*func(*main.Calculator, int32)- SRODATA dupok size=33
    0x0000 00 00 1e 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 43  ...*func(*main.C
    0x0010 61 6c 63 75 6c 61 74 6f 72 2c 20 69 6e 74 33 32  alculator, int32
    0x0020 29                                               )
type.*func(*"".Calculator, int32) SRODATA dupok size=56
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 1a 5b 75 c5 08 08 08 36 00 00 00 00 00 00 00 00  .[u....6........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 00 00 00 00 00                          ........
    rel 24+8 t=1 runtime.memequal64·f+0
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func(*main.Calculator, int32)-+0
    rel 48+8 t=1 type.func(*"".Calculator, int32)+0
type.func(*"".Calculator, int32) SRODATA dupok size=72
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 4b ca 77 d1 02 08 08 33 00 00 00 00 00 00 00 00  K.w....3........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 02 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0040 00 00 00 00 00 00 00 00                          ........
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func(*main.Calculator, int32)-+0
    rel 44+4 t=6 type.*func(*"".Calculator, int32)+0
    rel 56+8 t=1 type.*"".Calculator+0
    rel 64+8 t=1 type.int32+0
type..namedata.*func(*main.Calculator) int32- SRODATA dupok size=32
    0x0000 00 00 1d 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 43  ...*func(*main.C
    0x0010 61 6c 63 75 6c 61 74 6f 72 29 20 69 6e 74 33 32  alculator) int32
type.*func(*"".Calculator) int32 SRODATA dupok size=56
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 3f fe 39 b8 08 08 08 36 00 00 00 00 00 00 00 00  ?.9....6........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 00 00 00 00 00                          ........
    rel 24+8 t=1 runtime.memequal64·f+0
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func(*main.Calculator) int32-+0
    rel 48+8 t=1 type.func(*"".Calculator) int32+0
type.func(*"".Calculator) int32 SRODATA dupok size=72
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 5d c1 43 9d 02 08 08 33 00 00 00 00 00 00 00 00  ].C....3........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0040 00 00 00 00 00 00 00 00                          ........
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func(*main.Calculator) int32-+0
    rel 44+4 t=6 type.*func(*"".Calculator) int32+0
    rel 56+8 t=1 type.*"".Calculator+0
    rel 64+8 t=1 type.int32+0
type..namedata.Add. SRODATA dupok size=6
    0x0000 01 00 03 41 64 64                                ...Add
type..namedata.*func(int32)- SRODATA dupok size=15
    0x0000 00 00 0c 2a 66 75 6e 63 28 69 6e 74 33 32 29     ...*func(int32)
type.*func(int32) SRODATA dupok size=56
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 b8 13 8d 36 08 08 08 36 00 00 00 00 00 00 00 00  ...6...6........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 00 00 00 00 00                          ........
    rel 24+8 t=1 runtime.memequal64·f+0
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func(int32)-+0
    rel 48+8 t=1 type.func(int32)+0
type.func(int32) SRODATA dupok size=64
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 0f 2f 2a b2 02 08 08 33 00 00 00 00 00 00 00 00  ./*....3........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func(int32)-+0
    rel 44+4 t=6 type.*func(int32)+0
    rel 56+8 t=1 type.int32+0
type..namedata.SetValue. SRODATA dupok size=11
    0x0000 01 00 08 53 65 74 56 61 6c 75 65                 ...SetValue
type..namedata.Value. SRODATA dupok size=8
    0x0000 01 00 05 56 61 6c 75 65                          ...Value
type..namedata.*func() int32- SRODATA dupok size=16
    0x0000 00 00 0d 2a 66 75 6e 63 28 29 20 69 6e 74 33 32  ...*func() int32
type.*func() int32 SRODATA dupok size=56
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 37 4d 0d 79 08 08 08 36 00 00 00 00 00 00 00 00  7M.y...6........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 00 00 00 00 00                          ........
    rel 24+8 t=1 runtime.memequal64·f+0
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func() int32-+0
    rel 48+8 t=1 type.func() int32+0
type.func() int32 SRODATA dupok size=64
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 c5 b9 7d 20 02 08 08 33 00 00 00 00 00 00 00 00  ..} ...3........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func() int32-+0
    rel 44+4 t=6 type.*func() int32+0
    rel 56+8 t=1 type.int32+0
type.*"".Calculator SRODATA size=120
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 5e 33 ca c8 09 08 08 36 00 00 00 00 00 00 00 00  ^3.....6........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 00 00 00 00 00 00 00 00 00 03 00 03 00  ................
    0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0070 00 00 00 00 00 00 00 00                          ........
    rel 24+8 t=1 runtime.memequal64·f+0
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*main.Calculator.+0
    rel 48+8 t=1 type."".Calculator+0
    rel 56+4 t=5 type..importpath."".+0
    rel 72+4 t=5 type..namedata.Add.+0
    rel 76+4 t=25 type.func(int32)+0
    rel 80+4 t=25 "".(*Calculator).Add+0
    rel 84+4 t=25 "".(*Calculator).Add+0
    rel 88+4 t=5 type..namedata.SetValue.+0
    rel 92+4 t=25 type.func(int32)+0
    rel 96+4 t=25 "".(*Calculator).SetValue+0
    rel 100+4 t=25 "".(*Calculator).SetValue+0
    rel 104+4 t=5 type..namedata.Value.+0
    rel 108+4 t=25 type.func() int32+0
    rel 112+4 t=25 "".(*Calculator).Value+0
    rel 116+4 t=25 "".(*Calculator).Value+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.*func(main.Calculator) int32- SRODATA dupok size=31
    0x0000 00 00 1c 2a 66 75 6e 63 28 6d 61 69 6e 2e 43 61  ...*func(main.Ca
    0x0010 6c 63 75 6c 61 74 6f 72 29 20 69 6e 74 33 32     lculator) int32
type.*func("".Calculator) int32 SRODATA dupok size=56
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 e2 04 01 11 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 00 00 00 00 00                          ........
    rel 24+8 t=1 runtime.memequal64·f+0
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func(main.Calculator) int32-+0
    rel 48+8 t=1 type.func("".Calculator) int32+0
type.func("".Calculator) int32 SRODATA dupok size=72
    0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    0x0010 1c 77 64 14 02 08 08 33 00 00 00 00 00 00 00 00  .wd....3........
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0040 00 00 00 00 00 00 00 00                          ........
    rel 32+8 t=1 runtime.gcbits.01+0
    rel 40+4 t=5 type..namedata.*func(main.Calculator) int32-+0
    rel 44+4 t=6 type.*func("".Calculator) int32+0
    rel 56+8 t=1 type."".Calculator+0
    rel 64+8 t=1 type.int32+0
type..namedata.val- SRODATA dupok size=6
    0x0000 00 00 03 76 61 6c                                ...val
type."".Calculator SRODATA size=136
    0x0000 04 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0010 b2 95 ce 57 0f 04 04 19 00 00 00 00 00 00 00 00  ...W............
    0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
    0x0050 00 00 00 00 01 00 01 00 28 00 00 00 00 00 00 00  ........(.......
    0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    0x0080 00 00 00 00 00 00 00 00                          ........
    rel 24+8 t=1 runtime.memequal32·f+0
    rel 32+8 t=1 runtime.gcbits.+0
    rel 40+4 t=5 type..namedata.*main.Calculator.+0
    rel 44+4 t=5 type.*"".Calculator+0
    rel 48+8 t=1 type..importpath."".+0
    rel 56+8 t=1 type."".Calculator+96
    rel 80+4 t=5 type..importpath."".+0
    rel 96+8 t=1 type..namedata.val-+0
    rel 104+8 t=1 type.int32+0
    rel 120+4 t=5 type..namedata.Value.+0
    rel 124+4 t=25 type.func() int32+0
    rel 128+4 t=25 "".(*Calculator).Value+0
    rel 132+4 t=25 "".Calculator.Value+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
    0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
    0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
    0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·568470801006e5c0dc3947ea998fe279 SRODATA dupok size=10
    0x0000 02 00 00 00 02 00 00 00 00 02                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
    0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
