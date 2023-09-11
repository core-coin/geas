// Code generated by "stringer -type tokenType"; DO NOT EDIT.

package asm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[eof-0]
	_ = x[lineStart-1]
	_ = x[lineEnd-2]
	_ = x[invalidToken-3]
	_ = x[identifier-4]
	_ = x[dottedIdentifier-5]
	_ = x[labelRef-6]
	_ = x[dottedLabelRef-7]
	_ = x[label-8]
	_ = x[dottedLabel-9]
	_ = x[numberLiteral-10]
	_ = x[stringLiteral-11]
	_ = x[openParen-12]
	_ = x[closeParen-13]
	_ = x[comma-14]
	_ = x[arithPlus-15]
	_ = x[arithMinus-16]
	_ = x[arithMul-17]
	_ = x[arithDiv-18]
	_ = x[arithMod-19]
	_ = x[arithLshift-20]
	_ = x[arithRshift-21]
	_ = x[arithAnd-22]
	_ = x[arithOr-23]
	_ = x[arithHat-24]
	_ = x[directive-25]
	_ = x[instMacroIdent-26]
	_ = x[openBrace-27]
	_ = x[closeBrace-28]
}

const _tokenType_name = "eoflineStartlineEndinvalidTokenidentifierdottedIdentifierlabelRefdottedLabelReflabeldottedLabelnumberLiteralstringLiteralopenParencloseParencommaarithPlusarithMinusarithMularithDivarithModarithLshiftarithRshiftarithAndarithOrarithHatdirectiveinstMacroIdentopenBracecloseBrace"

var _tokenType_index = [...]uint16{0, 3, 12, 19, 31, 41, 57, 65, 79, 84, 95, 108, 121, 130, 140, 145, 154, 164, 172, 180, 188, 199, 210, 218, 225, 233, 242, 256, 265, 275}

func (i tokenType) String() string {
	if i >= tokenType(len(_tokenType_index)-1) {
		return "tokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _tokenType_name[_tokenType_index[i]:_tokenType_index[i+1]]
}