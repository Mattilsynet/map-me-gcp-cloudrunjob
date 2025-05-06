package managedenvironment

import (
	me_gcp "github.com/Mattilsynet/mapis/gen/go/managedgcpenvironment/v1"
)

func ToManagedEnvironment(meAsBytes []byte) (*me_gcp.ManagedGcpEnvironment, error) {
	me := me_gcp.ManagedGcpEnvironment{}
	err := me.UnmarshalVT(meAsBytes)
	if err != nil {
		return nil, err
	}
	return &me, nil
}

func ToBytes(me *me_gcp.ManagedGcpEnvironment) ([]byte, error) {
	return me.MarshalVT()
}
