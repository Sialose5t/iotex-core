// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package actpool

import (
	"time"
)

type ttlOption struct{ ttl time.Duration }

// WithTimeOut returns an option to overwrite time out setting.
func WithTimeOut(ttl time.Duration) interface{ ActQueueOption } {
	return &ttlOption{ttl}
}

func (o *ttlOption) SetActQueueOption(aq *actQueue) { aq.ttl = o.ttl }
