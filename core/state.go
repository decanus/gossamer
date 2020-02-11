// Copyright 2019 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/ChainSafe/gossamer/common"
	"github.com/ChainSafe/gossamer/core/types"
)

// BlockState interface for block state methods
type BlockState interface {
	LatestHeader() *types.Header
	AddBlock(*types.Block) error
	SetBlock(*types.Block) error
}

// NetworkState interface for network state methods
type NetworkState interface {
	GetHealth() (*common.Health, error)
	SetHealth(*common.Health) error
	GetNetworkState() (*common.NetworkState, error)
	SetNetworkState(*common.NetworkState) error
	GetPeers() (*[]common.PeerInfo, error)
	SetPeers(*[]common.PeerInfo) error
}

// StorageState interface for storage state methods
type StorageState interface {
	StorageRoot() (common.Hash, error)
}