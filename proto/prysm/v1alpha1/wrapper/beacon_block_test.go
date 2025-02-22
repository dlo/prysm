package wrapper_test

import (
	"bytes"
	"testing"

	types "github.com/prysmaticlabs/eth2-types"
	"github.com/prysmaticlabs/prysm/encoding/bytesutil"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
	v1alpha1 "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
	"github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1/wrapper"
	"github.com/prysmaticlabs/prysm/runtime/version"
	"github.com/prysmaticlabs/prysm/testing/assert"
	"github.com/prysmaticlabs/prysm/testing/require"
	"github.com/prysmaticlabs/prysm/testing/util"
)

func TestAltairSignedBeaconBlock_Signature(t *testing.T) {
	sig := []byte{0x11, 0x22}
	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(&ethpb.SignedBeaconBlockAltair{Block: &ethpb.BeaconBlockAltair{}, Signature: sig})
	require.NoError(t, err)

	if !bytes.Equal(sig, wsb.Signature()) {
		t.Error("Wrong signature returned")
	}
}

func TestAltairSignedBeaconBlock_Block(t *testing.T) {
	blk := &ethpb.BeaconBlockAltair{Slot: 54}
	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(&ethpb.SignedBeaconBlockAltair{Block: blk})
	require.NoError(t, err)

	assert.DeepEqual(t, blk, wsb.Block().Proto())
}

func TestAltairSignedBeaconBlock_IsNil(t *testing.T) {
	_, err := wrapper.WrappedAltairSignedBeaconBlock(nil)
	require.Equal(t, wrapper.ErrNilObjectWrapped, err)

	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(&ethpb.SignedBeaconBlockAltair{Block: &ethpb.BeaconBlockAltair{}})
	require.NoError(t, err)

	assert.Equal(t, false, wsb.IsNil())
}

func TestAltairSignedBeaconBlock_Copy(t *testing.T) {
	t.Skip("TODO: Missing mutation evaluation helpers")
}

func TestAltairSignedBeaconBlock_Proto(t *testing.T) {
	sb := &ethpb.SignedBeaconBlockAltair{
		Block:     &ethpb.BeaconBlockAltair{Slot: 66},
		Signature: []byte{0x11, 0x22},
	}
	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(sb)
	require.NoError(t, err)

	assert.Equal(t, sb, wsb.Proto())
}

func TestAltairSignedBeaconBlock_PbPhase0Block(t *testing.T) {
	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(&ethpb.SignedBeaconBlockAltair{Block: &ethpb.BeaconBlockAltair{}})
	require.NoError(t, err)

	if _, err := wsb.PbPhase0Block(); err != wrapper.ErrUnsupportedPhase0Block {
		t.Errorf("Wrong error returned. Want %v got %v", wrapper.ErrUnsupportedPhase0Block, err)
	}
}

func TestAltairSignedBeaconBlock_PbAltairBlock(t *testing.T) {
	sb := &ethpb.SignedBeaconBlockAltair{
		Block:     &ethpb.BeaconBlockAltair{Slot: 66},
		Signature: []byte{0x11, 0x22},
	}
	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(sb)
	require.NoError(t, err)

	got, err := wsb.PbAltairBlock()
	assert.NoError(t, err)
	assert.Equal(t, sb, got)
}

func TestAltairSignedBeaconBlock_MarshalSSZTo(t *testing.T) {
	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(util.HydrateSignedBeaconBlockAltair(&ethpb.SignedBeaconBlockAltair{}))
	assert.NoError(t, err)

	var b []byte
	b, err = wsb.MarshalSSZTo(b)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(b))
}

func TestAltairSignedBeaconBlock_SSZ(t *testing.T) {
	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(util.HydrateSignedBeaconBlockAltair(&ethpb.SignedBeaconBlockAltair{}))
	assert.NoError(t, err)

	b, err := wsb.MarshalSSZ()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(b))

	assert.NotEqual(t, 0, wsb.SizeSSZ())

	assert.NoError(t, wsb.UnmarshalSSZ(b))
}

func TestAltairSignedBeaconBlock_Version(t *testing.T) {
	wsb, err := wrapper.WrappedAltairSignedBeaconBlock(&ethpb.SignedBeaconBlockAltair{Block: &ethpb.BeaconBlockAltair{}})
	require.NoError(t, err)

	assert.Equal(t, version.Altair, wsb.Version())
}

func TestAltairBeaconBlock_Slot(t *testing.T) {
	slot := types.Slot(546)
	wb, err := wrapper.WrappedAltairBeaconBlock(&ethpb.BeaconBlockAltair{Slot: slot})
	require.NoError(t, err)

	assert.Equal(t, slot, wb.Slot())
}

func TestAltairBeaconBlock_ProposerIndex(t *testing.T) {
	pi := types.ValidatorIndex(555)
	wb, err := wrapper.WrappedAltairBeaconBlock(&ethpb.BeaconBlockAltair{ProposerIndex: pi})
	require.NoError(t, err)

	assert.Equal(t, pi, wb.ProposerIndex())
}

