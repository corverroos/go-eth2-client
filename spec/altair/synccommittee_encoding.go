// Code generated by fastssz. DO NOT EDIT.
// Hash: 649a003dbf727429461fbc6316a0a0cca130b264ccb563ee1f3607a4dae4fc6c
package altair

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SyncCommittee object
func (s *SyncCommittee) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SyncCommittee object to a target array
func (s *SyncCommittee) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Pubkeys'
	if len(s.Pubkeys) != 512 {
		err = ssz.ErrVectorLength
		return
	}
	for ii := 0; ii < 512; ii++ {
		dst = append(dst, s.Pubkeys[ii][:]...)
	}

	// Field (1) 'AggregatePubkey'
	dst = append(dst, s.AggregatePubkey[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SyncCommittee object
func (s *SyncCommittee) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 24624 {
		return ssz.ErrSize
	}

	// Field (0) 'Pubkeys'
	s.Pubkeys = make([]phase0.BLSPubKey, 512)
	for ii := 0; ii < 512; ii++ {
		copy(s.Pubkeys[ii][:], buf[0:24576][ii*48:(ii+1)*48])
	}

	// Field (1) 'AggregatePubkey'
	copy(s.AggregatePubkey[:], buf[24576:24624])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncCommittee object
func (s *SyncCommittee) SizeSSZ() (size int) {
	size = 24624
	return
}

// HashTreeRoot ssz hashes the SyncCommittee object
func (s *SyncCommittee) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncCommittee object with a hasher
func (s *SyncCommittee) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkeys'
	{
		if len(s.Pubkeys) != 512 {
			err = ssz.ErrVectorLength
			return
		}
		subIndx := hh.Index()
		for _, i := range s.Pubkeys {
			hh.PutBytes(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (1) 'AggregatePubkey'
	hh.PutBytes(s.AggregatePubkey[:])

	hh.Merkleize(indx)
	return
}
