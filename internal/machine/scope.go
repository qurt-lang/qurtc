package machine

import (
	"maps"
	"fmt"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/types"
)

type scope struct {
	parentVars map[string]types.Type
	childVars  map[string]types.Type
	isLoop     bool
	isContinue bool
	isBreak    bool
}

func newFuncScope(funcDecl *ast.FuncDecl, args []types.Type) (*scope, error) {
	if len(funcDecl.Args) != len(args) {
		return nil, ErrFuncArgMismatch
	}
	newScope := scope{
		parentVars: make(map[string]types.Type),
		childVars:  make(map[string]types.Type, len(funcDecl.Args)),
	}
	for i, arg := range funcDecl.Args {
		if !types.IsOfType(args[i], arg.Type) {
			return nil, ErrFuncArgMismatch
		}
		newScope.add(arg.Name, args[i])
	}
	return &newScope, nil
}

func (s *scope) add(name string, value types.Type) bool {
	_, ok := s.childVars[name]
	if ok {
		fmt.Println("hello", name, value, s.childVars)
		return false
	}
	s.childVars[name] = value
	return true
}

func (s *scope) get(name string) types.Type {
	val, ok := s.childVars[name]
	if !ok {
		return s.parentVars[name]
	}
	return val
}

func (s *scope) set(name string, value types.Type) bool {
	_, ok := s.childVars[name]
	if !ok {
		_, ok := s.parentVars[name]
		if !ok {
			return false
		}
		s.parentVars[name] = value
		return true
	}
	s.childVars[name] = value
	return true
}

func (s *scope) newBlockScope() *scope {
	return &scope{
		parentVars: mergeScopes(s.parentVars, s.childVars),
		childVars:  make(map[string]types.Type),
		isLoop:     s.isLoop,
	}
}

func mergeScopes(childScope, parentScope map[string]types.Type) map[string]types.Type {
	merged := make(map[string]types.Type, len(childScope)+len(parentScope))
	maps.Copy(merged, parentScope)
	maps.Copy(merged, childScope)
	return merged
}
