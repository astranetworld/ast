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

package conf

type DatabaseConfig struct {
	DBType     string   `json:"db_type" yaml:"db_type"`
	DBPath     string   `json:"path" yaml:"path"`
	DBName     string   `json:"name" yaml:"name"`
	SubDB      []string `json:"sub_name" yaml:"sub_name"`
	Debug      bool     `json:"debug" yaml:"debug"`
	IsMem      bool     `json:"memory" yaml:"memory"`
	MaxDB      uint64   `json:"max_db" yaml:"max_db"`
	MaxReaders uint64   `json:"max_readers" yaml:"max_readers"`
}
