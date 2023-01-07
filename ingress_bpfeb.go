// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64
// +build arm64be armbe mips mips64 mips64p32 ppc64 s390 s390x sparc sparc64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type ingressSeg6localKey struct {
	Base struct {
		Prefixlen uint32
		Data      [0]uint8
	}
	Prefix struct{ In6U struct{ U6Addr8 [16]uint8 } }
	Pad    [12]uint8
}

type ingressSeg6localVal struct {
	Action uint32
	Attr   struct{ Vrftable uint32 }
}

// loadIngress returns the embedded CollectionSpec for ingress.
func loadIngress() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_IngressBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load ingress: %w", err)
	}

	return spec, err
}

// loadIngressObjects loads ingress and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*ingressObjects
//	*ingressPrograms
//	*ingressMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadIngressObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadIngress()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// ingressSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ingressSpecs struct {
	ingressProgramSpecs
	ingressMapSpecs
}

// ingressSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ingressProgramSpecs struct {
	IngressMain *ebpf.ProgramSpec `ebpf:"ingress_main"`
}

// ingressMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ingressMapSpecs struct {
	Vrftable2ifindex  *ebpf.MapSpec `ebpf:"vrftable2ifindex"`
	ZebraSeg6localMap *ebpf.MapSpec `ebpf:"zebra_seg6local_map"`
}

// ingressObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadIngressObjects or ebpf.CollectionSpec.LoadAndAssign.
type ingressObjects struct {
	ingressPrograms
	ingressMaps
}

func (o *ingressObjects) Close() error {
	return _IngressClose(
		&o.ingressPrograms,
		&o.ingressMaps,
	)
}

// ingressMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadIngressObjects or ebpf.CollectionSpec.LoadAndAssign.
type ingressMaps struct {
	Vrftable2ifindex  *ebpf.Map `ebpf:"vrftable2ifindex"`
	ZebraSeg6localMap *ebpf.Map `ebpf:"zebra_seg6local_map"`
}

func (m *ingressMaps) Close() error {
	return _IngressClose(
		m.Vrftable2ifindex,
		m.ZebraSeg6localMap,
	)
}

// ingressPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadIngressObjects or ebpf.CollectionSpec.LoadAndAssign.
type ingressPrograms struct {
	IngressMain *ebpf.Program `ebpf:"ingress_main"`
}

func (p *ingressPrograms) Close() error {
	return _IngressClose(
		p.IngressMain,
	)
}

func _IngressClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed ingress_bpfeb.o
var _IngressBytes []byte
