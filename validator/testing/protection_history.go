package testing

import (
	"fmt"

	types "github.com/prysmaticlabs/eth2-types"
	"github.com/prysmaticlabs/prysm/config/params"
	"github.com/prysmaticlabs/prysm/crypto/bls"
	"github.com/prysmaticlabs/prysm/crypto/rand"
	"github.com/prysmaticlabs/prysm/encoding/bytesutil"
	"github.com/prysmaticlabs/prysm/validator/db/kv"
	"github.com/prysmaticlabs/prysm/validator/slashing-protection/local/standard-protection-format/format"
)

// MockSlashingProtectionJSON creates a mock, full slashing protection JSON struct
// using attesting and proposing histories provided.
func MockSlashingProtectionJSON(
	publicKeys [][48]byte,
	attestingHistories [][]*kv.AttestationRecord,
	proposalHistories []kv.ProposalHistoryForPubkey,
) (*format.EIPSlashingProtectionFormat, error) {
	standardProtectionFormat := &format.EIPSlashingProtectionFormat{}
	standardProtectionFormat.Metadata.GenesisValidatorsRoot = fmt.Sprintf("%#x", bytesutil.PadTo([]byte{32}, 32))
	standardProtectionFormat.Metadata.InterchangeFormatVersion = format.InterchangeFormatVersion
	for i := 0; i < len(publicKeys); i++ {
		data := &format.ProtectionData{
			Pubkey: fmt.Sprintf("%#x", publicKeys[i]),
		}
		if len(attestingHistories) > 0 {
			for _, att := range attestingHistories[i] {
				data.SignedAttestations = append(data.SignedAttestations, &format.SignedAttestation{
					TargetEpoch: fmt.Sprintf("%d", att.Target),
					SourceEpoch: fmt.Sprintf("%d", att.Source),
					SigningRoot: fmt.Sprintf("%#x", att.SigningRoot),
				})
			}
		}
		if len(proposalHistories) > 0 {
			for _, proposal := range proposalHistories[i].Proposals {
				block := &format.SignedBlock{
					Slot:        fmt.Sprintf("%d", proposal.Slot),
					SigningRoot: fmt.Sprintf("%#x", proposal.SigningRoot),
				}
				data.SignedBlocks = append(data.SignedBlocks, block)
			}
		}
		standardProtectionFormat.Data = append(standardProtectionFormat.Data, data)
	}
	return standardProtectionFormat, nil
}

// MockAttestingAndProposalHistories given a number of validators, creates mock attesting
// and proposing histories within WEAK_SUBJECTIVITY_PERIOD bounds.
func MockAttestingAndProposalHistories(pubkeys [][48]byte) ([][]*kv.AttestationRecord, []kv.ProposalHistoryForPubkey) {
	// deduplicate and transform them into our internal format.
	numValidators := len(pubkeys)
	attData := make([][]*kv.AttestationRecord, numValidators)
	proposalData := make([]kv.ProposalHistoryForPubkey, numValidators)
	gen := rand.NewGenerator()
	for v := 0; v < numValidators; v++ {
		latestTarget := types.Epoch(gen.Intn(int(params.BeaconConfig().WeakSubjectivityPeriod) / 1000))
		// If 0, we change the value to 1 as the we compute source by doing (target-1)
		// to prevent any underflows in this setup helper.
		if latestTarget == 0 {
			latestTarget = 1
		}
		historicalAtts := make([]*kv.AttestationRecord, 0)
		proposals := make([]kv.Proposal, 0)
		for i := types.Epoch(1); i < latestTarget; i++ {
			signingRoot := [32]byte{}
			signingRootStr := fmt.Sprintf("%d", i)
			copy(signingRoot[:], signingRootStr)
			historicalAtts = append(historicalAtts, &kv.AttestationRecord{
				Source:      i - 1,
				Target:      i,
				SigningRoot: signingRoot,
				PubKey:      pubkeys[v],
			})
		}
		for i := types.Epoch(1); i <= latestTarget; i++ {
			signingRoot := [32]byte{}
			signingRootStr := fmt.Sprintf("%d", i)
			copy(signingRoot[:], signingRootStr)
			proposals = append(proposals, kv.Proposal{
				Slot:        types.Slot(i),
				SigningRoot: signingRoot[:],
			})
		}
		proposalData[v] = kv.ProposalHistoryForPubkey{Proposals: proposals}
		attData[v] = historicalAtts
	}
	return attData, proposalData
}

// CreateRandomPubKeys --
func CreateRandomPubKeys(numValidators int) ([][48]byte, error) {
	pubKeys := make([][48]byte, numValidators)
	for i := 0; i < numValidators; i++ {
		randKey, err := bls.RandKey()
		if err != nil {
			return nil, err
		}
		copy(pubKeys[i][:], randKey.PublicKey().Marshal())
	}
	return pubKeys, nil
}

// CreateMockRoots --
func CreateMockRoots(numRoots int) [][32]byte {
	roots := make([][32]byte, numRoots)
	for i := 0; i < numRoots; i++ {
		var rt [32]byte
		copy(rt[:], fmt.Sprintf("%d", i))
	}
	return roots
}
