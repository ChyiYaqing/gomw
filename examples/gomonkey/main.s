main.a STEXT size=16 args=0x0 locals=0x0 funcid=0x0 align=0x0 leaf
	0x0000 00000 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:3)	TEXT	main.a(SB), LEAF|NOFRAME|ABIInternal, $0-0
	0x0000 00000 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:3)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:3)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:4)	MOVD	$1, R0
	0x0004 00004 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:4)	RET	(R30)
	0x0000 e0 03 40 b2 c0 03 5f d6 00 00 00 00 00 00 00 00  ..@..._.........
main.main STEXT size=80 args=0x0 locals=0x18 funcid=0x0 align=0x0
	0x0000 00000 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	TEXT	main.main(SB), ABIInternal, $32-0
	0x0000 00000 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	MOVD	16(g), R16
	0x0004 00004 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	PCDATA	$0, $-2
	0x0004 00004 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	CMP	R16, RSP
	0x0008 00008 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	BLS	56
	0x000c 00012 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	PCDATA	$0, $-1
	0x000c 00012 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	MOVD.W	R30, -32(RSP)
	0x0010 00016 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	MOVD	R29, -8(RSP)
	0x0014 00020 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	SUB	$8, RSP, R29
	0x0018 00024 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0018 00024 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0018 00024 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:8)	PCDATA	$1, $0
	0x0018 00024 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:8)	CALL	runtime.printlock(SB)
	0x001c 00028 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:8)	MOVD	$main.a·f(SB), R0
	0x0024 00036 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:8)	CALL	runtime.printpointer(SB)
	0x0028 00040 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:8)	CALL	runtime.printunlock(SB)
	0x002c 00044 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:9)	LDP	-8(RSP), (R29, R30)
	0x0030 00048 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:9)	ADD	$32, RSP
	0x0034 00052 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:9)	RET	(R30)
	0x0038 00056 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:9)	NOP
	0x0038 00056 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	PCDATA	$1, $-1
	0x0038 00056 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	PCDATA	$0, $-2
	0x0038 00056 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	MOVD	R30, R3
	0x003c 00060 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0040 00064 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	PCDATA	$0, $-1
	0x0040 00064 (/Users/chyiyaqing/chyi/github.com/gomw/examples/gomonkey/main.go:7)	JMP	0
	0x0000 90 0b 40 f9 ff 63 30 eb 89 01 00 54 fe 0f 1e f8  ..@..c0....T....
	0x0010 fd 83 1f f8 fd 23 00 d1 00 00 00 94 00 00 00 90  .....#..........
	0x0020 00 00 00 91 00 00 00 94 00 00 00 94 fd fb 7f a9  ................
	0x0030 ff 83 00 91 c0 03 5f d6 e3 03 1e aa 00 00 00 94  ......_.........
	0x0040 f0 ff ff 17 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+4 t=9 runtime.printlock+0
	rel 28+8 t=3 main.a·f+0
	rel 36+4 t=9 runtime.printpointer+0
	rel 40+4 t=9 runtime.printunlock+0
	rel 60+4 t=9 runtime.morestack_noctxt+0
go:cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
main..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
main.a·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 main.a+0
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
