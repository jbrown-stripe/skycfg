package skycfg

import (
	"fmt"

	"github.com/google/skylark"
)

// NewProtoPackage creates a Skylark value representing a named Protobuf package.
//
// Protobuf packagess are conceptually similar to a C++ namespace or Ruby
// module, in that they're aggregated from multiple .proto source files.
func newProtoPackage(name string) skylark.Value {
	return &skyProtoPackage{
		name: name,
	}
}

type skyProtoPackage struct {
	name string
}

func (pkg *skyProtoPackage) String() string      { return fmt.Sprintf("<proto.Package %q>", pkg.name) }
func (pkg *skyProtoPackage) Type() string        { return "proto.Package" }
func (pkg *skyProtoPackage) Freeze()             {}
func (pkg *skyProtoPackage) Truth() skylark.Bool { return skylark.True }
func (pkg *skyProtoPackage) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", pkg.Type())
}

func (pkg *skyProtoPackage) AttrNames() []string {
	// TODO: Implement when go-protobuf gains support for listing the
	// registered message types in a Protobuf package.
	//
	// https://github.com/golang/protobuf/issues/623
	return nil
}

func (pkg *skyProtoPackage) Attr(attrName string) (skylark.Value, error) {
	return newMessageType(fmt.Sprintf("%s.%s", pkg.name, attrName))
}