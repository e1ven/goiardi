/*
 * Copyright (c) 2013-2014, Jeremy Bingham (<jbingham@gmail.com>)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package shovey

/* MySQL funcs for shovey */

import (
	"github.com/ctdk/goiardi/datastore"
	"github.com/go-sql-driver/mysql"
	"github.com/ctdk/goiardi/util"
	"net/http"
	"time"
)

func (s *Shovey) fillShoveyFromMySQL(row datastore.ResRow) error {
	var ca, ua mysql.NullTime
	var nn util.StringSlice
	var tm int64
	err := row.Scan(&s.RunID, &nn, &s.Command, &ca, &ua, &s.Status, &tm, &s.Quorum)
	if err != nil {
		return err
	}
	if ca.Valid {
		s.CreatedAt = ca.Time
	}
	if ua.Valid {
		s.UpdatedAt = ua.Time
	}
	s.Timeout = time.Duration(tm)

	s.NodeNames = nn

	return nil
}

func (s *ShoveyRun) fillShoveyRunFromMySQL(row datastore.ResRow) error {
	var at, et mysql.NullTime
	err := row.Scan(&s.ID, &s.ShoveyUUID, &s.NodeName, &s.Status, &at, &et, &s.Output, &s.Error, &s.Stderr, &s.ExitStatus)
	if err != nil {
		return err
	}
	if at.Valid {
		s.AckTime = at.Time
	}
	if et.Valid {
		s.EndTime = et.Time
	}
	return nil
}

func (s *ShoveyRunStream) fillShoveyRunStreamFromMySQL(row datastore.ResRow) error {
	var ca mysql.NullTime
	err := row.Scan(&s.ShoveyUUID, &s.NodeName, &s.Seq, &s.OutputType, &s.Output, &s.IsLast, &ca)
	if err != nil {
		return err
	}
	if ca.Valid {
		s.CreatedAt = ca.Time
	}
	return nil
}

func (s *Shovey) saveMySQL() util.Gerror {
	tx, err := datastore.Dbh.Begin()
	if err != nil {
		gerr := util.CastErr(err)
		gerr.SetStatus(http.StatusInternalServerError)
		return gerr
	}
}

func (sr *ShoveyRun) saveMySQL() util.Gerror {
	tx, err := datastore.Dbh.Begin()
	if err != nil {
		gerr := util.CastErr(err)
		gerr.SetStatus(http.StatusInternalServerError)
		return gerr
	}
}
