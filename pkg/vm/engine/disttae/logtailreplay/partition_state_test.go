// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logtailreplay

import (
	"testing"

	"github.com/matrixorigin/matrixone/pkg/catalog"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/objectio"
	"github.com/stretchr/testify/assert"
)

func BenchmarkPartitionStateConcurrentWriteAndIter(b *testing.B) {
	partition := NewPartition()
	end := make(chan struct{})
	defer func() {
		close(end)
	}()

	// concurrent writer
	go func() {
		for {
			select {
			case <-end:
				return
			default:
			}
			state, end := partition.MutateState()
			_ = state
			end()
		}
	}()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			state := partition.state.Load()
			iter := state.NewRowsIter(types.BuildTS(0, 0), nil, false)
			iter.Close()
		}
	})

}

func TestTruncate(t *testing.T) {
	partition := NewPartitionState(true)
	addBlock(partition, types.BuildTS(1, 0), types.BuildTS(2, 0))
	addBlock(partition, types.BuildTS(1, 0), types.BuildTS(3, 0))
	addBlock(partition, types.BuildTS(1, 0), types.TS{})

	partition.truncate([2]uint64{0, 0}, types.BuildTS(1, 0))
	assert.Equal(t, 3, partition.blocks.Len())
	assert.Equal(t, 5, partition.blockIndexByTS.Len())

	partition.truncate([2]uint64{0, 0}, types.BuildTS(2, 0))
	assert.Equal(t, 2, partition.blocks.Len())
	assert.Equal(t, 3, partition.blockIndexByTS.Len())

	partition.truncate([2]uint64{0, 0}, types.BuildTS(3, 0))
	assert.Equal(t, 1, partition.blocks.Len())
	assert.Equal(t, 1, partition.blockIndexByTS.Len())

	partition.truncate([2]uint64{0, 0}, types.BuildTS(4, 0))
	assert.Equal(t, 1, partition.blocks.Len())
	assert.Equal(t, 1, partition.blockIndexByTS.Len())
}

func addBlock(p *PartitionState, create, delete types.TS) {
	blkID := objectio.NewBlockid(objectio.NewSegmentid(), 0, 0)
	blk1 := BlockEntry{
		BlockInfo: catalog.BlockInfo{
			BlockID: *blkID,
		},
		CreateTime: create,
		DeleteTime: delete,
	}
	p.blocks.Set(blk1)
	blkIndex1 := BlockIndexByTSEntry{
		Time:     create,
		BlockID:  *blkID,
		IsDelete: false,
	}
	p.blockIndexByTS.Set(blkIndex1)
	if !delete.IsEmpty() {
		blkIndex2 := BlockIndexByTSEntry{
			Time:     delete,
			BlockID:  *blkID,
			IsDelete: true,
		}
		p.blockIndexByTS.Set(blkIndex2)
	}

}
