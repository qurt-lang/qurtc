package machine

import (
	"maps"

	"github.com/nurtai325/qurtc/internal/ast"
	"github.com/nurtai325/qurtc/internal/types"
)

type scope struct {
	vars       map[string]types.Type
	isLoop     bool
	isContinue bool
	isBreak    bool
}

func newFuncScope(funcDecl *ast.FuncDecl, args []types.Type) (*scope, error) {
	if len(funcDecl.Args) != len(args) {
		return nil, ErrFuncArgMismatch
	}
	newScope := scope{
		vars: make(map[string]types.Type, len(funcDecl.Args)),
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
	_, ok := s.vars[name]
	if ok {
		return false
	}
	s.vars[name] = value
	return true
}

func (s *scope) get(name string) types.Type {
	return s.vars[name]
}

func (s *scope) set(name string, value types.Type) bool {
	_, ok := s.vars[name]
	if !ok {
		return false
	}
	s.vars[name] = value
	return true
}

func (s *scope) clone() *scope {
	return &scope{
		vars:   maps.Clone(s.vars),
		isLoop: s.isLoop,
	}
}
