/*
Copyright 2022 Huawei Cloud Computing Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package record_test

import (
	"sort"
	"testing"

	"github.com/openGemini/openGemini/engine/mutable"
	"github.com/openGemini/openGemini/lib/config"
	"github.com/openGemini/openGemini/lib/record"
	"github.com/openGemini/openGemini/open_src/vm/protoparser/influx"
)

func TestSortRecordByOrderTags(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "ttbool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 1, 1, 1}, []int64{1000, 0, 0, 1100},
		[]int{1, 1, 1, 1}, []float64{1001.3, 1002.4, 1002.4, 0},
		[]int{1, 1, 1, 1}, []string{"ha", "helloNew", "worldNew", "hb"},
		[]int{1, 1, 1, 1}, []bool{true, true, false, false},
		[]int64{22, 21, 20, 19})
	expRec := genRowRec(schema,
		[]int{1, 1, 1, 1}, []int64{0, 0, 1000, 1100},
		[]int{1, 1, 1, 1}, []float64{1002.4, 1002.4, 1001.3, 0},
		[]int{1, 1, 1, 1}, []string{"helloNew", "worldNew", "ha", "hb"},
		[]int{1, 1, 1, 1}, []bool{true, false, true, false},
		[]int64{21, 20, 22, 19})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}
	pk := []string{"order1_int", "order2_float"}
	sk := []string{"order1_int", "order2_float", "order3_string"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)
	tbl.SetMsInfo("cpu", msInfo)
	tbl.MTable.SortAndDedup(tbl, "cpu", nil)
	if !testRecsEqual(msInfo.GetWriteChunk().WriteRec.GetRecord(), expRec) {
		t.Fatal("error result")
	}
}

func TestSortRecordByOrderTags1(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "ttbool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int64{700, 600, 500, 400, 300, 200, 100, 30, 20, 10, 9, 8},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []float64{0, 5.3, 4.3, 0, 3.3, 0, 2.3, 0, 1003.5, 0, 1002.4, 1001.3},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []string{"test", "", "world", "", "", "hello", "", "", "testNew1", "worldNew", "helloNew", "a"},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []bool{false, false, true, false, true, false, false, false, true, false, true, true},
		[]int64{1, 2, 3, 4, 5, 6, 7, 18, 19, 20, 21, 22})
	expRec := genRowRec(schema,
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int64{8, 9, 10, 20, 30, 100, 200, 300, 400, 500, 600, 700},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []float64{1001.3, 1002.4, 0, 1003.5, 0, 2.3, 0, 3.3, 0, 4.3, 5.3, 0},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []string{"a", "helloNew", "worldNew", "testNew1", "", "", "hello", "", "", "world", "", "test"},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []bool{true, true, false, true, false, false, false, true, false, true, false, false},
		[]int64{22, 21, 20, 19, 18, 7, 6, 5, 4, 3, 2, 1})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}
	pk := []string{"order1_int", "order2_float", "order3_string"}
	sk := []string{"order1_int", "order2_float", "order3_string"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)
	tbl.SetMsInfo("cpu", msInfo)
	tbl.MTable.SortAndDedup(tbl, "cpu", nil)
	if !testRecsEqual(msInfo.GetWriteChunk().WriteRec.GetRecord(), expRec) {
		t.Fatal("error result")
	}
}

func TestSortRecordByOrderTags2(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "ttbool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 1, 1, 1}, []int64{1000, 0, 0, 1100},
		[]int{1, 1, 1, 1}, []float64{1001.3, 1002.4, 1002.4, 0},
		[]int{1, 1, 1, 1}, []string{"ha", "helloNew", "worldNew", "hb"},
		[]int{1, 1, 1, 1}, []bool{true, true, false, false},
		[]int64{22, 21, 20, 19})
	expRec := genRowRec(schema,
		[]int{1, 1, 1, 1}, []int64{0, 0, 1000, 1100},
		[]int{1, 1, 1, 1}, []float64{1002.4, 1002.4, 1001.3, 0},
		[]int{1, 1, 1, 1}, []string{"worldNew", "helloNew", "ha", "hb"},
		[]int{1, 1, 1, 1}, []bool{false, true, true, false},
		[]int64{20, 21, 22, 19})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}
	pk := []string{"order1_int"}
	sk := []string{"order1_int", "ttbool"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)
	tbl.SetMsInfo("cpu", msInfo)
	tbl.MTable.SortAndDedup(tbl, "cpu", nil)
	if !testRecsEqual(msInfo.GetWriteChunk().WriteRec.GetRecord(), expRec) {
		t.Fatal("error result")
	}
}

func TestSortRecordByOrderTags3(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "order4_bool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 0, 0, 1}, []int64{1000, 0, 0, 1100},
		[]int{1, 1, 1, 0}, []float64{1001.3, 1002.4, 1003.4, 0},
		[]int{1, 1, 1, 0}, []string{"ha", "helloNew", "helloNew", ""},
		[]int{1, 1, 0, 1}, []bool{true, true, false, false},
		[]int64{22, 21, 20, 19})
	expRec := genRowRec(schema,
		[]int{1, 1, 0, 0}, []int64{1100, 1000, 0, 0},
		[]int{0, 1, 1, 1}, []float64{0, 1001.3, 1002.4, 1003.4},
		[]int{0, 1, 1, 1}, []string{"", "ha", "helloNew", "helloNew"},
		[]int{1, 1, 1, 0}, []bool{false, true, true, false},
		[]int64{19, 22, 21, 20})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}
	pk := []string{"order3_string", "order2_float"}
	sk := []string{"order3_string", "order2_float"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)
	tbl.SetMsInfo("cpu", msInfo)
	tbl.MTable.SortAndDedup(tbl, "cpu", nil)
	if !testRecsEqual(msInfo.GetWriteChunk().WriteRec.GetRecord(), expRec) {
		t.Fatal("error result")
	}
}

func TestSortRecordByOrderTags4(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "order4_bool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 0, 0, 1}, []int64{1000, 0, 0, 1100},
		[]int{1, 1, 1, 0}, []float64{1001.3, 1002.4, 1002.4, 0},
		[]int{1, 0, 0, 0}, []string{"ha", "", "", ""},
		[]int{1, 1, 0, 1}, []bool{true, false, true, false},
		[]int64{22, 21, 21, 19})
	expRec := genRowRec(schema,
		[]int{0, 0, 1, 1}, []int64{0, 0, 1000, 1100},
		[]int{1, 1, 1, 0}, []float64{1002.4, 1002.4, 1001.3, 0},
		[]int{0, 0, 1, 0}, []string{"", "", "ha", ""},
		[]int{0, 1, 1, 1}, []bool{true, false, true, false},
		[]int64{21, 21, 22, 19})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}
	pk := []string{"order1_int", "order2_float", "order3_string"}
	sk := []string{"order1_int", "order2_float", "order3_string", "order4_bool"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)
	tbl.SetMsInfo("cpu", msInfo)
	tbl.MTable.SortAndDedup(tbl, "cpu", nil)
	if !testRecsEqual(msInfo.GetWriteChunk().WriteRec.GetRecord(), expRec) {
		t.Fatal("error result")
	}
}

func TestSortRecordByOrderTags5(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "order4_bool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 0, 0, 1}, []int64{1000, 0, 0, 1100},
		[]int{1, 1, 1, 0}, []float64{1001.3, 1002.4, 1002.4, 0},
		[]int{1, 1, 1, 0}, []string{"ha", "helloNew", "worldNew", ""},
		[]int{1, 1, 1, 1}, []bool{true, true, false, false},
		[]int64{22, 21, 21, 19})
	expRec := genRowRec(schema,
		[]int{0, 0, 1, 1}, []int64{0, 0, 1000, 1100},
		[]int{1, 1, 1, 0}, []float64{1002.4, 1002.4, 1001.3, 0},
		[]int{1, 1, 1, 0}, []string{"worldNew", "helloNew", "ha", ""},
		[]int{1, 1, 1, 1}, []bool{false, true, true, false},
		[]int64{21, 21, 22, 19})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}
	pk := []string{"order1_int", "order4_bool"}
	sk := []string{"order1_int", "order4_bool"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)
	tbl.SetMsInfo("cpu", msInfo)
	tbl.MTable.SortAndDedup(tbl, "cpu", nil)
	if !testRecsEqual(msInfo.GetWriteChunk().WriteRec.GetRecord(), expRec) {
		t.Fatal("error result")
	}
}

func TestSortRecordByOrderTags6(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "order4_bool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 0, 0, 1}, []int64{1000, 0, 0, 1100},
		[]int{1, 1, 1, 0}, []float64{1001.3, 1002.4, 1002.4, 0},
		[]int{1, 1, 1, 0}, []string{"ha", "worldNew", "helloNew", "a"},
		[]int{1, 1, 1, 1}, []bool{true, false, false, false},
		[]int64{22, 21, 21, 19})
	expRec := genRowRec(schema,
		[]int{1, 1, 0, 0}, []int64{1100, 1000, 0, 0},
		[]int{0, 1, 1, 1}, []float64{0, 1001.3, 1002.4, 1002.4},
		[]int{0, 1, 1, 1}, []string{"a", "ha", "helloNew", "worldNew"},
		[]int{1, 1, 1, 1}, []bool{false, true, false, false},
		[]int64{19, 22, 21, 21})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}
	pk := []string{"order3_string", "order4_bool"}
	sk := []string{"order3_string", "order4_bool"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)
	tbl.SetMsInfo("cpu", msInfo)
	tbl.MTable.SortAndDedup(tbl, "cpu", nil)
	if !testRecsEqual(msInfo.GetWriteChunk().WriteRec.GetRecord(), expRec) {
		t.Fatal("error result")
	}
}

func TestSortRecordByOrderTags7(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "order4_bool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 0, 0, 1}, []int64{1000, 0, 0, 1100},
		[]int{1, 1, 1, 0}, []float64{1001.3, 1002.4, 1002.4, 0},
		[]int{1, 1, 1, 0}, []string{"ha", "worldNew", "worldNew", ""},
		[]int{1, 1, 1, 1}, []bool{true, false, false, false},
		[]int64{22, 21, 21, 19})
	expRec := genRowRec(schema,
		[]int{1, 0, 0, 1}, []int64{1100, 0, 0, 1000},
		[]int{0, 1, 1, 1}, []float64{0, 1002.4, 1002.4, 1001.3},
		[]int{0, 1, 1, 1}, []string{"", "worldNew", "worldNew", "ha"},
		[]int{1, 1, 1, 1}, []bool{false, false, false, true},
		[]int64{19, 21, 21, 22})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}

	pk := []string{"time"}
	sk := []string{"time"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)

	tbl.SetMsInfo("cpu", msInfo)
	tbl.MTable.SortAndDedup(tbl, "cpu", nil)
	if !testRecsEqual(msInfo.GetWriteChunk().WriteRec.GetRecord(), expRec) {
		t.Fatal("error result")
	}
}

func TestBoolSliceSingleValueCompare(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "order4_bool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1}, []int64{1000},
		[]int{1}, []float64{1001.3},
		[]int{1}, []string{"ha"},
		[]int{1}, []bool{true},
		[]int64{22})
	cmpRec := genRowRec(schema,
		[]int{1}, []int64{1100},
		[]int{0}, []float64{0},
		[]int{0}, []string{"worldNew"},
		[]int{1}, []bool{false},
		[]int64{19})
	expResult := []int{1, -1, -1, -1}
	sort.Sort(rec)
	sort.Sort(cmpRec)
	times := rec.Times()
	pk := []record.PrimaryKey{{"order1_int", influx.Field_Type_Int}, {"order2_float", influx.Field_Type_Float},
		{"order3_string", influx.Field_Type_String}, {"order4_bool", influx.Field_Type_Boolean}}
	sk := pk
	data := record.SortData{}
	dataCmp := record.SortData{}
	data.Init(times, pk, sk, rec)
	dataCmp.Init(cmpRec.Times(), pk, sk, cmpRec)

	im := data.Data
	jm := dataCmp.Data
	res := make([]int, 0, len(pk))
	for idx := 0; idx < len(im); idx++ {
		v, err := im[idx].CompareSingleValue(jm[idx], 0, 0)
		if err != nil {
			t.Fatal()
		}
		res = append(res, v)
	}
	if !compareResult(res, expResult) {
		t.Fatal()
	}
}

func TestBoolSliceSingleValueCompare2(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "order4_bool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1}, []int64{1000},
		[]int{1}, []float64{1001.3},
		[]int{1}, []string{"ha"},
		[]int{1}, []bool{true},
		[]int64{22})
	cmpRec := genRowRec(schema,
		[]int{1}, []int64{1100},
		[]int{1}, []float64{200},
		[]int{1}, []string{"world"},
		[]int{1}, []bool{true},
		[]int64{19})
	expResult := []int{1, -1, 1, 0}
	sort.Sort(rec)
	sort.Sort(cmpRec)
	times := rec.Times()
	pk := []record.PrimaryKey{{"order1_int", influx.Field_Type_Int}, {"order2_float", influx.Field_Type_Float},
		{"order3_string", influx.Field_Type_String}, {"order4_bool", influx.Field_Type_Boolean}}
	sk := pk
	data := record.SortData{}
	dataCmp := record.SortData{}
	data.Init(times, pk, sk, rec)
	dataCmp.Init(cmpRec.Times(), pk, sk, cmpRec)

	im := data.Data
	jm := dataCmp.Data
	res := make([]int, 0, len(pk))
	for idx := 0; idx < len(im); idx++ {
		v, err := im[idx].CompareSingleValue(jm[idx], 0, 0)
		if err != nil {
			t.Fatal()
		}
		res = append(res, v)
	}
	if !compareResult(res, expResult) {
		t.Fatal()
	}
}

func TestBoolSliceSingleValueCompare3(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "order4_bool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1}, []int64{1100},
		[]int{1}, []float64{1001.3},
		[]int{1}, []string{"hb"},
		[]int{0}, []bool{false},
		[]int64{22})
	cmpRec := genRowRec(schema,
		[]int{1}, []int64{1100},
		[]int{1}, []float64{1001.3},
		[]int{1}, []string{"ha"},
		[]int{1}, []bool{false},
		[]int64{19})
	expResult := []int{0, 0, -1, 1}
	sort.Sort(rec)
	sort.Sort(cmpRec)
	times := rec.Times()
	pk := []record.PrimaryKey{{"order1_int", influx.Field_Type_Int}, {"order2_float", influx.Field_Type_Float},
		{"order3_string", influx.Field_Type_String}, {"order4_bool", influx.Field_Type_Boolean}}
	sk := pk
	data := record.SortData{}
	dataCmp := record.SortData{}
	data.Init(times, pk, sk, rec)
	dataCmp.Init(cmpRec.Times(), pk, sk, cmpRec)

	im := data.Data
	jm := dataCmp.Data
	res := make([]int, 0, len(pk))
	for idx := 0; idx < len(im); idx++ {
		v, err := im[idx].CompareSingleValue(jm[idx], 0, 0)
		if err != nil {
			t.Fatal()
		}
		res = append(res, v)
	}
	if !compareResult(res, expResult) {
		t.Fatal()
	}
}

func compareResult(res, expRes []int) bool {
	if len(res) != len(expRes) {
		return false
	}
	if len(res) == 0 {
		return false
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expRes[i] {
			return false
		}
	}
	return true
}

func TestSortRecordAndDeduplicate(t *testing.T) {
	schema := record.Schemas{
		record.Field{Type: influx.Field_Type_Int, Name: "order1_int"},
		record.Field{Type: influx.Field_Type_Float, Name: "order2_float"},
		record.Field{Type: influx.Field_Type_String, Name: "order3_string"},
		record.Field{Type: influx.Field_Type_Boolean, Name: "ttbool"},
		record.Field{Type: influx.Field_Type_Int, Name: "time"},
	}
	rec := genRowRec(schema,
		[]int{1, 1, 1, 1}, []int64{1000, 0, 0, 1100},
		[]int{1, 1, 1, 1}, []float64{1001.3, 1002.4, 1002.4, 0},
		[]int{1, 1, 1, 1}, []string{"ha", "helloNew", "helloNew", "hb"},
		[]int{1, 1, 1, 1}, []bool{true, true, false, false},
		[]int64{22, 21, 20, 19})
	expRec := genRowRec(schema,
		[]int{1, 1, 1}, []int64{0, 1000, 1100},
		[]int{1, 1, 1}, []float64{1002.4, 1001.3, 0},
		[]int{1, 1, 1}, []string{"helloNew", "ha", "hb"},
		[]int{1, 1, 1}, []bool{false, true, false},
		[]int64{20, 22, 19})
	sort.Sort(rec)
	sort.Sort(expRec)

	tbl := mutable.NewMemTable(config.COLUMNSTORE)
	msInfo := &mutable.MsInfo{
		Name:   "cpu",
		Schema: schema,
	}
	pk := []string{"order1_int", "order2_float"}
	sk := []string{"order1_int", "order2_float", "order3_string"}
	msInfo.CreateWriteChunkForColumnStore(pk, sk)
	wk := msInfo.GetWriteChunk()
	wk.WriteRec.SetWriteRec(rec)
	tbl.SetMsInfo("cpu", msInfo)

	hlp := record.NewSortHelper()
	defer hlp.Release()
	rec = hlp.SortForColumnStore(wk.WriteRec.GetRecord(), hlp.SortData, mutable.GetPrimaryKeys(schema, pk), mutable.GetPrimaryKeys(schema, sk), true)

	if !testRecsEqual(rec, expRec) {
		t.Fatal("error result")
	}
}
