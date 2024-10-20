<template>
  <div>
    <el-card class="container-card" shadow="always">
      <el-form size="mini" :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-button :loading="loading" icon="el-icon-plus" type="warning" @click="create">新增</el-button>
        </el-form-item>
      </el-form>

      <el-table v-loading="loading" :data="tableData" border stripe style="width: 100%">
        <el-table-column show-overflow-tooltip prop="key" label="密钥">
          <template #default="{ row }">
            <i style="margin-left: 4px; cursor: pointer" class="el-icon-document-copy" @click="copyText(row.key)" />
            {{ row.key }}
          </template>
        </el-table-column>
        <el-table-column v-slot="scope" show-overflow-tooltip label="可重用" align="center">
          <el-tag size="small" :type="scope.row.reusable?'success':'danger'">{{ scope.row.reusable?"是":"否" }}</el-tag>
        </el-table-column>
        <el-table-column show-overflow-tooltip prop="ephemeral" label="短暂的" align="center">
          <template slot="header">
            <el-tooltip content="当设备离线时，自动将其从尾网中删除" placement="top">
              <span>短暂的<i class="el-icon-question" /></span>
            </el-tooltip>
          </template>
          <template v-slot="scope">
            <el-tag size="small" :type="scope.row.ephemeral?'success':'danger'">{{ scope.row.reusable?"是":"否" }}</el-tag>
          </template>

        </el-table-column>
        <el-table-column v-slot="scope" show-overflow-tooltip prop="used" label="已使用">
          <el-tag size="small" :type="scope.row.used?'danger':'success'">{{ scope.row.used?"是":"否" }}</el-tag>
        </el-table-column>
        <el-table-column v-slot="scope" show-overflow-tooltip sortable prop="expiration" label="过期时间">
          <el-tag size="small" :type="UtilsDateFormat.fromTimeStamp(scope.row.expiration).after()?'primary':'danger'">{{ UtilsDateFormat.fromTimeStamp(scope.row.expiration).toDateTimeString() }}</el-tag>
        </el-table-column>
        <el-table-column fixed="right" label="操作" align="center" width="120">
          <template slot-scope="scope">
            <el-tooltip class="delete-popover" content="过期" effect="dark" placement="top">
              <el-popconfirm title="确定过此期key吗？" @onConfirm="singleDelete(scope.row.key)">
                <el-button slot="reference" size="mini" icon="el-icon-delete" circle type="danger" />
              </el-popconfirm>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog title="新建密钥" :visible.sync="dialogFormVisible" width="30%">
        <el-form ref="dialogForm" size="small" :model="dialogFormData" label-width="100px">
          <el-form-item label="过期时间">
            <el-date-picker
              v-model="dialogFormData.expireDate"
              type="datetime"
              :picker-options="pickerOptions"
              placeholder="选择时间"
              align="bottom"
            />
          </el-form-item>
          <el-form-item>
            <el-checkbox v-model="dialogFormData.reusable" label="可重用" />
            <el-checkbox v-model="dialogFormData.ephemeral">
              <el-tooltip content="当设备离线时，自动将其从尾网中删除" placement="top">
                <span>短暂的<i class="el-icon-question" /></span>
              </el-tooltip>
            </el-checkbox>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="mini" @click="cancelForm()">取 消</el-button>
          <el-button size="mini" :loading="submitLoading" type="primary" @click="submitForm()">确 定</el-button>
        </div>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
import { getKeys, createKey, expireKey } from '@/api/headscale/keys'
import { UtilsDateFormat } from '@/utils/date'

export default {
  name: 'User',
  data() {
    const today = new Date() // 获取当前时间
    return {
      // 表格数据
      tableData: [],
      loading: false,
      // dialog对话框
      submitLoading: false,
      dialogFormVisible: false,
      dialogFormData: {
        expireDate: new Date(today.getFullYear(), today.getMonth(), today.getDate(), 23, 59, 59).toISOString(),
        reusable: false,
        expiration: {
          seconds: 0,
          nanos: 0
        },
        ephemeral: false
      },
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() < (Date.now() - 24 * 60 * 60 * 1000)
        }
      }
    }
  },
  computed: {
    UtilsDateFormat() {
      return UtilsDateFormat
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    // 复制
    copyText(msg) {
      const save = function(e) {
        e.clipboardData.setData('text/plain', msg)
        e.preventDefault() // 阻止默认行为
      }
      document.addEventListener('copy', save) // 添加一个copy事件
      document.execCommand('copy') // 执行copy方法
      this.$message({ message: '复制成功', type: 'success' })
    },
    // 获取表格数据
    async getTableData() {
      this.loading = true
      try {
        const { data } = await getKeys()
        this.tableData = data.keys
      } finally {
        this.loading = false
      }
    },

    // 新增
    create() {
      this.dialogFormVisible = true
    },

    // 提交表单
    async submitForm() {
      this.submitLoading = true
      const date = new Date(this.dialogFormData.expireDate)
      const seconds = Math.floor(date.getTime() / 1000)
      this.dialogFormData.expiration.seconds = seconds
      this.dialogFormData.expiration.nanos = 0
      await createKey(this.dialogFormData).then(res => {
        this.$message({
          showClose: true,
          message: '创建成功',
          type: 'success'
        })
      }).finally(() => {
        this.resetForm()
        this.getTableData()
        this.submitLoading = false
      })
    },

    // 提交表单
    cancelForm() {
      this.resetForm()
    },

    resetForm() {
      this.dialogFormVisible = false
      const today = new Date() // 获取当前时间
      this.dialogFormData = {
        expireDate: new Date(today.getFullYear(), today.getMonth(), today.getDate(), 23, 59, 59).toISOString(),
        expiration: {
          seconds: 0,
          nanos: 0
        },
        reusable: false,
        ephemeral: false
      }
    },

    // 单个删除
    async singleDelete(key) {
      this.loading = true
      await expireKey({ key: key }).then(res => {
        this.$message({
          showClose: true,
          message: '成功',
          type: 'success'
        })
      }).finally(() => {
        this.resetForm()
        this.getTableData()
        this.submitLoading = false
      })
    }
  }
}
</script>

<style scoped>
  .container-card{
    margin: 10px;
  }

  .delete-popover{
    margin-left: 10px;
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 3px;
    font-size: 16px;
    color: #889aa4;
    cursor: pointer;
    user-select: none;
  }
</style>
