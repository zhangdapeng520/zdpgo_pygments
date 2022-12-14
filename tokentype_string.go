package zdpgo_pygments

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Background - -1]
	_ = x[PreWrapper - -2]
	_ = x[Line - -3]
	_ = x[LineNumbers - -4]
	_ = x[LineNumbersTable - -5]
	_ = x[LineHighlight - -6]
	_ = x[LineTable - -7]
	_ = x[LineTableTD - -8]
	_ = x[CodeLine - -9]
	_ = x[Error - -10]
	_ = x[Other - -11]
	_ = x[None - -12]
	_ = x[EOFType-0]
	_ = x[Keyword-1000]
	_ = x[KeywordConstant-1001]
	_ = x[KeywordDeclaration-1002]
	_ = x[KeywordNamespace-1003]
	_ = x[KeywordPseudo-1004]
	_ = x[KeywordReserved-1005]
	_ = x[KeywordType-1006]
	_ = x[Name-2000]
	_ = x[NameAttribute-2001]
	_ = x[NameBuiltin-2002]
	_ = x[NameBuiltinPseudo-2003]
	_ = x[NameClass-2004]
	_ = x[NameConstant-2005]
	_ = x[NameDecorator-2006]
	_ = x[NameEntity-2007]
	_ = x[NameException-2008]
	_ = x[NameFunction-2009]
	_ = x[NameFunctionMagic-2010]
	_ = x[NameKeyword-2011]
	_ = x[NameLabel-2012]
	_ = x[NameNamespace-2013]
	_ = x[NameOperator-2014]
	_ = x[NameOther-2015]
	_ = x[NamePseudo-2016]
	_ = x[NameProperty-2017]
	_ = x[NameTag-2018]
	_ = x[NameVariable-2019]
	_ = x[NameVariableAnonymous-2020]
	_ = x[NameVariableClass-2021]
	_ = x[NameVariableGlobal-2022]
	_ = x[NameVariableInstance-2023]
	_ = x[NameVariableMagic-2024]
	_ = x[Literal-3000]
	_ = x[LiteralDate-3001]
	_ = x[LiteralOther-3002]
	_ = x[LiteralString-3100]
	_ = x[LiteralStringAffix-3101]
	_ = x[LiteralStringAtom-3102]
	_ = x[LiteralStringBacktick-3103]
	_ = x[LiteralStringBoolean-3104]
	_ = x[LiteralStringChar-3105]
	_ = x[LiteralStringDelimiter-3106]
	_ = x[LiteralStringDoc-3107]
	_ = x[LiteralStringDouble-3108]
	_ = x[LiteralStringEscape-3109]
	_ = x[LiteralStringHeredoc-3110]
	_ = x[LiteralStringInterpol-3111]
	_ = x[LiteralStringName-3112]
	_ = x[LiteralStringOther-3113]
	_ = x[LiteralStringRegex-3114]
	_ = x[LiteralStringSingle-3115]
	_ = x[LiteralStringSymbol-3116]
	_ = x[LiteralNumber-3200]
	_ = x[LiteralNumberBin-3201]
	_ = x[LiteralNumberFloat-3202]
	_ = x[LiteralNumberHex-3203]
	_ = x[LiteralNumberInteger-3204]
	_ = x[LiteralNumberIntegerLong-3205]
	_ = x[LiteralNumberOct-3206]
	_ = x[Operator-4000]
	_ = x[OperatorWord-4001]
	_ = x[Punctuation-5000]
	_ = x[Comment-6000]
	_ = x[CommentHashbang-6001]
	_ = x[CommentMultiline-6002]
	_ = x[CommentSingle-6003]
	_ = x[CommentSpecial-6004]
	_ = x[CommentPreproc-6100]
	_ = x[CommentPreprocFile-6101]
	_ = x[Generic-7000]
	_ = x[GenericDeleted-7001]
	_ = x[GenericEmph-7002]
	_ = x[GenericError-7003]
	_ = x[GenericHeading-7004]
	_ = x[GenericInserted-7005]
	_ = x[GenericOutput-7006]
	_ = x[GenericPrompt-7007]
	_ = x[GenericStrong-7008]
	_ = x[GenericSubheading-7009]
	_ = x[GenericTraceback-7010]
	_ = x[GenericUnderline-7011]
	_ = x[Text-8000]
	_ = x[TextWhitespace-8001]
	_ = x[TextSymbol-8002]
	_ = x[TextPunctuation-8003]
}

const _TokenType_name = "NoneOtherErrorCodeLineLineTableTDLineTableLineHighlightLineNumbersTableLineNumbersLinePreWrapperBackgroundEOFTypeKeywordKeywordConstantKeywordDeclarationKeywordNamespaceKeywordPseudoKeywordReservedKeywordTypeNameNameAttributeNameBuiltinNameBuiltinPseudoNameClassNameConstantNameDecoratorNameEntityNameExceptionNameFunctionNameFunctionMagicNameKeywordNameLabelNameNamespaceNameOperatorNameOtherNamePseudoNamePropertyNameTagNameVariableNameVariableAnonymousNameVariableClassNameVariableGlobalNameVariableInstanceNameVariableMagicLiteralLiteralDateLiteralOtherLiteralStringLiteralStringAffixLiteralStringAtomLiteralStringBacktickLiteralStringBooleanLiteralStringCharLiteralStringDelimiterLiteralStringDocLiteralStringDoubleLiteralStringEscapeLiteralStringHeredocLiteralStringInterpolLiteralStringNameLiteralStringOtherLiteralStringRegexLiteralStringSingleLiteralStringSymbolLiteralNumberLiteralNumberBinLiteralNumberFloatLiteralNumberHexLiteralNumberIntegerLiteralNumberIntegerLongLiteralNumberOctOperatorOperatorWordPunctuationCommentCommentHashbangCommentMultilineCommentSingleCommentSpecialCommentPreprocCommentPreprocFileGenericGenericDeletedGenericEmphGenericErrorGenericHeadingGenericInsertedGenericOutputGenericPromptGenericStrongGenericSubheadingGenericTracebackGenericUnderlineTextTextWhitespaceTextSymbolTextPunctuation"

