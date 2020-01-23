// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.

package confp

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/params/types/ctypes"
	"github.com/fatih/camelcase"
	"github.com/go-openapi/spec"
	openRPCTypes "github.com/gregdhill/go-openrpc/types"
	"github.com/iancoleman/strcase"
)

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *uint64
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func NewCompatError(what string, storedblock, newblock *uint64) *ConfigCompatError {
	var rew *uint64
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || *storedblock < *newblock:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && *rew > 0 {
		err.RewindTo = *rew - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	var have, want interface{}
	if err.StoredConfig != nil {
		have = *err.StoredConfig
	}
	if err.NewConfig != nil {
		want = *err.NewConfig
	}
	if have == nil {
		have = "nil"
	}
	if want == nil {
		want = "nil"
	}
	return fmt.Sprintf("mismatching %s in database (old: %v, new: %v, rewindto %d)", err.What, have, want, err.RewindTo)
}

type ConfigValidError struct {
	What string
	A, B interface{}
}

func NewValidErr(what string, a, b interface{}) *ConfigValidError {
	return &ConfigValidError{
		What: what,
		A:    a,
		B:    b,
	}
}

func (err *ConfigValidError) Error() string {
	return fmt.Sprintf("%s, %v/%v", err.What, err.A, err.B)
}

func IsEmpty(anything interface{}) bool {
	if anything == nil {
		return true
	}
	return reflect.DeepEqual(anything, reflect.Zero(reflect.TypeOf(anything)).Interface())
}

type DiscoverMethodParameterItem struct {
	Index   int
	Name    string
	Kind    string
	Nilable bool
}
type DiscoverMethodItem struct {
	Base       string
	Name       string
	NumParams  int
	NumReturns int
	Params     []DiscoverMethodParameterItem
	Returns    []DiscoverMethodParameterItem
}

func (i DiscoverMethodItem) AsWeb3Ext() string {
	s := fmt.Sprintf(`new web3._extend.Method({
	name: '%s',
	call: '%s_%s',
	params: %d,
}),`, strcase.ToLowerCamel(i.Name),
		i.Base, strcase.ToLowerCamel(i.Name),
		i.NumParams, // -1 b/c receiver methods
	)
	return s
}

func Discover(conf ctypes.ChainConfigurator) (items []DiscoverMethodItem) {
	k := reflect.TypeOf(conf)
	for i := 0; i < k.NumMethod(); i++ {
		method := k.Method(i)

		if method.Name == "String" {
			continue
		}

		tm := method.Type // func

		numIn := tm.NumIn()
		numOut := tm.NumOut()
		item := DiscoverMethodItem{
			Base:       "chainconfig",
			Name:       method.Name,
			NumParams:  numIn - 1,
			NumReturns: numOut,
			Params:     nil,
			Returns:    nil,
		}

		if numIn > 0 {
			item.Params = []DiscoverMethodParameterItem{}
		}
		if numOut > 0 {
			item.Returns = []DiscoverMethodParameterItem{}
		}
		for i := 1; i < numIn; i++ {
			inV := tm.In(i)
			inKind := inV.Kind()
			inName := inV.Name()
			nilable := false
			if inKind == reflect.Ptr {
				inKind = inV.Elem().Kind()
				inName = inV.Elem().Name()
				nilable = true
			}
			it := DiscoverMethodParameterItem{
				Index:   i,
				Name:    inName,
				Kind:    fmt.Sprintf("%v", inKind),
				Nilable: nilable,
			}
			item.Params = append(item.Params, it)
		}
		for i := 0; i < numOut; i++ {
			outV := tm.Out(i)
			outKind := outV.Kind()
			outName := outV.Name()
			nilable := false
			if outKind == reflect.Ptr {
				outKind = outV.Elem().Kind()
				outName = outV.Elem().Name()
				nilable = true
			}
			it := DiscoverMethodParameterItem{
				Index:   i,
				Name:    outName,
				Kind:    fmt.Sprintf("%v", outKind),
				Nilable: nilable,
			}
			item.Returns = append(item.Returns, it)
		}

		items = append(items, item)
	}
	return
}

func DiscoverOpenRPC(conf ctypes.ChainConfigurator) (*openRPCTypes.OpenRPCSpec1) {
	openrpcSpec := openRPCTypes.NewOpenRPCSpec1()
	its := Discover(conf)
	
	for _, it := range its {
		
		var params = []*openRPCTypes.ContentDescriptor{}
		for _, p := range it.Params {
			schemaName := "blockNumber"
			if !strings.HasSuffix(it.Name, "Transition") {
				strs := camelcase.Split(it.Name)
				schemaName = strs[len(strs)-1]
				schemaName = strcase.ToLowerCamel(schemaName)

			}
			param := &openRPCTypes.ContentDescriptor{
				Content: openRPCTypes.Content{
					Name:        schemaName,
					Summary:     "",
					Description: "",
					Required:    p.Nilable,
					Deprecated:  false,
					Schema:      spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type:                 spec.StringOrArray{p.Name},
							Title:                schemaName,
						},
					},
				},
			}
			params = append(params, param)
		}

		var result = &openRPCTypes.ContentDescriptor{}
		for _, r := range it.Returns {
			if r.Name == "error" {
				continue
			}
			schemaName := "blockNumber"
			if !strings.HasSuffix(it.Name, "Transition") {
				strs := camelcase.Split(it.Name)
				schemaName = strs[len(strs)-1]
				schemaName = strcase.ToLowerCamel(schemaName)
			}
			result.Name = schemaName
			result.Schema = spec.Schema{
				SchemaProps:        spec.SchemaProps{
					Type: spec.StringOrArray{r.Name},
					Title: schemaName,
				},
			}
		}

		if result.Name == "" {
			result = nil
		}
		method := openRPCTypes.Method{
			Name:           it.Name,
			Tags:           nil,
			Summary:        "",
			Description:    "",
			ExternalDocs:   openRPCTypes.ExternalDocs{},
			Params:         params,
			Result:         result,
			Deprecated:     false,
			Servers:        nil,
			Errors:         nil,
			Links:          nil,
			ParamStructure: "",
			Examples:       nil,
		}
		openrpcSpec.Methods = append(openrpcSpec.Methods, method)
	}
	return openrpcSpec
}

