// AUTO-GENERATED by mkbuiltin.go; DO NOT EDIT

package gc

var runtimeDecls = [...]struct {
	name string
	tag  int
	typ  int
}{
	{"newobject", funcTag, 4},
	{"panicindex", funcTag, 5},
	{"panicslice", funcTag, 5},
	{"panicdivide", funcTag, 5},
	{"throwinit", funcTag, 5},
	{"panicwrap", funcTag, 7},
	{"gopanic", funcTag, 9},
	{"gorecover", funcTag, 12},
	{"printbool", funcTag, 14},
	{"printfloat", funcTag, 16},
	{"printint", funcTag, 18},
	{"printhex", funcTag, 20},
	{"printuint", funcTag, 20},
	{"printcomplex", funcTag, 22},
	{"printstring", funcTag, 23},
	{"printpointer", funcTag, 24},
	{"printiface", funcTag, 24},
	{"printeface", funcTag, 24},
	{"printslice", funcTag, 24},
	{"printnl", funcTag, 5},
	{"printsp", funcTag, 5},
	{"printlock", funcTag, 5},
	{"printunlock", funcTag, 5},
	{"concatstring2", funcTag, 27},
	{"concatstring3", funcTag, 28},
	{"concatstring4", funcTag, 29},
	{"concatstring5", funcTag, 30},
	{"concatstrings", funcTag, 32},
	{"cmpstring", funcTag, 34},
	{"eqstring", funcTag, 35},
	{"intstring", funcTag, 38},
	{"slicebytetostring", funcTag, 40},
	{"slicebytetostringtmp", funcTag, 41},
	{"slicerunetostring", funcTag, 44},
	{"stringtoslicebyte", funcTag, 45},
	{"stringtoslicebytetmp", funcTag, 46},
	{"stringtoslicerune", funcTag, 49},
	{"decoderune", funcTag, 50},
	{"slicecopy", funcTag, 52},
	{"slicestringcopy", funcTag, 53},
	{"convI2E", funcTag, 54},
	{"convI2I", funcTag, 55},
	{"convT2E", funcTag, 56},
	{"convT2I", funcTag, 56},
	{"assertE2E", funcTag, 57},
	{"assertE2E2", funcTag, 58},
	{"assertE2I", funcTag, 57},
	{"assertE2I2", funcTag, 58},
	{"assertE2T", funcTag, 57},
	{"assertE2T2", funcTag, 58},
	{"assertI2E", funcTag, 57},
	{"assertI2E2", funcTag, 58},
	{"assertI2I", funcTag, 57},
	{"assertI2I2", funcTag, 58},
	{"assertI2T", funcTag, 57},
	{"assertI2T2", funcTag, 58},
	{"panicdottype", funcTag, 59},
	{"ifaceeq", funcTag, 60},
	{"efaceeq", funcTag, 60},
	{"makemap", funcTag, 62},
	{"mapaccess1", funcTag, 63},
	{"mapaccess1_fast32", funcTag, 64},
	{"mapaccess1_fast64", funcTag, 64},
	{"mapaccess1_faststr", funcTag, 64},
	{"mapaccess1_fat", funcTag, 65},
	{"mapaccess2", funcTag, 66},
	{"mapaccess2_fast32", funcTag, 67},
	{"mapaccess2_fast64", funcTag, 67},
	{"mapaccess2_faststr", funcTag, 67},
	{"mapaccess2_fat", funcTag, 68},
	{"mapassign", funcTag, 63},
	{"mapiterinit", funcTag, 69},
	{"mapdelete", funcTag, 69},
	{"mapiternext", funcTag, 70},
	{"makechan", funcTag, 72},
	{"chanrecv1", funcTag, 74},
	{"chanrecv2", funcTag, 75},
	{"chansend1", funcTag, 77},
	{"closechan", funcTag, 24},
	{"writeBarrier", varTag, 78},
	{"writebarrierptr", funcTag, 79},
	{"typedmemmove", funcTag, 80},
	{"typedslicecopy", funcTag, 81},
	{"selectnbsend", funcTag, 82},
	{"selectnbrecv", funcTag, 83},
	{"selectnbrecv2", funcTag, 85},
	{"newselect", funcTag, 86},
	{"selectsend", funcTag, 82},
	{"selectrecv", funcTag, 75},
	{"selectrecv2", funcTag, 87},
	{"selectdefault", funcTag, 88},
	{"selectgo", funcTag, 89},
	{"block", funcTag, 5},
	{"makeslice", funcTag, 91},
	{"makeslice64", funcTag, 92},
	{"growslice", funcTag, 93},
	{"memmove", funcTag, 94},
	{"memclr", funcTag, 95},
	{"memequal", funcTag, 96},
	{"memequal8", funcTag, 97},
	{"memequal16", funcTag, 97},
	{"memequal32", funcTag, 97},
	{"memequal64", funcTag, 97},
	{"memequal128", funcTag, 97},
	{"int64div", funcTag, 98},
	{"uint64div", funcTag, 99},
	{"int64mod", funcTag, 98},
	{"uint64mod", funcTag, 99},
	{"float64toint64", funcTag, 100},
	{"float64touint64", funcTag, 101},
	{"float64touint32", funcTag, 103},
	{"int64tofloat64", funcTag, 104},
	{"uint64tofloat64", funcTag, 105},
	{"uint32tofloat64", funcTag, 106},
	{"complex128div", funcTag, 107},
	{"racefuncenter", funcTag, 108},
	{"racefuncexit", funcTag, 5},
	{"raceread", funcTag, 108},
	{"racewrite", funcTag, 108},
	{"racereadrange", funcTag, 109},
	{"racewriterange", funcTag, 109},
	{"msanread", funcTag, 109},
	{"msanwrite", funcTag, 109},
}

