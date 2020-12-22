package models

import (
	"github.com/zaccari/gosmi/types"
)

type Node struct {
	Access      types.Access
	Decl        types.Decl
	Description string
	Kind        types.NodeKind
	Name        string
	Oid         types.Oid
	OidLen      int
	Status      types.Status
	Type        *Type
}

// https://sourceforge.net/p/net-snmp/mailman/message/16580196/
//
// An Object type (an object class) that is a columnar or leaf object is
// identified by an OID value (called it's name). The last sub-identifier cannot
// be zero. (See RFC 2578, last paragraph of section 7.10).
// ...
// The result is that the variable (an OID value) that identifies any instance
// of a scalar always ends with a sub-identifier with a value of zero.
func (n Node) IsScalar() bool {
	if len(n.Oid) == 0 {
		return false
	}

	lastSubId := n.Oid[len(n.Oid)-1]
	return lastSubId == 0
}