func IsValid(conf ctypes.ChainConfigurator, head *uint64) *ConfigValidError {

	// head-agnostic logic
	if conf.GetNetworkID() == nil || *conf.GetNetworkID() == 0 {
		return NewValidErr("NetworkID cannot be empty nor zero", ">=0", conf.GetNetworkID())
	}
	if head == nil {
		return nil
	}

	// head-full logic
	var bhead = new(big.Int).SetUint64(*head)

	if conf.IsForked(conf.GetEIP155Transition, bhead) && conf.GetChainID() == nil {
		return NewValidErr("EIP155 requires ChainID. A:EIP155/B:ChainID", conf.GetEIP155Transition(), conf.GetChainID())
	}

	return nil
}

func Compatible(head *uint64, a, b ctypes.ChainConfigurator) *ConfigCompatError {
	// Iterate checkCompatible to find the lowest conflict.
	var lastErr *ConfigCompatError
	for {
		err := compatible(head, a, b)
		if err == nil || (lastErr != nil && err.RewindTo == lastErr.RewindTo) {
			break
		}
		lastErr = err
		head = &err.RewindTo
	}
	return lastErr
}

func compatible(head *uint64, a, b ctypes.ChainConfigurator) *ConfigCompatError {
	aFns, aNames := Transitions(a)
	bFns, _ := Transitions(b)
	for i, afn := range aFns {
		if err := func(c1, c2, head *uint64) *ConfigCompatError {
			if isForkIncompatible(c1, c2, head) {
				return NewCompatError("incompatible fork value: "+aNames[i], c1, c2)
			}
			return nil
		}(afn(), bFns[i](), head); err != nil {
			return err
		}
	}
	if head == nil {
		return nil
	}
	if a.IsForked(a.GetEIP155Transition, new(big.Int).SetUint64(*head)) {
		if a.GetChainID().Cmp(b.GetChainID()) != 0 {
			return NewCompatError("mismatching chain ids after EIP155 transition", a.GetEIP155Transition(), b.GetEIP155Transition())
		}
	}

	return nil
}

