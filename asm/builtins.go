// Copyright 2023 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package asm

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"

	"github.com/core-coin/go-core/v2/common"
	"github.com/core-coin/go-core/v2/signer/fourbyte"
	"golang.org/x/crypto/sha3"
)

var builtinMacros = map[string]builtinMacroFn{
	"bitlen":    bitlenMacro,
	"bytelen":   bytelenMacro,
	"abs":       absMacro,
	"address":   addressMacro,
	"selector":  selectorMacro,
	"keccak256": keccak256Macro,
	"sha256":    sha256Macro,
}

type builtinMacroFn func(*evaluator, *evalEnvironment, *macroCallExpr) (*big.Int, error)

func bitlenMacro(e *evaluator, env *evalEnvironment, call *macroCallExpr) (*big.Int, error) {
	if err := call.checkArgCount(1); err != nil {
		return nil, err
	}
	v, err := call.args[0].eval(e, env)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(v.BitLen())), nil
}

func bytelenMacro(e *evaluator, env *evalEnvironment, call *macroCallExpr) (*big.Int, error) {
	if err := call.checkArgCount(1); err != nil {
		return nil, err
	}
	v, err := call.args[0].eval(e, env)
	if err != nil {
		return nil, err
	}
	bytes := (v.BitLen() + 7) / 8
	return big.NewInt(int64(bytes)), nil
}

func absMacro(e *evaluator, env *evalEnvironment, call *macroCallExpr) (*big.Int, error) {
	if err := call.checkArgCount(1); err != nil {
		return nil, err
	}
	v, err := call.args[0].eval(e, env)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Abs(v), nil
}

func sha256Macro(e *evaluator, env *evalEnvironment, call *macroCallExpr) (*big.Int, error) {
	if err := call.checkArgCount(1); err != nil {
		return nil, err
	}
	bytes, err := evalAsBytes(e, env, call.args[0])
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(bytes)
	return new(big.Int).SetBytes(hash[:]), nil
}

func keccak256Macro(e *evaluator, env *evalEnvironment, call *macroCallExpr) (*big.Int, error) {
	if err := call.checkArgCount(1); err != nil {
		return nil, err
	}
	bytes, err := evalAsBytes(e, env, call.args[0])
	if err != nil {
		return nil, err
	}
	w := sha3.New256()
	w.Write(bytes)
	hash := w.Sum(nil)
	return new(big.Int).SetBytes(hash[:]), nil
}

var (
	errSelectorWantsLiteral = fmt.Errorf(".selector(...) requires literal string argument")
)

func selectorMacro(e *evaluator, env *evalEnvironment, call *macroCallExpr) (*big.Int, error) {
	if err := call.checkArgCount(1); err != nil {
		return nil, err
	}
	lit, ok := call.args[0].(*literalExpr)
	if !ok || lit.tok.typ != stringLiteral {
		return nil, errSelectorWantsLiteral
	}
	text := lit.tok.text
	if _, err := fourbyte.ParseSelector(text); err != nil {
		return nil, fmt.Errorf("invalid ABI selector")
	}
	w := sha3.New256()
	w.Write([]byte(text))
	hash := w.Sum(nil)
	return new(big.Int).SetBytes(hash[:4]), nil
}

var (
	errAddressWantsLiteral = errors.New(".address(...) requires literal argument")
	errAddressInvalid      = errors.New("invalid Core Blockchain address")
)

func addressMacro(e *evaluator, env *evalEnvironment, call *macroCallExpr) (*big.Int, error) {
	if err := call.checkArgCount(1); err != nil {
		return nil, err
	}
	lit, ok := call.args[0].(*literalExpr)
	if !ok {
		return nil, errAddressWantsLiteral
	}
	text := lit.tok.text
	addr, err := common.HexToAddress(text)
	if err != nil {
		return nil, errAddressInvalid
	}

	return new(big.Int).SetBytes(addr[:]), nil
}

// evalAsBytes gives the byte value of an expression.
// In most cases, this is just the big-endian encoding of an integer value.
// For string literals, the bytes of the literal are used directly.
func evalAsBytes(e *evaluator, env *evalEnvironment, expr astExpr) ([]byte, error) {
	lit, ok := expr.(*literalExpr)
	if ok && lit.tok.typ == stringLiteral {
		return []byte(lit.tok.text), nil
	}
	v, err := expr.eval(e, env)
	if err != nil {
		return nil, err
	}
	return v.Bytes(), nil
}