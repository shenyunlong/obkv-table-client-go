/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2023 OceanBase
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

package main

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/client/option"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

// CREATE TABLE test(c1 bigint, c2 varchar(20), PRIMARY KEY(c1)) PARTITION BY hash(c1) partitions 2;
func main() {
	const (
		configUrl    = "xxx"
		fullUserName = "user@tenant#cluster"
		passWord     = ""
		sysUserName  = "sysUser"
		sysPassWord  = ""
		tableName    = "test"
	)
	cfg := config.NewDefaultClientConfig()
	cli, err := client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	if err != nil {
		panic(err)
	}

	// max result size for each packet from server
	batchSize := 10
	// max result size for each partition
	limit := 5
	// partition scan order
	scanOrder := table.Reverse
	// offset of partition result
	offset := 3

	// filter c2 < 100 and c2 > 50
	lt100 := filter.CompareVal(filter.LessThan, "c2", int64(100))
	gt50 := filter.CompareVal(filter.GreaterThan, "c2", int64(50))
	filterList := filter.AndList(lt100, gt50)

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", table.Min)}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", table.Max)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2"}),
		option.WithQueryBatchSize(batchSize),
		option.WithQueryLimit(limit),
		option.WithQueryScanOrder(scanOrder),
		option.WithQueryOffset(offset),
		option.WithQueryFilter(filterList),
	)
	res, err := resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		values := res.Values()
		println(values[0].(int64), values[1].(string))
		println(res.Value("c1").(int64), res.Value("c2").(string))
	}
	if err != nil {
		panic(err)
	}
}