var _TokenType_map = map[TokenType]string{
	-12:  _TokenType_name[0:4],
	-11:  _TokenType_name[4:9],
	-10:  _TokenType_name[9:14],
	-9:   _TokenType_name[14:22],
	-8:   _TokenType_name[22:33],
	-7:   _TokenType_name[33:42],
	-6:   _TokenType_name[42:55],
	-5:   _TokenType_name[55:71],
	-4:   _TokenType_name[71:82],
	-3:   _TokenType_name[82:86],
	-2:   _TokenType_name[86:96],
	-1:   _TokenType_name[96:106],
	0:    _TokenType_name[106:113],
	1000: _TokenType_name[113:120],
	1001: _TokenType_name[120:135],
	1002: _TokenType_name[135:153],
	1003: _TokenType_name[153:169],
	1004: _TokenType_name[169:182],
	1005: _TokenType_name[182:197],
	1006: _TokenType_name[197:208],
	2000: _TokenType_name[208:212],
	2001: _TokenType_name[212:225],
	2002: _TokenType_name[225:236],
	2003: _TokenType_name[236:253],
	2004: _TokenType_name[253:262],
	2005: _TokenType_name[262:274],
	2006: _TokenType_name[274:287],
	2007: _TokenType_name[287:297],
	2008: _TokenType_name[297:310],
	2009: _TokenType_name[310:322],
	2010: _TokenType_name[322:339],
	2011: _TokenType_name[339:350],
	2012: _TokenType_name[350:359],
	2013: _TokenType_name[359:372],
	2014: _TokenType_name[372:384],
	2015: _TokenType_name[384:393],
	2016: _TokenType_name[393:403],
	2017: _TokenType_name[403:415],
	2018: _TokenType_name[415:422],
	2019: _TokenType_name[422:434],
	2020: _TokenType_name[434:455],
	2021: _TokenType_name[455:472],
	2022: _TokenType_name[472:490],
	2023: _TokenType_name[490:510],
	2024: _TokenType_name[510:527],
	3000: _TokenType_name[527:534],
	3001: _TokenType_name[534:545],
	3002: _TokenType_name[545:557],
	3100: _TokenType_name[557:570],
	3101: _TokenType_name[570:588],
	3102: _TokenType_name[588:605],
	3103: _TokenType_name[605:626],
	3104: _TokenType_name[626:646],
	3105: _TokenType_name[646:663],
	3106: _TokenType_name[663:685],
	3107: _TokenType_name[685:701],
	3108: _TokenType_name[701:720],
	3109: _TokenType_name[720:739],
	3110: _TokenType_name[739:759],
	3111: _TokenType_name[759:780],
	3112: _TokenType_name[780:797],
	3113: _TokenType_name[797:815],
	3114: _TokenType_name[815:833],
	3115: _TokenType_name[833:852],
	3116: _TokenType_name[852:871],
	3200: _TokenType_name[871:884],
	3201: _TokenType_name[884:900],
	3202: _TokenType_name[900:918],
	3203: _TokenType_name[918:934],
	3204: _TokenType_name[934:954],
	3205: _TokenType_name[954:978],
	3206: _TokenType_name[978:994],
	4000: _TokenType_name[994:1002],
	4001: _TokenType_name[1002:1014],
	5000: _TokenType_name[1014:1025],
	6000: _TokenType_name[1025:1032],
	6001: _TokenType_name[1032:1047],
	6002: _TokenType_name[1047:1063],
	6003: _TokenType_name[1063:1076],
	6004: _TokenType_name[1076:1090],
	6100: _TokenType_name[1090:1104],
	6101: _TokenType_name[1104:1122],
	7000: _TokenType_name[1122:1129],
	7001: _TokenType_name[1129:1143],
	7002: _TokenType_name[1143:1154],
	7003: _TokenType_name[1154:1166],
	7004: _TokenType_name[1166:1180],
	7005: _TokenType_name[1180:1195],
	7006: _TokenType_name[1195:1208],
	7007: _TokenType_name[1208:1221],
	7008: _TokenType_name[1221:1234],
	7009: _TokenType_name[1234:1251],
	7010: _TokenType_name[1251:1267],
	7011: _TokenType_name[1267:1283],
	8000: _TokenType_name[1283:1287],
	8001: _TokenType_name[1287:1301],
	8002: _TokenType_name[1301:1311],
	8003: _TokenType_name[1311:1326],
}

func (i TokenType) String() string {
	if str, ok := _TokenType_map[i]; ok {
		return str
	}
	return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
}
