<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="接口名称">
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
        <el-table-column align="left" label="接口名称" prop="name" width="250" />
        <!-- <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="text" icon="el-icon-files" size="small" class="table-button" @click="goToCaseList(scope.row)">查看用例</el-button>
            <el-button type="text" icon="el-icon-data-analysis" size="small" class="table-button" @click="goToReport(scope.row)">测试报告</el-button>
            <el-button type="text" icon="el-icon-s-tools" size="small" class="table-button" @click="changeVisible(scope.row)">测试</el-button>
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
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="跑测试用例">
          <el-form :model="formData" :rules="rules" ref="formData" label-position="right" label-width="100px">
            <el-form-item label="测试环境" prop="env">
              <el-select v-model="formData.env" placeholder="请下拉选择" clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in envOptions" :key="index"  :value="item"
                :disabled="item.disabled"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="测试模块:" prop="module">
              <el-input v-model="formData.module" readonly placeholder="请输入" />
            </el-form-item>
            <el-form-item label="测试接口:" prop="api">
              <el-input v-model="formData.api" readonly placeholder="请输入" />
            </el-form-item>
          </el-form>
          <template #footer>
            <div class="dialog-footer">
              <el-button size="small" @click="closeDialog">取 消</el-button>
              <el-button size="small" type="primary" @click="runCase()">确 定</el-button>
            </div>
          </template>
        </el-dialog>
    </div>

  </div>
</template>

<script>
import {
  getApiList,
  runCase
} from '@/api/apiTest' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'ApiList',
  mixins: [infoList],
  data() {
    return {
      dialogFormVisible: false,
      listApi: getApiList,
      multipleSelection: [],
      formData: {
        module: '',
        api: '',
        env:''
      },
      envOptions: ["demo","api2","api"],
      rules: {
        env: [{ required:true, message: "请选择测试环境", trigger:"blur" }],
      }
    }
  },
  async created() {
    this.searchInfo.module = this.$route.query.module
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

    goToCaseList(item){
      this.$router.push({name:"case", query:{api:item.name, module:item.module}})
    },
    goToReport(item){
      this.$router.push({name:"report",})
    },
    async changeVisible(row) {
        this.dialogFormVisible = true
        this.formData.api =  row.name
        this.formData.module =  row.module
      
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.$refs.formData.resetFields();
      this.formData = {
        module: '',
        api:''
      }
    },
    async runCase(){
      let res
      this.$refs.formData.validate( async (valid) => {
        if (!valid){

        }else{
            res = await runCase(this.formData)
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '发起测试成功，请稍后查询报告'
            })
            this.closeDialog()
          }
        }
       })

    }
  },
}
</script>

<style>
</style>
