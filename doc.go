// Copyright 2012 GAEGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package ds provides cached presistence for the Google App Engine datastore.

ds has the same API as the "appengine/datastore" so It will work as a drop in replacement.

E.g.

	import "github.com/gaego/ds"

	u = &User{Name: "Bob"}
	key := datastore.NewKey(c, "User", "bob", 0, nil)
	key, err := ds.Put(c, key, u)

	u = new(User)
	err = ds.Get(c, key, u)


By default it will cache all `Put`s and `Get`s to memcache, but you can modify this behavior by 
calling the `ds.Register` method:

	ds.Register("User", true, false, false)

`ds.Register` takes a string representing the `Kind` and 3 `bool` 
- `userDatastore`, `useMemcache`, `useMemory`. Passing a `true` value will cause `AEGo/ds` to 
persist the record to that store. The `Memory` store is useful for records that you do not expect 
to change, but could contain stale data if you have more then one instance running.

Supported methods are:

    Put
    PutMulti
    Get
    GetMulti
    Delete
    DeleteMulti
    AllocateIDs

Note: Currently casheing only occures with `Get`. `GetMulti` pulls from the datastore.

`AEGo/ds` is a work in progress, but the code is well tested. Any feedback would be appreciated. 

*/
package ds
