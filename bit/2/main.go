package main

import (
	"fmt"
	"strings"
)

const (
	UPPER = 1 << iota
	LOWER
	CAP
	REV
)

func main() {
	// Decimal
	// 00,01,02,03,04,05,06,07,08,09,0A,0B,0C,0D,0E,0F -> 16
	var bitss = []uint8{0x41, 0x42, 0x43}
	var x uint8 = 0x14
	var y uint8 = 0xA9
	fmt.Println(x)

	fmt.Printf("%8b\n", y)
	fmt.Println(string(y))
	fmt.Println(y)

	fmt.Println(string(bitss[:]))

	// 1111 1111 1111 1111 = 65535
	fmt.Println(0xffff)
	fmt.Println(0xf29a)
	// 1111 1111
	fmt.Println(0xff)

	fmt.Println(procstr("HELLO PEOPLE", LOWER|CAP|REV))

}

func procstr(str string, conf byte) string {
	// reverse string
	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	// query config bits
	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str
}

// Dec	Hex	Char	Octal	Raw
// encoding	UTF8
// encoding	HTML
// entity	Description
// 0	0000		000	0x00	0x00		NUL
// 1	0001		001	0x01	0x01		SOH
// 2	0002		002	0x02	0x02		STX
// 3	0003		003	0x03	0x03		ETX
// 4	0004		004	0x04	0x04		EOT
// 5	0005		005	0x05	0x05		ENQ
// 6	0006		006	0x06	0x06		ACK
// 7	0007		007	0x07	0x07		BEL, bell, alarm, \a
// 8	0008		010	0x08	0x08		BS, backspace, \b
// 9	0009		011	0x09	0x09		HT, tab, \t
// 10	000a		012	0x0A	0x0A		LF, line feed, \cj
// 11	000b		013	0x0B	0x0B		VT
// 12	000c		014	0x0C	0x0C		FF, NP, form feed, \f
// 13	000d		015	0x0D	0x0D		CR, carriage return, \cm
// 14	000e		016	0x0E	0x0E		SO
// 15	000f		017	0x0F	0x0F		SI
// 16	0010		020	0x10	0x10		DLE
// 17	0011		021	0x11	0x11		DC1
// 18	0012		022	0x12	0x12		DC2
// 19	0013		023	0x13	0x13		DC3
// 20	0014		024	0x14	0x14		DC4
// 21	0015		025	0x15	0x15		NAK
// 22	0016		026	0x16	0x16		SYN
// 23	0017		027	0x17	0x17		ETB
// 24	0018		030	0x18	0x18		CAN
// 25	0019		031	0x19	0x19		EM
// 26	001a		032	0x1A	0x1A		SUB
// 27	001b		033	0x1B	0x1B		ESC, escape, \e
// 28	001c		034	0x1C	0x1C		FS
// 29	001d		035	0x1D	0x1D		GS
// 30	001e		036	0x1E	0x1E		RS
// 31	001f		037	0x1F	0x1F		US
// 32	0020		040	0x20	0x20		SPC, space
// 33	0021	!	041	0x21	0x21		Exclamation point, bang
// 34	0022	"	042	0x22	0x22	&quot;	Quote, double quote
// 35	0023	#	043	0x23	0x23		Number, pound, hash
// 36	0024	$	044	0x24	0x24		Dollar
// 37	0025	%	045	0x25	0x25		Percent
// 38	0026	&	046	0x26	0x26	&amp;	Ampersand, and
// 39	0027	'	047	0x27	0x27	&apos;	Apostrophe, single quote
// 40	0028	(	050	0x28	0x28		Open parenthesis, open parens
// 41	0029	)	051	0x29	0x29		Close parenthesis, close parens
// 42	002a	*	052	0x2A	0x2A		Asterisk, star, glob
// 43	002b	+	053	0x2B	0x2B		Plus
// 44	002c	,	054	0x2C	0x2C		Comma
// 45	002d	-	055	0x2D	0x2D		Hyphen, dash, minus
// 46	002e	.	056	0x2E	0x2E		Period, dot, decimal, full stop
// 47	002f	/	057	0x2F	0x2F		Slash, forward slash, stroke, virgule, solidus
// 48	0030	0	060	0x30	0x30
// 49	0031	1	061	0x31	0x31
// 50	0032	2	062	0x32	0x32
// 51	0033	3	063	0x33	0x33
// 52	0034	4	064	0x34	0x34
// 53	0035	5	065	0x35	0x35
// 54	0036	6	066	0x36	0x36
// 55	0037	7	067	0x37	0x37
// 56	0038	8	070	0x38	0x38
// 57	0039	9	071	0x39	0x39
// 58	003a	:	072	0x3A	0x3A		Colon
// 59	003b	;	073	0x3B	0x3B		Semicolon
// 60	003c	<	074	0x3C	0x3C	&lt;	Less-than sign
// 61	003d	=	075	0x3D	0x3D		Equals sign
// 62	003e	>	076	0x3E	0x3E	&gt;	Greater-than sign
// 63	003f	?	077	0x3F	0x3F		Question mark
// 64	0040	@	100	0x40	0x40		At sign
// 65	0041	A	101	0x41	0x41
// 66	0042	B	102	0x42	0x42
// 67	0043	C	103	0x43	0x43
// 68	0044	D	104	0x44	0x44
// 69	0045	E	105	0x45	0x45
// 70	0046	F	106	0x46	0x46
// 71	0047	G	107	0x47	0x47
// 72	0048	H	110	0x48	0x48
// 73	0049	I	111	0x49	0x49
// 74	004a	J	112	0x4A	0x4A
// 75	004b	K	113	0x4B	0x4B
// 76	004c	L	114	0x4C	0x4C
// 77	004d	M	115	0x4D	0x4D
// 78	004e	N	116	0x4E	0x4E
// 79	004f	O	117	0x4F	0x4F
// 80	0050	P	120	0x50	0x50
// 81	0051	Q	121	0x51	0x51
// 82	0052	R	122	0x52	0x52
// 83	0053	S	123	0x53	0x53
// 84	0054	T	124	0x54	0x54
// 85	0055	U	125	0x55	0x55
// 86	0056	V	126	0x56	0x56
// 87	0057	W	127	0x57	0x57
// 88	0058	X	130	0x58	0x58
// 89	0059	Y	131	0x59	0x59
// 90	005a	Z	132	0x5A	0x5A
// 91	005b	[	133	0x5B	0x5B		Left (square) bracket, open (square) bracket
// 92	005c	\	134	0x5C	0x5C		Backslash
// 93	005d	]	135	0x5D	0x5D		Right (square) bracket, close (square) bracket
// 94	005e	^	136	0x5E	0x5E		Caret, up-arrow, circumflex
// 95	005f	_	137	0x5F	0x5F		Underscore
// 96	0060	`	140	0x60	0x60		Backtick, backquote
// 97	0061	a	141	0x61	0x61
// 98	0062	b	142	0x62	0x62
// 99	0063	c	143	0x63	0x63
// 100	0064	d	144	0x64	0x64
// 101	0065	e	145	0x65	0x65
// 102	0066	f	146	0x66	0x66
// 103	0067	g	147	0x67	0x67
// 104	0068	h	150	0x68	0x68
// 105	0069	i	151	0x69	0x69
// 106	006a	j	152	0x6A	0x6A
// 107	006b	k	153	0x6B	0x6B
// 108	006c	l	154	0x6C	0x6C
// 109	006d	m	155	0x6D	0x6D
// 110	006e	n	156	0x6E	0x6E
// 111	006f	o	157	0x6F	0x6F
// 112	0070	p	160	0x70	0x70
// 113	0071	q	161	0x71	0x71
// 114	0072	r	162	0x72	0x72
// 115	0073	s	163	0x73	0x73
// 116	0074	t	164	0x74	0x74
// 117	0075	u	165	0x75	0x75
// 118	0076	v	166	0x76	0x76
// 119	0077	w	167	0x77	0x77
// 120	0078	x	170	0x78	0x78
// 121	0079	y	171	0x79	0x79
// 122	007a	z	172	0x7A	0x7A
// 123	007b	{	173	0x7B	0x7B		Open brace
// 124	007c	|	174	0x7C	0x7C		Pipe, vertical bar
// 125	007d	}	175	0x7D	0x7D		Close brace
// 126	007e	~	176	0x7E	0x7E		Tilde, twiddle, squiggle
// 127	007f		177	0x7F	0x7F		DEL, delete
// 128	0080		200	0x80	0xC2,0x80		(Undefined)
// 129	0081		201	0x81	0xC2,0x81		(Undefined)
// 130	0082		202	0x82	0xC2,0x82		(Undefined)
// 131	0083		203	0x83	0xC2,0x83		(Undefined)
// 132	0084		204	0x84	0xC2,0x84		(Undefined)
// 133	0085		205	0x85	0xC2,0x85		(Undefined)
// 134	0086		206	0x86	0xC2,0x86		(Undefined)
// 135	0087		207	0x87	0xC2,0x87		(Undefined)
// 136	0088		210	0x88	0xC2,0x88		(Undefined)
// 137	0089		211	0x89	0xC2,0x89		(Undefined)
// 138	008a		212	0x8A	0xC2,0x8A		(Undefined)
// 139	008b		213	0x8B	0xC2,0x8B		(Undefined)
// 140	008c		214	0x8C	0xC2,0x8C		(Undefined)
// 141	008d		215	0x8D	0xC2,0x8D		(Undefined)
// 142	008e		216	0x8E	0xC2,0x8E		(Undefined)
// 143	008f		217	0x8F	0xC2,0x8F		(Undefined)
// 144	0090		220	0x90	0xC2,0x90		(Undefined)
// 145	0091		221	0x91	0xC2,0x91		(Undefined)
// 146	0092		222	0x92	0xC2,0x92		(Undefined)
// 147	0093		223	0x93	0xC2,0x93		(Undefined)
// 148	0094		224	0x94	0xC2,0x94		(Undefined)
// 149	0095		225	0x95	0xC2,0x95		(Undefined)
// 150	0096		226	0x96	0xC2,0x96		(Undefined)
// 151	0097		227	0x97	0xC2,0x97		(Undefined)
// 152	0098		230	0x98	0xC2,0x98		(Undefined)
// 153	0099		231	0x99	0xC2,0x99		(Undefined)
// 154	009a		232	0x9A	0xC2,0x9A		(Undefined)
// 155	009b		233	0x9B	0xC2,0x9B		(Undefined)
// 156	009c		234	0x9C	0xC2,0x9C		(Undefined)
// 157	009d		235	0x9D	0xC2,0x9D		(Undefined)
// 158	009e		236	0x9E	0xC2,0x9E		(Undefined)
// 159	009f		237	0x9F	0xC2,0x9F		(Undefined)
// 160	00a0		240	0xA0	0xC2,0xA0	&nbsp;	No-break space, nonbreaking space
// 161	00a1	¡	241	0xA1	0xC2,0xA1	&iexcl;	Inverted exclamation mark
// 162	00a2	¢	242	0xA2	0xC2,0xA2	&cent;	Cent sign
// 163	00a3	£	243	0xA3	0xC2,0xA3	&pound;	Pound sign
// 164	00a4	¤	244	0xA4	0xC2,0xA4	&curren;	Currency sign
// 165	00a5	¥	245	0xA5	0xC2,0xA5	&yen;	Yen sign, yuan sign
// 166	00a6	¦	246	0xA6	0xC2,0xA6	&brvbar;	Broken bar, broken vertical bar
// 167	00a7	§	247	0xA7	0xC2,0xA7	&sect;	Section sign
// 168	00a8	¨	250	0xA8	0xC2,0xA8	&uml;	Diaeresis, spacing diaeresis
// 169	00a9	©	251	0xA9	0xC2,0xA9	&copy;	Copyright sign
// 170	00aa	ª	252	0xAA	0xC2,0xAA	&ordf;	Feminine ordinal indicator
// 171	00ab	«	253	0xAB	0xC2,0xAB	&laquo;	Left-pointing double angle quotation mark, left pointing guillemet
// 172	00ac	¬	254	0xAC	0xC2,0xAC	&not;	Not sign, angled dash
// 173	00ad	(-)	255	0xAD	0xC2,0xAD	&shy;	Soft hyphen, discretionary hyphen
// 174	00ae	®	256	0xAE	0xC2,0xAE	&reg;	Registered sign, registered trademark sign
// 175	00af	¯	257	0xAF	0xC2,0xAF	&macr;	Macron, spacing macron, overline, APL overbar
// 176	00b0	°	260	0xB0	0xC2,0xB0	&deg;	Degree sign
// 177	00b1	±	261	0xB1	0xC2,0xB1	&plusmn;	Plus-minus sign, plus-or-minus sign
// 178	00b2	²	262	0xB2	0xC2,0xB2	&sup2;	Superscript two, superscript digit two, squared
// 179	00b3	³	263	0xB3	0xC2,0xB3	&sup3;	Superscript three, superscript digit three, cubed
// 180	00b4	´	264	0xB4	0xC2,0xB4	&acute;	Acute accent, spacing acute
// 181	00b5	μ	265	0xB5	0xC2,0xB5	&micro;	Micro sign
// 182	00b6	¶	266	0xB6	0xC2,0xB6	&para;	Pilcrow sign, paragraph sign
// 183	00b7	·	267	0xB7	0xC2,0xB7	&middot;	Middle dot, Georgian comma, Greek middle dot
// 184	00b8	¸	270	0xB8	0xC2,0xB8	&cedil;	Cedilla, spacing cedilla
// 185	00b9	¹	271	0xB9	0xC2,0xB9	&sup1;	Superscript one, superscript digit one
// 186	00ba	º	272	0xBA	0xC2,0xBA	&ordm;	Masculine ordinal indicator
// 187	00bb	»	273	0xBB	0xC2,0xBB	&raquo;	Right-pointing double angle quotation mark, right pointing guillemet
// 188	00bc	¼	274	0xBC	0xC2,0xBC	&frac14;	Vulgar fraction one quarter, fraction one quarter
// 189	00bd	½	275	0xBD	0xC2,0xBD	&frac12;	Vulgar fraction one half, fraction one half
// 190	00be	¾	276	0xBE	0xC2,0xBE	&frac34;	Vulgar fraction three quarters, fraction three quarters
// 191	00bf	¿	277	0xBF	0xC2,0xBF	&iquest;	Inverted question mark, turned question mark
// 192	00c0	À	300	0xC0	0xC3,0x80	&Agrave;	Capital A grave, capital A grave
// 193	00c1	Á	301	0xC1	0xC3,0x81	&Aacute;	Capital A acute
// 194	00c2	Â	302	0xC2	0xC3,0x82	&Acirc;	Capital A circumflex
// 195	00c3	Ã	303	0xC3	0xC3,0x83	&Atilde;	Capital A tilde
// 196	00c4	Ä	304	0xC4	0xC3,0x84	&Auml;	Capital A diaeresis
// 197	00c5	Å	305	0xC5	0xC3,0x85	&Aring;	Capital A ring above, capital A ring
// 198	00c6	Æ	306	0xC6	0xC3,0x86	&AElig;	Capital AE, capital ligature AE
// 199	00c7	Ç	307	0xC7	0xC3,0x87	&Ccedil;	Capital C cedilla
// 200	00c8	È	310	0xC8	0xC3,0x88	&Egrave;	Capital E grave
// 201	00c9	É	311	0xC9	0xC3,0x89	&Eacute;	Capital E acute
// 202	00ca	Ê	312	0xCA	0xC3,0x8A	&Ecirc;	Capital E circumflex
// 203	00cb	Ë	313	0xCB	0xC3,0x8B	&Euml;	Capital E diaeresis
// 204	00cc	Ì	314	0xCC	0xC3,0x8C	&Igrave;	Capital I grave
// 205	00cd	Í	315	0xCD	0xC3,0x8D	&Iacute;	Capital I acute
// 206	00ce	Î	316	0xCE	0xC3,0x8E	&Icirc;	Capital I circumflex
// 207	00cf	Ï	317	0xCF	0xC3,0x8F	&Iuml;	Capital I diaeresis
// 208	00d0	Ð Ð	320	0xD0	0xC3,0x90	&ETH;	Capital Eth, Edh, crossed D
// 209	00d1	Ñ	321	0xD1	0xC3,0x91	&Ntilde;	Capital N tilde
// 210	00d2	Ò	322	0xD2	0xC3,0x92	&Ograve;	Capital O grave
// 211	00d3	Ó	323	0xD3	0xC3,0x93	&Oacute;	Capital O acute
// 212	00d4	Ô	324	0xD4	0xC3,0x94	&Ocirc;	Capital O circumflex
// 213	00d5	Õ	325	0xD5	0xC3,0x95	&Otilde;	Capital O tilde
// 214	00d6	Ö	326	0xD6	0xC3,0x96	&Ouml;	Capital O diaeresis
// 215	00d7	×	327	0xD7	0xC3,0x97	&times;	Multiplication sign
// 216	00d8	Ø	330	0xD8	0xC3,0x98	&Oslash;	Capital O stroke, capital O slash
// 217	00d9	Ù	331	0xD9	0xC3,0x99	&Ugrave;	Capital U grave
// 218	00da	Ú	332	0xDA	0xC3,0x9A	&Uacute;	Capital U acute
// 219	00db	û	333	0xDB	0xC3,0x9B	&Ucirc;	Capital U circumflex
// 220	00dc	Ü	334	0xDC	0xC3,0x9C	&Uuml;	Capital U diaeresis
// 221	00dd	Ý	335	0xDD	0xC3,0x9D	&Yacute;	Capital Y acute
// 222	00de	Þ Þ	336	0xDE	0xC3,0x9E	&THORN;	Capital Thorn
// 223	00df	ß	337	0xDF	0xC3,0x9F	&szlig;	Sharp s, ess-zed, Eszett
// 224	00e0	à	340	0xE0	0xC3,0xA0	&agrave;	a grave
// 225	00e1	á	341	0xE1	0xC3,0xA1	&aacute;	a acute
// 226	00e2	â	342	0xE2	0xC3,0xA2	&acirc;	a circumflex
// 227	00e3	ã	343	0xE3	0xC3,0xA3	&atilde;	a tilde
// 228	00e4	ä	344	0xE4	0xC3,0xA4	&auml;	a diaeresis
// 229	00e5	å	345	0xE5	0xC3,0xA5	&aring;	a ring above, a ring
// 230	00e6	æ	346	0xE6	0xC3,0xA6	&aelig;	ae, ligature ae
// 231	00e7	ç	347	0xE7	0xC3,0xA7	&ccedil;	c cedilla
// 232	00e8	è	350	0xE8	0xC3,0xA8	&egrave;	e grave
// 233	00e9	é	351	0xE9	0xC3,0xA9	&eacute;	e acute
// 234	00ea	ê	352	0xEA	0xC3,0xAA	&ecirc;	e circumflex
// 235	00eb	ë	353	0xEB	0xC3,0xAB	&euml;	e diaeresis
// 236	00ec	ì	354	0xEC	0xC3,0xAC	&igrave;	i grave
// 237	00ed	í	355	0xED	0xC3,0xAD	&iacute;	i acute
// 238	00ee	î	356	0xEE	0xC3,0xAE	&icirc;	i circumflex
// 239	00ef	ï	357	0xEF	0xC3,0xAF	&iuml;	i diaeresis
// 240	00f0	ð ð	360	0xF0	0xC3,0xB0	&eth;	eth, edh, crossed d
// 241	00f1	ñ	361	0xF1	0xC3,0xB1	&ntilde;	n tilde
// 242	00f2	ò	362	0xF2	0xC3,0xB2	&ograve;	o grave
// 243	00f3	ó	363	0xF3	0xC3,0xB3	&oacute;	o acute
// 244	00f4	ô	364	0xF4	0xC3,0xB4	&ocirc;	o circumflex
// 245	00f5	õ	365	0xF5	0xC3,0xB5	&otilde;	o tilde
// 246	00f6	ö	366	0xF6	0xC3,0xB6	&ouml;	o diaeresis
// 247	00f7	÷	367	0xF7	0xC3,0xB7	&divide;	Division sign
// 248	00f8	⊘	370	0xF8	0xC3,0xB8	&oslash;	o stroke, o slash
// 249	00f9	ù	371	0xF9	0xC3,0xB9	&ugrave;	u grave
// 250	00fa	ú	372	0xFA	0xC3,0xBA	&uacute;	u acute
// 251	00fb	Û	373	0xFB	0xC3,0xBB	&ucirc;	u circumflex
// 252	00fc	ü	374	0xFC	0xC3,0xBC	&uuml;	u diaeresis
// 253	00fd	ý ý	375	0xFD	0xC3,0xBD	&yacute;	y acute
// 254	00fe	þ þ	376	0xFE	0xC3,0xBE	&thorn;	thorn
// 255	00ff	ÿ	377	0xFF	0xC3,0xBF	&yuml;	y diaeresis
// 338	0152	Œ Œ			0xC5,0x92	&OElig;	Capital ligature OE
// 339	0153	œ œ			0xC5,0x93	&oelig;	Ligature oe
// 352	0160	Š Š			0xC5,0xA0	&Scaron;	Capital S caron
// 353	0161	š š			0xC5,0xA1	&scaron;	s caron
// 376	0178	Ÿ Ÿ			0xC5,0xB8	&Yuml;	Capital Y diaeresis
// 402	0192	ƒ ƒ			0xC6,0x92	&fnof;	F hook, function, florin
// 710	02c6	ˆ			0xCB,0x86	&circ;	Modifier letter circumflex accent
// 732	02dc	˜			0xCB,0x9C	&tilde;	Small tilde
// 913	0391	Α Α			0xCE,0x91	&Alpha;	Capital Alpha
// 914	0392	Β Β			0xCE,0x92	&Beta;	Capital Beta
// 915	0393	Γ Γ			0xCE,0x93	&Gamma;	Capital Gamma
// 916	0394	Δ Δ			0xCE,0x94	&Delta;	Capital Delta
// 917	0395	Ε Ε			0xCE,0x95	&Epsilon;	Capital Epsilon
// 918	0396	Ζ Ζ			0xCE,0x96	&Zeta;	Capital Zeta
// 919	0397	Η Η			0xCE,0x97	&Eta;	Capital Eta
// 920	0398	Θ Θ			0xCE,0x98	&Theta;	Capital Theta
// 921	0399	Ι Ι			0xCE,0x99	&Iota;	Capital Iota
// 922	039a	Κ			0xCE,0x9A	&Kappa;	Capital Kappa
// 923	039b	Λ Λ			0xCE,0x9B	&Lambda;	Capital Lambda
// 924	039c	Μ Μ			0xCE,0x9C	&Mu;	Capital Mu
// 925	039d	Ν Ν			0xCE,0x9D	&Nu;	Capital Nu
// 926	039e	Ξ Ξ			0xCE,0x9E	&Xi;	Capital Xi
// 927	039f	Ο Ο			0xCE,0x9F	&Omicron;	Capital Omicron
// 928	03a0	Π Π			0xCE,0xA0	&Pi;	Capital Pi
// 929	03a1	Ρ Ρ			0xCE,0xA1	&Rho;	Capital Rho
// 931	03a3	Σ Σ			0xCE,0xA3	&Sigma;	Capital Sigma
// 932	03a4	Τ Τ			0xCE,0xA4	&Tau;	Capital Tau
// 933	03a5	Υ Υ			0xCE,0xA5	&Upsilon;	Capital Upsilon
// 934	03a6	Φ Φ			0xCE,0xA6	&Phi;	Capital Phi
// 935	03a7	Χ Χ			0xCE,0xA7	&Chi;	Capital Chi
// 936	03a8	Ψ Ψ			0xCE,0xA8	&Psi;	Capital Psi
// 937	03a9	Ω Ω			0xCE,0xA9	&Omega;	Capital Omega
// 945	03b1	α α			0xCE,0xB1	&alpha;	alpha
// 946	03b2	β β			0xCE,0xB2	&beta;	beta
// 947	03b3	γ γ			0xCE,0xB3	&gamma;	gamma
// 948	03b4	δ δ			0xCE,0xB4	&delta;	delta
// 949	03b5	ε ε			0xCE,0xB5	&epsilon;	epsilon
// 950	03b6	ζ ζ			0xCE,0xB6	&zeta;	zeta
// 951	03b7	η η			0xCE,0xB7	&eta;	eta
// 952	03b8	θ θ			0xCE,0xB8	&theta;	theta
// 953	03b9	ι ι			0xCE,0xB9	&iota;	iota
// 954	03ba	κ κ			0xCE,0xBA	&kappa;	kappa
// 955	03bb	λ λ			0xCE,0xBB	&lambda;	lambda
// 956	03bc	μ			0xCE,0xBC	&mu;	mu
// 957	03bd	ν ν			0xCE,0xBD	&nu;	nu
// 958	03be	ξ ξ			0xCE,0xBE	&xi;	xi
// 959	03bf	ο ο			0xCE,0xBF	&omicron;	omicron
// 960	03c0	π π			0xCF,0x80	&pi;	pi
// 961	03c1	ρ ρ			0xCF,0x81	&rho;	rho
// 962	03c2	ς ς			0xCF,0x82	&sigmaf;	final sigma
// 963	03c3	σ σ			0xCF,0x83	&sigma;	sigma
// 964	03c4	τ τ			0xCF,0x84	&tau;	tau
// 965	03c5	υ υ			0xCF,0x85	&upsilon;	upsilon
// 966	03c6	φ φ			0xCF,0x86	&phi;	phi
// 967	03c7	χ χ			0xCF,0x87	&chi;	chi
// 968	03c8	ψ ψ			0xCF,0x88	&psi;	psi
// 969	03c9	ω ω			0xCF,0x89	&omega;	omega
// 977	03d1	ϑ ϑ			0xCF,0x91	&thetasym;	theta symbol
// 978	03d2	ϒ ϒ			0xCF,0x92	&upsih;	Greek Upsilon with hook symbol
// 982	03d6	ϖ ϖ			0xCF,0x96	&piv;	Greek pi symbol
// 8194	2002	 			0xE2,0x80,0x82	&ensp;	En space
// 8195	2003	 			0xE2,0x80,0x83	&emsp;	Em space
// 8201	2009	 			0xE2,0x80,0x89	&thinsp;	Thin space
// 8204	200c				0xE2,0x80,0x8C	&zwnj;	Zero width non-joiner
// 8205	200d				0xE2,0x80,0x8D	&zwj;	Zero width joiner
// 8206	200e				0xE2,0x80,0x8E	&lrm;	Left-to-right mark
// 8207	200f				0xE2,0x80,0x8F	&rlm;	Right-to-left mark
// 8211	2013	–			0xE2,0x80,0x93	&ndash;	En dash
// 8212	2014	—			0xE2,0x80,0x94	&mdash;	Em dash
// 8216	2018	‘			0xE2,0x80,0x98	&lsquo;	Left single quotation mark
// 8217	2019	’			0xE2,0x80,0x99	&rsquo;	Right single quotation mark
// 8218	201a	‚ ‚			0xE2,0x80,0x9A	&sbquo;	Single low-9 quotation mark
// 8220	201c	“			0xE2,0x80,0x9C	&ldquo;	Left double quotation mark
// 8221	201d	”			0xE2,0x80,0x9D	&rdquo;	Right double quotation mark
// 8222	201e	„ „			0xE2,0x80,0x9E	&bdquo;	Double low-9 quotation mark
// 8224	2020	† †			0xE2,0x80,0xA0	&dagger;	Dagger
// 8225	2021	‡ ‡			0xE2,0x80,0xA1	&Dagger;	Double dagger
// 8226	2022	·			0xE2,0x80,0xA2	&bull;	Bullet, black small circle
// 8230	2026	…			0xE2,0x80,0xA6	&hellip;	Horizontal ellipsis, three dot leader
// 8240	2030	‰ ‰			0xE2,0x80,0xB0	&permil;	Per mille sign
// 8242	2032	′			0xE2,0x80,0xB2	&prime;	Prime, minutes, feet
// 8243	2033	″ ″			0xE2,0x80,0xB3	&Prime;	Double prime, seconds, inches
// 8249	2039	‹ ‹			0xE2,0x80,0xB9	&lsaquo;	Single left-pointing angle quotation mark
// 8250	203a	› ›			0xE2,0x80,0xBA	&rsaquo;	Single right-pointing angle quotation mark
// 8254	203e	‾			0xE2,0x80,0xBE	&oline;	Overline, spacing overscore
// 8260	2044	⁄			0xE2,0x81,0x84	&frasl;	Fraction slash
// 8364	20ac	€ €			0xE2,0x82,0xAC	&euro;	Euro sign
// 8465	2111	ℑ ℑ			0xE2,0x84,0x91	&image;	Blackletter capital I, imaginary part
// 8472	2118	℘ ℘			0xE2,0x84,0x98	&weierp;	Script capital P, power set, Weierstrass p
// 8476	211c	ℜ ℜ			0xE2,0x84,0x9C	&real;	Blackletter capital R, real part symbol
// 8482	2122	™ ™			0xE2,0x84,0xA2	&trade;	Trademark sign
// 8501	2135	ℵ ℵ			0xE2,0x84,0xB5	&alefsym;	Alef symbol, first transfinite cardinal
// 8592	2190	← ←			0xE2,0x86,0x90	&larr;	Leftward arrow
// 8593	2191	↑ ↑			0xE2,0x86,0x91	&uarr;	Upward arrow
// 8594	2192	→ →			0xE2,0x86,0x92	&rarr;	Rightward arrow
// 8595	2193	↓ ↓			0xE2,0x86,0x93	&darr;	Downward arrow
// 8596	2194	↔ ↔			0xE2,0x86,0x94	&harr;	Left-right arrow
// 8629	21b5	↵ ↵			0xE2,0x86,0xB5	&crarr;	Downward arrow with corner leftward, carriage return
// 8656	21d0	⇐ ⇐			0xE2,0x87,0x90	&lArr;	Leftward double arrow
// 8657	21d1	⇑ ⇑			0xE2,0x87,0x91	&uArr;	Upward double arrow
// 8658	21d2	⇒ ⇒			0xE2,0x87,0x92	&rArr;	Rightward double arrow
// 8659	21d3	⇓ ⇓			0xE2,0x87,0x93	&dArr;	Downward double arrow
// 8660	21d4	⇔ ⇔			0xE2,0x87,0x94	&hArr;	Left-right double arrow
// 8704	2200	∀ ∀			0xE2,0x88,0x80	&forall;	For all
// 8706	2202	∂ ∂			0xE2,0x88,0x82	&part;	Partial differential
// 8707	2203	∃ ∃			0xE2,0x88,0x83	&exist;	There exists
// 8709	2205	∅			0xE2,0x88,0x85	&empty;	Empty set, null set, diameter
// 8711	2207	∇ ∇			0xE2,0x88,0x87	&nabla;	Nabla, backward difference
// 8712	2208	∈ ∈			0xE2,0x88,0x88	&isin;	Element of
// 8713	2209	∉ ∉			0xE2,0x88,0x89	&notin;	Not an element of
// 8715	220b	∋ ∋			0xE2,0x88,0x8B	&ni;	Contains as member
// 8719	220f	∏ ∏			0xE2,0x88,0x8F	&prod;	n-ary product, product sign
// 8721	2211	∑ ∑			0xE2,0x88,0x91	&sum;	n-ary sumation
// 8722	2212	−			0xE2,0x88,0x92	&minus;	Minus sign
// 8727	2217	∗			0xE2,0x88,0x97	&lowast;	Asterisk operator
// 8730	221a	√ √			0xE2,0x88,0x9A	&radic;	Square root, radical sign
// 8733	221d	∝ ∝			0xE2,0x88,0x9D	&prop;	Proportional to
// 8734	221e	∞ ∞			0xE2,0x88,0x9E	&infin;	Infinity
// 8736	2220	∠ ∠			0xE2,0x88,0xA0	&ang;	Angle
// 8743	2227	∧ ∧			0xE2,0x88,0xA7	&and;	Logical and, wedge
// 8744	2228	∨ ∨			0xE2,0x88,0xA8	&or;	Logical or, vee
// 8745	2229	∩ ∩			0xE2,0x88,0xA9	&cap;	Intersection, cap
// 8746	222a	∪ ∪			0xE2,0x88,0xAA	&cup;	Union, cup
// 8747	222b	∫ ∫			0xE2,0x88,0xAB	&int;	Integral
// 8756	2234	∴ ∴			0xE2,0x88,0xB4	&there4;	Therefore
// 8764	223c	∼ ∼			0xE2,0x88,0xBC	&sim;	Tilde operator, varies with, similar to
// 8773	2245	≅ ≅			0xE2,0x89,0x85	&cong;	Approximately equal to
// 8776	2248	≍ ≍			0xE2,0x89,0x88	&asymp;	Almost equal to, asymptotic to
// 8800	2260	≠ ≠			0xE2,0x89,0xA0	&ne;	Not equal to
// 8801	2261	≡ ≡			0xE2,0x89,0xA1	&equiv;	Identical to
// 8804	2264	≤ ≤			0xE2,0x89,0xA4	&le;	Less-than or equal to
// 8805	2265	≥ ≥			0xE2,0x89,0xA5	&ge;	Greater-than or equal to
// 8834	2282	⊂ ⊂			0xE2,0x8A,0x82	&sub;	Subset of
// 8835	2283	⊃ ⊃			0xE2,0x8A,0x83	&sup;	Superset of
// 8836	2284	⊄ ⊄			0xE2,0x8A,0x84	&nsub;	Not a subset of
// 8838	2286	⊆ ⊆			0xE2,0x8A,0x86	&sube;	Subset of or equal to
// 8839	2287	⊇ ⊇			0xE2,0x8A,0x87	&supe;	Superset of or equal to
// 8853	2295	⊕ ⊕			0xE2,0x8A,0x95	&oplus;	Circled plus, direct sum
// 8855	2297	⊗ ⊗			0xE2,0x8A,0x97	&otimes;	Circled times, vector product
// 8869	22a5	⊥ ⊥			0xE2,0x8A,0xA5	&perp;	Up tack, orthogonal to, perpendicular
// 8901	22c5	⋅ ⋅			0xE2,0x8B,0x85	&sdot;	Dot operator
// 8968	2308	⌈ ⌈			0xE2,0x8C,0x88	&lceil;	Left ceiling, APL upstile
// 8969	2309	⌉ ⌉			0xE2,0x8C,0x89	&rceil;	Right ceiling
// 8970	230a	⌊ ⌊			0xE2,0x8C,0x8A	&lfloor;	Left floor, APL downstile
// 8971	230b	⌋ ⌋			0xE2,0x8C,0x8B	&rfloor;	Right floor
// 9001	2329	〈 〈			0xE2,0x8C,0xA9	&lang;	Left-pointing angle bracket, bra
// 9002	232a	〉 〉			0xE2,0x8C,0xAA	&rang;	Right-pointing angle bracket, ket
// 9674	25ca	◊ ◊			0xE2,0x97,0x8A	&loz;	Lozenge
// 9824	2660	♠ ♠			0xE2,0x99,0xA0	&spades;	Black spade suit
// 9827	2663	♣ ♣			0xE2,0x99,0xA3	&clubs;	Black club suit, shamrock
// 9829	2665	♥ ♥			0xE2,0x99,0xA5	&hearts;	Black heart suit, valentine
// 9830	2666	♦ ♦			0xE2,0x99,0xA6	&diams;	Black diamond suit
