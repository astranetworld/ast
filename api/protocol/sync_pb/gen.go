// Copyright 2022 The N42 Authors
// This file is part of the N42 library.
//
// The N42 library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The N42 library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the N42 library. If not, see <http://www.gnu.org/licenses/>.

package sync_pb

////go:generate protoc --plugin=/Users/mac/go/bin/protoc-gen-go-cast -I=../ -I=. -I=../include --go-cast_out=plugins=protoc-gen-go-cast,paths=source_relative:. types.proto
//go:generate protoc  -I=../ -I=. -I=../include --go-cast_out=paths=source_relative:. sync_pb.proto
//go:generate sszgen -path=. -objs=BodiesByRangeRequest,HeadersByRangeRequest,Ping,ForkData,Status --include=../types_pb -output=generated.ssz.go
