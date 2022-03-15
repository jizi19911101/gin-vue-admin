<template>
  <div>
    <div class="gva-search-box">
      <!-- 初始版本自动化代码工具 -->
      <el-form ref="form" :rules="rules" :model="form" label-width="120px" >
        <el-form-item label="选择设备" prop="device">
          <!-- <el-input v-model="form.device" placeholder="请选择" :style="{width: '30%'}"/> -->
            <el-select v-model="form.device"  placeholder="请选择" size="small" clearable :style="{width: '30%'}">
                <el-option
                v-for="item in devices"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                >
                </el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="测试时长" prop="duration">
          <el-input v-model="form.duration" placeholder="设置测试时长（分钟）" :style="{width: '30%'}"/>
        </el-form-item>
        <!-- <el-tooltip content="注：请先确认手机安装了测试app" placement="bottom" effect="light"> -->
        <el-form-item label="测试app" prop="app">
        <!-- </el-tooltip> -->
            <!-- <el-input v-model="form.app" placeholder="请选择" /> -->
              <el-select v-model="form.app"  placeholder="请选择" size="small" clearable :style="{width: '30%'}">
                <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                >
                </el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="报告名称" prop="report">
          <el-input v-model="form.report" placeholder="报告名称" :style="{width: '30%'}"/>
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip content="注：会在app把之前跑出来的崩溃日志清除" placement="bottom" effect="light">
              <div> 是否清空日志 </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.cleanLog" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="startMonkey">发起测试</el-button>
          <el-button @click="reset">清空</el-button>
        </el-form-item>
      </el-form>
    </div>

  </div>
</template>

<script>
import { store } from '@/store'
import { getDeviceList, startMonkey } from '@/api/monkeyTest'
export default {
  name: 'StartMonkey',
  data() {
    return {
      form: {
        device: '',
        app: '',
        duration: '',
        report: '',
        cleanLog: false,
      },
      // options: ["触漫debug", "触漫release", "猫条debug", "猫条release"],
      options: [{ value: "com.mallestudio.gugu.app.debug",  label: "触漫debug" }, { value: "com.mallestudio.gugu.app",  label: "触漫release" }, { value: "cn.dreampix.dreamland.debug",  label: "猫条debug" }, { value: "22",  label: "猫条release" }],
      devices: [],
      rules: {
        device: [
          { required: true, message: '请选择设备', trigger: 'blur' }
        ],
        app: [
          { required: true, message: '请选择app', trigger: 'blur' }
        ],
        duration: [
          { required: true, message: '请设置测试时长', trigger: 'blur' }
        ],
        report: [
          { required: true, message: '请输入报告名称', trigger: 'blur' }
        ]
      },
    }
  },
  async created() {
        const userInfo = store.getters['user/userInfo']
        const res = await getDeviceList({ user_id: userInfo.userName+"@anonymous.com", "present":true })
        if (res.success === true) {
          var len = res.devices.length
          if (len != 0) {
            for (var i=0; i<len; i++){ 
              this.devices[i] = { value: res.devices[i].properties.serial,  label: res.devices[i].properties.name}
            }
          }
        }

  },
  methods: {
   reset() {
      // this.dialogFormVisible = false
      this.$refs.form.resetFields();
      this.form = {
        device: '',
        app: '',
        duration: '',
        report: '',
        cleanLog: false,
      }
    },
    // async startMonkey(){
    startMonkey(){
      let res
      this.$refs.form.validate( async (valid) => {
        if (!valid){

        }else{
            console.log(this.form)
            res = await startMonkey(this.form)
            if (res.code === 0) {
              this.$message({
                type: 'success',
                message: '发起monkey测试成功，请稍后查询报告'
              })
              this.reset()
          }
        }
       })

    }
  },
}
</script>

