package objects

import "github.com/MontFerret/ferret/pkg/runtime/core"

func NewLib() map[string]core.Function {
	return map[string]core.Function{
		"HAS":             Has,
		"KEYS":            Keys,
		"KEEP_KEYS":       KeepKeys,
		"MERGE":           Merge,
		"ZIP":             Zip,
		"VALUES":          Values,
		"MERGE_RECURSIVE": MergeRecursive,
	}
}
