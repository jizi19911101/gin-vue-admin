<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="报告名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">

        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="报告名称" prop="name" width="200" />
        <el-table-column align="left" label="报告地址" prop="url" width="300" />
        <el-table-column align="left" label="测试环境" prop="env" width="100" />
        <el-table-column align="left" label="报告描述" prop="description" width="250" />
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="text" icon="el-icon-data-analysis" size="small" class="table-button" @click="viewReport(scope.row)">查看</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>

  </div>
</template>

<script>
import {
  getReportList
} from '@/api/apiTest' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Report',
  mixins: [infoList],
  data() {
    return {
      listApi: getReportList,
      multipleSelection: [],
      formData: {
        name: '',
      },

    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
  onReset() {
    this.searchInfo = {}
  },
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    viewReport(item){
      window.open(item.url)
    }
  },
}
</script>

<style>
</style>
