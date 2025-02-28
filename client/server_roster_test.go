/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package client

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/route"
)

func TestObServerRoster(t *testing.T) {
	r := &obServerRoster{}
	assert.Equal(t, "obServerRoster{maxPriority:0, roster:[]}", r.String())
	s1 := route.NewObServerAddr("127.0.0.1", 4001, 4000)
	s2 := route.NewObServerAddr("127.0.0.2", 4001, 4000)
	s3 := route.NewObServerAddr("127.0.0.3", 4001, 4000)

	r = &obServerRoster{atomic.Int32{}, nil}
	r.Reset([]*route.ObServerAddr{s1, s2, s3})
	r.maxPriority.Store(1)
	assert.EqualValues(t, 1, r.MaxPriority())
	assert.EqualValues(t, 3, r.Size())
	assert.NotEqual(t, nil, r.GetServer())
	assert.Equal(t, "obServerRoster{maxPriority:1, roster:[ObServerAddr{ip:127.0.0.1, sqlPort:4001, svrPort:4000}, ObServerAddr{ip:127.0.0.2, sqlPort:4001, svrPort:4000}, ObServerAddr{ip:127.0.0.3, sqlPort:4001, svrPort:4000}]}", r.String())
}
