<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用例名称">
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
        <el-table-column align="left" label="所属模块" prop="module" width="200" />
        <el-table-column align="left" label="所属接口" prop="api" width="250" />
        <el-table-column align="left" label="用例名称" prop="name" width="250" />
        <!-- <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="text" icon="el-icon-data-analysis" size="small" class="table-button" @click="goToReport(scope.row)">测试报告</el-button>
            <el-button type="text" icon="el-icon-s-tools" size="small" class="table-button" @click="updateOrganization(scope.row)">测试</el-button>
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
  getCaseList
} from '@/api/apiTest' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Case',
  mixins: [infoList],
  data() {
    return {
      listApi: getCaseList,
      multipleSelection: [],
      formData: {
        name: '',
      },
    }
  },
  async created() {
    this.searchInfo.module = this.$route.query.module
    this.searchInfo.api = this.$route.query.api
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
    goToReport(item){
      this.$router.push({name:"report",})
    },
  },
}
</script>

<style>
</style>
