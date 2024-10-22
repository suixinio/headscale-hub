<template>
  <div>
    <el-card class="container-card" shadow="always">
      <el-form size="mini" :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-button :loading="loading" icon="el-icon-plus" type="success" @click="create">注册节点</el-button>
        </el-form-item>
      </el-form>

      <el-table v-loading="loading" :data="tableData" border stripe style="width: 100%">
        <el-table-column show-overflow-tooltip prop="name" label="设备名" />
        <el-table-column show-overflow-tooltip prop="ip" label="IP地址">
          <template v-slot="scope">
            <div v-for="ip in scope.row.ip_addresses" :key="ip">{{ ip }}</div>
          </template>
        </el-table-column>
        <el-table-column show-overflow-tooltip prop="status" label="最后连接时间" align="center">
          <template v-slot="scope">
            <el-tag size="small" :type="scope.row.online ? 'success':'danger'">{{ scope.row.online === true ? '已连接': formatAfterDateTime(scope.row.last_seen) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column show-overflow-tooltip prop="user.name" label="所有人" />
        <el-table-column show-overflow-tooltip sortable prop="expiry" label="过期时间">
          <template v-slot="scope">
            <el-tag size="small" :type="UtilsDateFormat.fromTimeStamp(scope.row.expiry).after()?'primary':'danger'">{{ UtilsDateFormat.fromTimeStamp(scope.row.expiry).isMin()?'永不过期':UtilsDateFormat.fromTimeStamp(scope.row.expiry).toDateTimeString() }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" align="center" width="120">
          <template slot-scope="scope">
            <el-tooltip class="delete-popover" content="删除" effect="dark" placement="top">
              <el-popconfirm title="确定删除吗？" @onConfirm="singleDelete(scope.row.id)">
                <el-button slot="reference" size="mini" icon="el-icon-delete" circle type="danger" />
              </el-popconfirm>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <el-dialog title="注册节点" :visible.sync="dialogFormVisible" width="30%">
        <el-form ref="dialogForm" size="small" :model="dialogFormData" label-width="100px">
          <el-form-item label="nodekey" prop="nodekey">
            <el-input ref="nodekey" v-model.trim="dialogFormData.nodeKey" placeholder="nodekey" />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="mini" @click="cancelForm()">取 消</el-button>
          <el-button size="mini" :loading="submitLoading" type="primary" @click="registerNodeBtn()">确 定</el-button>
        </div>
      </el-dialog>

    </el-card>
  </div>
</template>

<script>
import { getNodeList, registerNode, deleteNode } from '@/api/headscale/node'
import { formatAfterDateTime, UtilsDateFormat } from '@/utils/date'

export default {
  name: 'User',
  data() {
    return {
      // 表格数据
      tableData: [],
      loading: false,
      // dialog对话框
      submitLoading: false,
      dialogFormTitle: '',
      dialogType: '',
      dialogFormVisible: false,
      dialogFormData: {
        nodeKey: ''
      },
      // 删除按钮弹出框
      popoverVisible: false
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
    formatAfterDateTime,
    // 获取表格数据
    async getTableData() {
      this.loading = true
      try {
        const { data } = await getNodeList()
        this.tableData = data.nodes
      } finally {
        this.loading = false
      }
    },

    // 新增
    create() {
      this.dialogFormVisible = true
    },

    // 提交表单
    async registerNodeBtn() {
      this.submitLoading = true
      await registerNode({ key: this.dialogFormData.nodeKey }).then(res => {
        this.$message({
          showClose: true,
          message: '注册成功',
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
      this.dialogFormData = {
        nodeKey: ''
      }
    },

    // 单个删除
    async singleDelete(Id) {
      this.loading = true
      await deleteNode({ node_id: Id }).then(res => {
        this.$message({
          showClose: true,
          message: '删除成功',
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
</style>
