// Copyright 2017 AMIS Technologies
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package extra

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func Decode(extraData string) ([]byte, *IstanbulExtra, error) {
	extra, err := hexutil.Decode(extraData)
	if err != nil {
		return nil, nil, err
	}

	istanbulExtra, err := ExtractIstanbulExtra(&types.Header{Extra: extra})
	if err != nil {
		return nil, nil, err
	}
	return extra[:types.IstanbulExtraVanity], istanbulExtra, nil
}

// ExtractIstanbulExtra extracts all values of the IstanbulExtra from the header. It returns an
// error if the length of the given extra-data is less than 32 bytes or the extra-data can not
// be decoded.
func ExtractIstanbulExtra(h *types.Header) (*IstanbulExtra, error) {
	if len(h.Extra) < types.IstanbulExtraVanity {
		return nil, types.ErrInvalidIstanbulHeaderExtra
	}

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(h.Extra[types.IstanbulExtraVanity:], &istanbulExtra)
	if err != nil {
		return nil, err
	}
	return istanbulExtra, nil
}
