/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package routetable

import (
	"fmt"

	"hcm/pkg/api/core"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/dal/dao/audit"
	idgenerator "hcm/pkg/dal/dao/id-generator"
	"hcm/pkg/dal/dao/orm"
	"hcm/pkg/dal/dao/tools"
	"hcm/pkg/dal/dao/types"
	"hcm/pkg/dal/table"
	tableaudit "hcm/pkg/dal/table/audit"
	routetable "hcm/pkg/dal/table/cloud/route-table"
	"hcm/pkg/kit"
	"hcm/pkg/logs"
	"hcm/pkg/runtime/filter"
	"hcm/pkg/tools/converter"

	"github.com/jmoiron/sqlx"
)

// GcpRoute defines gcp route dao operations.
type GcpRoute interface {
	BatchCreateWithTx(kt *kit.Kit, tx *sqlx.Tx, models []routetable.GcpRouteTable) ([]string, error)
	List(kt *kit.Kit, opt *types.ListOption, whereOpts ...*filter.SQLWhereOption) (*types.GcpRouteListResult, error)
	BatchDeleteWithTx(kt *kit.Kit, tx *sqlx.Tx, expr *filter.Expression) error
}

var _ GcpRoute = new(gcpRouteDao)

// gcpRouteDao gcp route dao.
type gcpRouteDao struct {
	orm   orm.Interface
	idGen idgenerator.IDGenInterface
	audit audit.Interface
}

// NewGcpRouteDao create a gcp route dao.
func NewGcpRouteDao(orm orm.Interface, idGen idgenerator.IDGenInterface, audit audit.Interface) GcpRoute {
	return &gcpRouteDao{
		orm:   orm,
		idGen: idGen,
		audit: audit,
	}
}

// BatchCreateWithTx create gcp route with transaction.
func (r *gcpRouteDao) BatchCreateWithTx(kt *kit.Kit, tx *sqlx.Tx, models []routetable.GcpRouteTable) (
	[]string, error) {

	if len(models) == 0 {
		return nil, errf.New(errf.InvalidParameter, "models to create cannot be empty")
	}

	for _, model := range models {
		if err := model.InsertValidate(); err != nil {
			return nil, err
		}
	}

	// generate gcp route id
	ids, err := r.idGen.Batch(kt, table.GcpRouteTable, len(models))
	if err != nil {
		return nil, err
	}

	for idx := range models {
		models[idx].ID = ids[idx]
	}

	sql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, models[0].TableName(),
		routetable.GcpRouteColumns.ColumnExpr(), routetable.GcpRouteColumns.ColonNameExpr())

	err = r.orm.Txn(tx).BulkInsert(kt.Ctx, sql, models)
	if err != nil {
		return nil, fmt.Errorf("insert %s failed, err: %v", models[0].TableName(), err)
	}

	if err = r.batchCreateAudit(kt, tx, models); err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *gcpRouteDao) batchCreateAudit(kt *kit.Kit, tx *sqlx.Tx, routes []routetable.GcpRouteTable) error {
	rtIDMap := make(map[string]struct{}, 0)
	for _, route := range routes {
		rtIDMap[route.RouteTableID] = struct{}{}
	}

	rtIDs := make([]string, 0, len(rtIDMap))
	for id := range rtIDMap {
		rtIDs = append(rtIDs, id)
	}

	idRtMap, err := listRouteTable(kt, r.orm, tx, rtIDs)
	if err != nil {
		return err
	}

	audits := make([]*tableaudit.AuditTable, 0, len(routes))
	for _, route := range routes {
		rt, exist := idRtMap[route.RouteTableID]
		if !exist {
			return errf.Newf(errf.RecordNotFound, "route table: %s not found", route.RouteTableID)
		}

		audits = append(audits, &tableaudit.AuditTable{
			ResID:      rt.ID,
			CloudResID: rt.CloudID,
			ResName:    converter.PtrToVal(rt.Name),
			ResType:    enumor.RouteTableAuditResType,
			Action:     enumor.Update,
			BkBizID:    rt.BkBizID,
			Vendor:     rt.Vendor,
			AccountID:  rt.AccountID,
			Operator:   kt.User,
			Source:     kt.GetRequestSource(),
			Rid:        kt.Rid,
			AppCode:    kt.AppCode,
			Detail: &tableaudit.BasicDetail{
				Data: &tableaudit.ChildResAuditData{
					ChildResType: enumor.RouteAuditResType,
					Action:       enumor.Create,
					ChildRes:     route,
				},
			},
		})
	}

	if err = r.audit.BatchCreateWithTx(kt, tx, audits); err != nil {
		logs.Errorf("batch create audit failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	return nil
}

// List gcp routes.
func (r *gcpRouteDao) List(kt *kit.Kit, opt *types.ListOption, whereOpts ...*filter.SQLWhereOption) (
	*types.GcpRouteListResult, error) {

	if opt == nil {
		return nil, errf.New(errf.InvalidParameter, "list gcp route options is nil")
	}

	if err := opt.Validate(filter.NewExprOption(filter.RuleFields(routetable.GcpRouteColumns.ColumnTypes())),
		core.DefaultPageOption); err != nil {
		return nil, err
	}

	whereOpt := tools.DefaultSqlWhereOption
	if len(whereOpts) != 0 && whereOpts[0] != nil {
		err := whereOpts[0].Validate()
		if err != nil {
			return nil, err
		}
		whereOpt = whereOpts[0]
	}
	whereExpr, whereValue, err := opt.Filter.SQLWhereExpr(whereOpt)
	if err != nil {
		return nil, err
	}

	if opt.Page.Count {
		// this is a count request, do count operation only.
		sql := fmt.Sprintf(`SELECT COUNT(*) FROM %s %s`, table.GcpRouteTable, whereExpr)

		count, err := r.orm.Do().Count(kt.Ctx, sql, whereValue)
		if err != nil {
			logs.ErrorJson("count gcp routes failed, err: %v, filter: %s, rid: %s", err, opt.Filter, kt.Rid)
			return nil, err
		}

		return &types.GcpRouteListResult{Count: count}, nil
	}

	pageExpr, err := types.PageSQLExpr(opt.Page, types.DefaultPageSQLOption)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(`SELECT %s FROM %s %s %s`, routetable.GcpRouteColumns.FieldsNamedExpr(opt.Fields),
		table.GcpRouteTable, whereExpr, pageExpr)

	details := make([]routetable.GcpRouteTable, 0)
	if err = r.orm.Do().Select(kt.Ctx, &details, sql, whereValue); err != nil {
		return nil, err
	}

	return &types.GcpRouteListResult{Details: details}, nil
}

// BatchDeleteWithTx batch delete gcp route with transaction.
func (r *gcpRouteDao) BatchDeleteWithTx(kt *kit.Kit, tx *sqlx.Tx, filterExpr *filter.Expression) error {
	if filterExpr == nil {
		return errf.New(errf.InvalidParameter, "filter expr is required")
	}

	whereExpr, whereValue, err := filterExpr.SQLWhereExpr(tools.DefaultSqlWhereOption)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(`DELETE FROM %s %s`, table.GcpRouteTable, whereExpr)
	if _, err = r.orm.Txn(tx).Delete(kt.Ctx, sql, whereValue); err != nil {
		logs.ErrorJson("delete gcp route failed, err: %v, filter: %s, rid: %s", err, filterExpr, kt.Rid)
		return err
	}

	return nil
}