func TestAltairBeaconBlock_ParentRoot(t *testing.T) {
	root := []byte{0xAA, 0xBF, 0x33, 0x01}
	wb, err := wrapper.WrappedAltairBeaconBlock(&ethpb.BeaconBlockAltair{ParentRoot: root})
	require.NoError(t, err)

	assert.DeepEqual(t, root, wb.ParentRoot())
}

func TestAltairBeaconBlock_StateRoot(t *testing.T) {
	root := []byte{0xAA, 0xBF, 0x33, 0x01}
	wb, err := wrapper.WrappedAltairBeaconBlock(&ethpb.BeaconBlockAltair{StateRoot: root})
	require.NoError(t, err)

	assert.DeepEqual(t, root, wb.StateRoot())
}

func TestAltairBeaconBlock_Body(t *testing.T) {
	body := &ethpb.BeaconBlockBodyAltair{Graffiti: []byte{0x44}}
	wb, err := wrapper.WrappedAltairBeaconBlock(&ethpb.BeaconBlockAltair{Body: body})
	require.NoError(t, err)

	assert.Equal(t, body, wb.Body().Proto())
}

func TestAltairBeaconBlock_IsNil(t *testing.T) {
	_, err := wrapper.WrappedAltairBeaconBlock(nil)
	require.Equal(t, wrapper.ErrNilObjectWrapped, err)

	wb, err := wrapper.WrappedAltairBeaconBlock(&ethpb.BeaconBlockAltair{})
	require.NoError(t, err)

	assert.Equal(t, false, wb.IsNil())
}

func TestAltairBeaconBlock_HashTreeRoot(t *testing.T) {
	wb, err := wrapper.WrappedAltairBeaconBlock(util.HydrateBeaconBlockAltair(&ethpb.BeaconBlockAltair{}))
	require.NoError(t, err)

	rt, err := wb.HashTreeRoot()
	assert.NoError(t, err)
	assert.NotEmpty(t, rt)
}

func TestAltairBeaconBlock_Proto(t *testing.T) {
	blk := &ethpb.BeaconBlockAltair{ProposerIndex: 234}
	wb, err := wrapper.WrappedAltairBeaconBlock(blk)
	require.NoError(t, err)

	assert.Equal(t, blk, wb.Proto())
}

func TestAltairBeaconBlock_SSZ(t *testing.T) {
	wb, err := wrapper.WrappedAltairBeaconBlock(util.HydrateBeaconBlockAltair(&ethpb.BeaconBlockAltair{}))
	assert.NoError(t, err)

	b, err := wb.MarshalSSZ()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(b))

	assert.NotEqual(t, 0, wb.SizeSSZ())

	assert.NoError(t, wb.UnmarshalSSZ(b))
}

func TestAltairBeaconBlock_Version(t *testing.T) {
	wb, err := wrapper.WrappedAltairBeaconBlock(&ethpb.BeaconBlockAltair{})
	require.NoError(t, err)

	assert.Equal(t, version.Altair, wb.Version())
}

func TestAltairBeaconBlockBody_RandaoReveal(t *testing.T) {
	root := []byte{0xAA, 0xBF, 0x33, 0x01}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(&ethpb.BeaconBlockBodyAltair{RandaoReveal: root})
	require.NoError(t, err)

	assert.DeepEqual(t, root, wbb.RandaoReveal())
}

func TestAltairBeaconBlockBody_Eth1Data(t *testing.T) {
	data := &v1alpha1.Eth1Data{}
	body := &ethpb.BeaconBlockBodyAltair{
		Eth1Data: data,
	}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(body)
	require.NoError(t, err)
	assert.Equal(t, data, wbb.Eth1Data())
}

func TestAltairBeaconBlockBody_Graffiti(t *testing.T) {
	graffiti := []byte{0x66, 0xAA}
	body := &ethpb.BeaconBlockBodyAltair{Graffiti: graffiti}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(body)
	require.NoError(t, err)

	assert.DeepEqual(t, graffiti, wbb.Graffiti())
}

func TestAltairBeaconBlockBody_ProposerSlashings(t *testing.T) {
	ps := []*v1alpha1.ProposerSlashing{
		{Header_1: &v1alpha1.SignedBeaconBlockHeader{
			Signature: []byte{0x11, 0x20},
		}},
	}
	body := &ethpb.BeaconBlockBodyAltair{ProposerSlashings: ps}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(body)
	require.NoError(t, err)

	assert.DeepEqual(t, ps, wbb.ProposerSlashings())
}

func TestAltairBeaconBlockBody_AttesterSlashings(t *testing.T) {
	as := []*v1alpha1.AttesterSlashing{
		{Attestation_1: &v1alpha1.IndexedAttestation{Signature: []byte{0x11}}},
	}
	body := &ethpb.BeaconBlockBodyAltair{AttesterSlashings: as}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(body)
	require.NoError(t, err)

	assert.DeepEqual(t, as, wbb.AttesterSlashings())
}

