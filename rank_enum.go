/*
 * Copyright (c) 2021 OmegaRogue.
 *  This file is part of eliteJournal.
 *
 *     eliteJournal is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or any later version.
 *
 *     eliteJournal is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with eliteJournal.  If not, see <https://www.gnu.org/licenses/>.
 */

// Code generated by go-enum
// DO NOT EDIT!

package elite

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/rs/zerolog"
)

const (
	// None is a Rank of type None.
	None Rank = iota + -1
	// Rookie is a Rank of type Rookie.
	Rookie
	// Agent is a Rank of type Agent.
	Agent
	// Officer is a Rank of type Officer.
	Officer
	// SeniorOfficer is a Rank of type SeniorOfficer.
	SeniorOfficer
	// Leader is a Rank of type Leader.
	Leader
)

const _RankName = "NoneRookieAgentOfficerSeniorOfficerLeader"

var _RankMap = map[Rank]string{
	-1: _RankName[0:4],
	0:  _RankName[4:10],
	1:  _RankName[10:15],
	2:  _RankName[15:22],
	3:  _RankName[22:35],
	4:  _RankName[35:41],
}

// String implements the Stringer interface.
func (x Rank) String() string {
	if str, ok := _RankMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Rank(%d)", x)
}

var _RankValue = map[string]Rank{
	_RankName[0:4]:   -1,
	_RankName[4:10]:  0,
	_RankName[10:15]: 1,
	_RankName[15:22]: 2,
	_RankName[22:35]: 3,
	_RankName[35:41]: 4,
}

// ParseRank attempts to convert a string to a Rank
func ParseRank(name string) (Rank, error) {
	if x, ok := _RankValue[name]; ok {
		return x, nil
	}
	return Rank(0), fmt.Errorf("%s is not a valid Rank", name)
}

// MarshalText implements the text marshaller method
func (x Rank) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *Rank) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseRank(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var _RankErrNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *Rank) Scan(value interface{}) (err error) {
	if value == nil {
		*x = Rank(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = Rank(v)
	case string:
		*x, err = ParseRank(v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(v); verr == nil {
				*x, err = Rank(val), nil
			}
		}
	case []byte:
		*x, err = ParseRank(string(v))
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(string(v)); verr == nil {
				*x, err = Rank(val), nil
			}
		}
	case Rank:
		*x = v
	case int:
		*x = Rank(v)
	case *Rank:
		if v == nil {
			return _RankErrNilPtr
		}
		*x = *v
	case uint:
		*x = Rank(v)
	case uint64:
		*x = Rank(v)
	case *int:
		if v == nil {
			return _RankErrNilPtr
		}
		*x = Rank(*v)
	case *int64:
		if v == nil {
			return _RankErrNilPtr
		}
		*x = Rank(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = Rank(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return _RankErrNilPtr
		}
		*x = Rank(*v)
	case *uint:
		if v == nil {
			return _RankErrNilPtr
		}
		*x = Rank(*v)
	case *uint64:
		if v == nil {
			return _RankErrNilPtr
		}
		*x = Rank(*v)
	case *string:
		if v == nil {
			return _RankErrNilPtr
		}
		*x, err = ParseRank(*v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(*v); verr == nil {
				*x, err = Rank(val), nil
			}
		}
	}

	return
}

// Value implements the driver Valuer interface.
func (x Rank) Value() (driver.Value, error) {
	return int64(x), nil
}

type NullRank struct {
	Rank  Rank
	Valid bool
	Set   bool
}

func NewNullRank(val interface{}) (x NullRank) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Scan implements the Scanner interface.
func (x *NullRank) Scan(value interface{}) (err error) {
	x.Set = true
	if value == nil {
		x.Rank, x.Valid = Rank(0), false
		return
	}

	err = x.Rank.Scan(value)
	x.Valid = err == nil
	return
}

// Value implements the driver Valuer interface.
func (x NullRank) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return int64(x.Rank), nil
}

// MarshalJSON correctly serializes a NullRank to JSON.
func (n NullRank) MarshalJSON() ([]byte, error) {
	const nullStr = "null"
	if n.Valid {
		return json.Marshal(n.Rank)
	}
	return []byte(nullStr), nil
}

// UnmarshalJSON correctly deserializes a NullRank from JSON.
func (n *NullRank) UnmarshalJSON(b []byte) error {
	n.Set = true
	var x interface{}
	err := json.Unmarshal(b, &x)
	if err != nil {
		return err
	}
	err = n.Scan(x)
	return err
}

func (x Rank) MarshalZerologObject(e *zerolog.Event) {
	e.Str("Rank", fmt.Sprint(x))
}
