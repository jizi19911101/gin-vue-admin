<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="接口名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="请求url">
          <el-input v-model="searchInfo.url" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="所属项目">
          <el-input v-model="searchInfo.project" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="所属模块">
          <el-input v-model="searchInfo.module" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="el-icon-delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180"  v-if="false">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="接口名称" prop="name" width="120" />
        <el-table-column align="left" label="请求方法" prop="method" width="120" />
        <el-table-column align="left" label="请求url" prop="url" width="360" />
        <el-table-column align="left" label="请求参数" prop="params" width="120" v-if="false"/>
        <el-table-column align="left" label="所属项目" prop="project" width="120" v-if="false"/>
        <el-table-column align="left" label="所属模块" prop="module" width="120" v-if="false"/>
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button" @click="updateApiInfo(scope.row)">变更</el-button>
            <el-button type="text" icon="el-icon-caret-right" size="small" class="table-button" @click="updateApiInfo(scope.row)">调试</el-button>
            <el-button type="text" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新增接口">
      <el-form :model="formData" :rules="rules" ref="formData" label-position="left" label-width="150px">
        <el-form-item label="接口名称:" prop="name">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="请求方法:" prop="method" >
          <el-select v-model="formData.method" placeholder="请下拉选择" clearable :style="{width: '100%'}">
            <el-option v-for="(item, index) in methodOptions" :key="index" :label="item"
              :value="item" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="请求url:"  prop="url">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>
            请求参数：
          <el-form-item
            v-for="(item, index) in formData.params"
            :label="'参数' + index + ':'"
            :key="index"
            :prop="formData.params[index]"
          >
            <el-input v-model="formData.params[index]" :style="{width: '80%'}"></el-input><el-button size="small" @click.prevent="removeDomain(index)"  :style="{width: '20%'}" icon="el-icon-delete" type="info"></el-button>
          </el-form-item>
          <el-form-item>
          <el-button @click="addDomain" size="small" :style="{float: 'right', width: '20%'}" type="primary">新增参数</el-button>
          </el-form-item>
        <el-form-item label="所属项目:" prop="project">
          <el-input v-model="formData.project" clearable placeholder="请输入" :style="{width: '100%'}"/>
        </el-form-item>
        <el-form-item label="所属模块:"  prop="module">
          <el-input v-model="formData.module" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog()">取 消</el-button>
          <el-button size="small" type="primary"  @click="enterDialog('formData')">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createApiInfo,
  deleteApiInfo,
  deleteApiInfoByIds,
  updateApiInfo,
  findApiInfo,
  getApiInfoList
} from '@/api/apiInfo' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'ApiInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getApiInfoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        name: '',
        method: '',
        url: '',
        params: [''],
        project: '',
        module: '',
      },
      rules: {
        name: [{ required: true, message: '请输入接口名称', trigger: 'blur' }],
        method: [{required: true, message: '请选择下拉选择',trigger: 'change'}],
        url: [{ required: true, message: '请输入url', trigger: 'blur' }],
        project: [{ required: true, message: '请输入所属项目', trigger: 'blur' }],
        module: [{ required: true, message: '请输入所属', trigger: 'blur' }],

      },
      methodOptions: ["get","post","put","delete"],
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
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteApiInfo(row)
      }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
      // removeDomain(item) {
      //   var index = this.formData.params.indexOf(item)
      //   if (index !== -1) {
      //     this.formData.params.splice(index, 1)
      //   }
      // },
      removeDomain(index) {
        this.formData.params.splice(index, 1)
      },
      addDomain() {
        this.formData.params.push('');
      },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteApiInfoByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateApiInfo(row) {
      const res = await findApiInfo({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.reapiInfo
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.$refs.formData.resetFields();
      this.formData = {
        name: '',
        method: '',
        url: '',
        params: [''],
        project: '',
        module: '',
      }
    },
    async deleteApiInfo(row) {
      const res = await deleteApiInfo({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog(formData) {
      let res
      this.$refs[formData].validate(async (valid) => {
        if (!valid) {
        }else {
          switch (this.type) {
              case 'create':
                res = await createApiInfo(this.formData)
                break
              case 'update':
                res = await updateApiInfo(this.formData)
                break
              default:
                res = await createApiInfo(this.formData)
                break
            }
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '创建/更改成功'
            })
            this.closeDialog()
            this.getTableData()
          }
        }
      });
    },

  
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}


</script>

<style>
  .el-dropdown-link {
    cursor: pointer;
    color: #409EFF;
  }
  .el-icon-arrow-down {
    font-size: 12px;
  }
</style>