func TestAltairBeaconBlockBody_Attestations(t *testing.T) {
	atts := []*v1alpha1.Attestation{{Signature: []byte{0x88}}}

	body := &ethpb.BeaconBlockBodyAltair{Attestations: atts}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(body)
	require.NoError(t, err)

	assert.DeepEqual(t, atts, wbb.Attestations())
}

func TestAltairBeaconBlockBody_Deposits(t *testing.T) {
	deposits := []*v1alpha1.Deposit{
		{Proof: [][]byte{{0x54, 0x10}}},
	}
	body := &ethpb.BeaconBlockBodyAltair{Deposits: deposits}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(body)
	require.NoError(t, err)

	assert.DeepEqual(t, deposits, wbb.Deposits())
}

func TestAltairBeaconBlockBody_VoluntaryExits(t *testing.T) {
	exits := []*v1alpha1.SignedVoluntaryExit{
		{Exit: &v1alpha1.VoluntaryExit{Epoch: 54}},
	}
	body := &ethpb.BeaconBlockBodyAltair{VoluntaryExits: exits}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(body)
	require.NoError(t, err)

	assert.DeepEqual(t, exits, wbb.VoluntaryExits())
}

func TestAltairBeaconBlockBody_IsNil(t *testing.T) {
	_, err := wrapper.WrappedAltairBeaconBlockBody(nil)
	require.Equal(t, wrapper.ErrNilObjectWrapped, err)

	wbb, err := wrapper.WrappedAltairBeaconBlockBody(&ethpb.BeaconBlockBodyAltair{})
	require.NoError(t, err)
	assert.Equal(t, false, wbb.IsNil())

}

func TestAltairBeaconBlockBody_HashTreeRoot(t *testing.T) {
	wb, err := wrapper.WrappedAltairBeaconBlockBody(util.HydrateBeaconBlockBodyAltair(&ethpb.BeaconBlockBodyAltair{}))
	assert.NoError(t, err)

	rt, err := wb.HashTreeRoot()
	assert.NoError(t, err)
	assert.NotEmpty(t, rt)
}

func TestAltairBeaconBlockBody_Proto(t *testing.T) {
	body := &ethpb.BeaconBlockBodyAltair{Graffiti: []byte{0x66, 0xAA}}
	wbb, err := wrapper.WrappedAltairBeaconBlockBody(body)
	require.NoError(t, err)

	assert.Equal(t, body, wbb.Proto())
}

func TestPhase0SignedBeaconBlock_Header(t *testing.T) {
	root := bytesutil.PadTo([]byte("root"), 32)
	signature := bytesutil.PadTo([]byte("sig"), 96)
	body := &ethpb.BeaconBlockBody{}
	body = util.HydrateBeaconBlockBody(body)
	bodyRoot, err := body.HashTreeRoot()
	require.NoError(t, err)
	block := &ethpb.SignedBeaconBlock{
		Block: &ethpb.BeaconBlock{
			Slot:          1,
			ProposerIndex: 1,
			ParentRoot:    root,
			StateRoot:     root,
			Body:          body,
		},
		Signature: signature,
	}
	wrapped := wrapper.WrappedPhase0SignedBeaconBlock(block)

	header, err := wrapped.Header()
	require.NoError(t, err)
	assert.Equal(t, types.ValidatorIndex(1), header.Header.ProposerIndex)
	assert.Equal(t, types.Slot(1), header.Header.Slot)
	assert.DeepEqual(t, bodyRoot[:], header.Header.BodyRoot)
	assert.DeepEqual(t, root, header.Header.StateRoot)
	assert.DeepEqual(t, root, header.Header.ParentRoot)
	assert.DeepEqual(t, signature, header.Signature)
}

func TestAltairSignedBeaconBlock_Header(t *testing.T) {
	root := bytesutil.PadTo([]byte("root"), 32)
	signature := bytesutil.PadTo([]byte("sig"), 96)
	body := &ethpb.BeaconBlockBodyAltair{}
	body = util.HydrateBeaconBlockBodyAltair(body)
	bodyRoot, err := body.HashTreeRoot()
	require.NoError(t, err)
	block := &ethpb.SignedBeaconBlockAltair{
		Block: &ethpb.BeaconBlockAltair{
			Slot:          1,
			ProposerIndex: 1,
			ParentRoot:    root,
			StateRoot:     root,
			Body:          body,
		},
		Signature: signature,
	}
	wrapped, err := wrapper.WrappedAltairSignedBeaconBlock(block)
	require.NoError(t, err)

	header, err := wrapped.Header()
	require.NoError(t, err)
	assert.Equal(t, types.ValidatorIndex(1), header.Header.ProposerIndex)
	assert.Equal(t, types.Slot(1), header.Header.Slot)
	assert.DeepEqual(t, bodyRoot[:], header.Header.BodyRoot)
	assert.DeepEqual(t, root, header.Header.StateRoot)
	assert.DeepEqual(t, root, header.Header.ParentRoot)
	assert.DeepEqual(t, signature, header.Signature)
}