func Equivalent(a, b ctypes.ChainConfigurator) error {
	if a.GetConsensusEngineType() != b.GetConsensusEngineType() {
		return fmt.Errorf("mismatch consensus engine types, A: %s, B: %s", a.GetConsensusEngineType(), b.GetConsensusEngineType())
	}

	// Check forks sameness.
	fa, fb := Forks(a), Forks(b)
	if len(fa) != len(fb) {
		return fmt.Errorf("different fork count: %d / %d (%v / %v)", len(fa), len(fb), fa, fb)
	}
	for i := range fa {
		if fa[i] != fb[i] {
			if fa[i] == math.MaxUint64 {
				return fmt.Errorf("fa bigmax: %d", fa[i])
			}
			if fb[i] == math.MaxUint64 {
				return fmt.Errorf("fb bigmax: %d", fb[i])
			}
			return fmt.Errorf("fork index %d not same: %d / %d", i, fa[i], fb[i])
		}
	}

	// Check initial, at- and around-fork, and eventual compatibility.
	var testForks = []uint64{}
	copy(testForks, fa)
	// Don't care about dupes.
	for _, f := range fa {
		testForks = append(testForks, f-1)
	}
	testForks = append(testForks, 0, math.MaxUint64)

	// essentiallyEquivalent treats nil and bitsize-max numbers as essentially equivalent.
	essentiallyEquivalent := func(x, y *uint64) bool {
		if x == nil && y != nil {
			return *y == math.MaxUint64 ||
				*y == 0x7FFFFFFFFFFFFFFF ||
				*y == 0x7FFFFFFFFFFFFFF ||
				*y == 0x7FFFFFFFFFFFFF
		}
		if x != nil && y == nil {
			return *x == math.MaxUint64 ||
				*x == 0x7FFFFFFFFFFFFFFF ||
				*x == 0x7FFFFFFFFFFFFFF ||
				*x == 0x7FFFFFFFFFFFFF
		}
		return false
	}
	for _, h := range testForks {
		if err := Compatible(&h, a, b); err != nil {
			if !essentiallyEquivalent(err.StoredConfig, err.NewConfig) {
				return err
			}
		}
	}

	if a.GetConsensusEngineType() == ctypes.ConsensusEngineT_Ethash {
		for _, f := range fa { // fa and fb are fork-equivalent
			ar := ctypes.EthashBlockReward(a, new(big.Int).SetUint64(f))
			br := ctypes.EthashBlockReward(b, new(big.Int).SetUint64(f))
			if ar.Cmp(br) != 0 {
				return fmt.Errorf("mismatch block reward, fork block: %v, A: %v, B: %v", f, ar, br)
			}
			// TODO: add difficulty comparison
			// Currently tough/complex to do because of necessary overhead (ie build a parent block).
		}
	} else if a.GetConsensusEngineType() == ctypes.ConsensusEngineT_Clique {
		if a.GetCliqueEpoch() != b.GetCliqueEpoch() {
			return fmt.Errorf("mismatch clique epochs: A: %v, B: %v", a.GetCliqueEpoch(), b.GetCliqueEpoch())
		}
		if a.GetCliquePeriod() != b.GetCliquePeriod() {
			return fmt.Errorf("mismatch clique periods: A: %v, B: %v", a.GetCliquePeriod(), b.GetCliquePeriod())
		}
	}
	return nil
}

// Transitions gets all available transition (fork) functions and their names for a ChainConfigurator.
func Transitions(conf ctypes.ChainConfigurator) (fns []func() *uint64, names []string) {
	names = []string{}
	fns = []func() *uint64{}
	k := reflect.TypeOf(conf)
	for i := 0; i < k.NumMethod(); i++ {
		method := k.Method(i)
		if !strings.HasPrefix(method.Name, "Get") || !strings.HasSuffix(method.Name, "Transition") {
			continue
		}
		m := reflect.ValueOf(conf).MethodByName(method.Name).Interface()
		fns = append(fns, m.(func() *uint64))
		names = append(names, method.Name)
	}
	return fns, names
}

// Forks returns non-nil, non <maxUin64>, unique sorted forks for a ChainConfigurator.
func Forks(conf ctypes.ChainConfigurator) []uint64 {
	var forks []uint64
	var forksM = make(map[uint64]struct{}) // Will key for uniqueness as fork numbers are appended to slice.

	transitions, _ := Transitions(conf)
	for _, tr := range transitions {
		// Extract the fork rule block number and aggregate it
		response := tr()
		if response == nil ||
			*response == math.MaxUint64 ||
			*response == 0x7fffffffffffff ||
			*response == 0x7FFFFFFFFFFFFFFF {
			continue
		}

		// Only append unique fork numbers, excluding 0 (genesis config is not considered a fork)
		if _, ok := forksM[*response]; !ok && *response != 0 {
			forks = append(forks, *response)
			forksM[*response] = struct{}{}
		}
	}
	sort.Slice(forks, func(i, j int) bool {
		return forks[i] < forks[j]
	})

	return forks
}

func isForkIncompatible(a, b, head *uint64) bool {
	return (isForked(a, head) || isForked(b, head)) && !u2Equal(a, b)
}

func isForked(x, head *uint64) bool {
	if x == nil || head == nil {
		return false
	}
	return *x <= *head
}

func u2Equal(x, y *uint64) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return *x == *y
}