func runtimeTypes() []*Type {
	var typs [110]*Type
	typs[0] = bytetype
	typs[1] = typPtr(typs[0])
	typs[2] = Types[TANY]
	typs[3] = typPtr(typs[2])
	typs[4] = functype(nil, []*Node{anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[5] = functype(nil, nil, nil)
	typs[6] = Types[TSTRING]
	typs[7] = functype(nil, []*Node{anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6])}, nil)
	typs[8] = Types[TINTER]
	typs[9] = functype(nil, []*Node{anonfield(typs[8])}, nil)
	typs[10] = Types[TINT32]
	typs[11] = typPtr(typs[10])
	typs[12] = functype(nil, []*Node{anonfield(typs[11])}, []*Node{anonfield(typs[8])})
	typs[13] = Types[TBOOL]
	typs[14] = functype(nil, []*Node{anonfield(typs[13])}, nil)
	typs[15] = Types[TFLOAT64]
	typs[16] = functype(nil, []*Node{anonfield(typs[15])}, nil)
	typs[17] = Types[TINT64]
	typs[18] = functype(nil, []*Node{anonfield(typs[17])}, nil)
	typs[19] = Types[TUINT64]
	typs[20] = functype(nil, []*Node{anonfield(typs[19])}, nil)
	typs[21] = Types[TCOMPLEX128]
	typs[22] = functype(nil, []*Node{anonfield(typs[21])}, nil)
	typs[23] = functype(nil, []*Node{anonfield(typs[6])}, nil)
	typs[24] = functype(nil, []*Node{anonfield(typs[2])}, nil)
	typs[25] = typArray(typs[0], 32)
	typs[26] = typPtr(typs[25])
	typs[27] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[6])})
	typs[28] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[6])})
	typs[29] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[6])})
	typs[30] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[6])})
	typs[31] = typSlice(typs[6])
	typs[32] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[31])}, []*Node{anonfield(typs[6])})
	typs[33] = Types[TINT]
	typs[34] = functype(nil, []*Node{anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[33])})
	typs[35] = functype(nil, []*Node{anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[13])})
	typs[36] = typArray(typs[0], 4)
	typs[37] = typPtr(typs[36])
	typs[38] = functype(nil, []*Node{anonfield(typs[37]), anonfield(typs[17])}, []*Node{anonfield(typs[6])})
	typs[39] = typSlice(typs[0])
	typs[40] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[39])}, []*Node{anonfield(typs[6])})
	typs[41] = functype(nil, []*Node{anonfield(typs[39])}, []*Node{anonfield(typs[6])})
	typs[42] = runetype
	typs[43] = typSlice(typs[42])
	typs[44] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[43])}, []*Node{anonfield(typs[6])})
	typs[45] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6])}, []*Node{anonfield(typs[39])})
	typs[46] = functype(nil, []*Node{anonfield(typs[6])}, []*Node{anonfield(typs[39])})
	typs[47] = typArray(typs[42], 32)
	typs[48] = typPtr(typs[47])
	typs[49] = functype(nil, []*Node{anonfield(typs[48]), anonfield(typs[6])}, []*Node{anonfield(typs[43])})
	typs[50] = functype(nil, []*Node{anonfield(typs[6]), anonfield(typs[33])}, []*Node{anonfield(typs[42]), anonfield(typs[33])})
	typs[51] = Types[TUINTPTR]
	typs[52] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2]), anonfield(typs[51])}, []*Node{anonfield(typs[33])})
	typs[53] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[33])})
	typs[54] = functype(nil, []*Node{anonfield(typs[2])}, []*Node{anonfield(typs[2])})
	typs[55] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2])}, []*Node{anonfield(typs[2])})
	typs[56] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, []*Node{anonfield(typs[2])})
	typs[57] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[3])}, nil)
	typs[58] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[3])}, []*Node{anonfield(typs[13])})
	typs[59] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[1])}, nil)
	typs[60] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[13])})
	typs[61] = typMap(typs[2], typs[2])
	typs[62] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[17]), anonfield(typs[3]), anonfield(typs[3])}, []*Node{anonfield(typs[61])})
	typs[63] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[61]), anonfield(typs[3])}, []*Node{anonfield(typs[3])})
	typs[64] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[61]), anonfield(typs[2])}, []*Node{anonfield(typs[3])})
	typs[65] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[61]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[66] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[61]), anonfield(typs[3])}, []*Node{anonfield(typs[3]), anonfield(typs[13])})
	typs[67] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[61]), anonfield(typs[2])}, []*Node{anonfield(typs[3]), anonfield(typs[13])})
	typs[68] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[61]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3]), anonfield(typs[13])})
	typs[69] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[61]), anonfield(typs[3])}, nil)
	typs[70] = functype(nil, []*Node{anonfield(typs[3])}, nil)
	typs[71] = typChan(typs[2], Cboth)
	typs[72] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[17])}, []*Node{anonfield(typs[71])})
	typs[73] = typChan(typs[2], Crecv)
	typs[74] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[73]), anonfield(typs[3])}, nil)
	typs[75] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[73]), anonfield(typs[3])}, []*Node{anonfield(typs[13])})
	typs[76] = typChan(typs[2], Csend)
	typs[77] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[76]), anonfield(typs[3])}, nil)
	typs[78] = tostruct([]*Node{namedfield("enabled", typs[13]), namedfield("needed", typs[13]), namedfield("cgo", typs[13])})
	typs[79] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[2])}, nil)
	typs[80] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[3])}, nil)
	typs[81] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[33])})
	typs[82] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[76]), anonfield(typs[3])}, []*Node{anonfield(typs[13])})
	typs[83] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[73])}, []*Node{anonfield(typs[13])})
	typs[84] = typPtr(typs[13])
	typs[85] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[84]), anonfield(typs[73])}, []*Node{anonfield(typs[13])})
	typs[86] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[17]), anonfield(typs[10])}, nil)
	typs[87] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[73]), anonfield(typs[3]), anonfield(typs[84])}, []*Node{anonfield(typs[13])})
	typs[88] = functype(nil, []*Node{anonfield(typs[1])}, []*Node{anonfield(typs[13])})
	typs[89] = functype(nil, []*Node{anonfield(typs[1])}, nil)
	typs[90] = typSlice(typs[2])
	typs[91] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[33]), anonfield(typs[33])}, []*Node{anonfield(typs[90])})
	typs[92] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[17]), anonfield(typs[17])}, []*Node{anonfield(typs[90])})
	typs[93] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[90]), anonfield(typs[33])}, []*Node{anonfield(typs[90])})
	typs[94] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[51])}, nil)
	typs[95] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[51])}, nil)
	typs[96] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[51])}, []*Node{anonfield(typs[13])})
	typs[97] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3])}, []*Node{anonfield(typs[13])})
	typs[98] = functype(nil, []*Node{anonfield(typs[17]), anonfield(typs[17])}, []*Node{anonfield(typs[17])})
	typs[99] = functype(nil, []*Node{anonfield(typs[19]), anonfield(typs[19])}, []*Node{anonfield(typs[19])})
	typs[100] = functype(nil, []*Node{anonfield(typs[15])}, []*Node{anonfield(typs[17])})
	typs[101] = functype(nil, []*Node{anonfield(typs[15])}, []*Node{anonfield(typs[19])})
	typs[102] = Types[TUINT32]
	typs[103] = functype(nil, []*Node{anonfield(typs[15])}, []*Node{anonfield(typs[102])})
	typs[104] = functype(nil, []*Node{anonfield(typs[17])}, []*Node{anonfield(typs[15])})
	typs[105] = functype(nil, []*Node{anonfield(typs[19])}, []*Node{anonfield(typs[15])})
	typs[106] = functype(nil, []*Node{anonfield(typs[102])}, []*Node{anonfield(typs[15])})
	typs[107] = functype(nil, []*Node{anonfield(typs[21]), anonfield(typs[21])}, []*Node{anonfield(typs[21])})
	typs[108] = functype(nil, []*Node{anonfield(typs[51])}, nil)
	typs[109] = functype(nil, []*Node{anonfield(typs[51]), anonfield(typs[51])}, nil)
	return typs[:]
}