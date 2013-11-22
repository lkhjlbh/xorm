// Copyright 2013 The XORM Authors. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

/*
Package xorm is a simple and powerful ORM for Go.

First, we should new an engine for a database

	engine, err := xorm.NewEngine(driverName, dataSourceName)

Method NewEngine's parameters is the same as sql.Open. It depends
drivers' implementation. Generally, one engine is enough.

There are 7 major methods and many helpful methods to use to operate database.

1. Insert one or multipe records to database

	engine.Insert(...)

2. Query one record from database

	engine.Get(...)

3. Query multiple records from database

	engine.Find(...)

4. Query multiple records and record by record handle

	engine.Iterate(...)

5. Update one or more records

	engine.Update(...)

6. Delete one or more records

	engine.Delete(...)

7. Count records

	engine.Count(...)
*/
package xorm
