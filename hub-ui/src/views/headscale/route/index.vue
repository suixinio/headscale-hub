<template>
  <div>
    <el-card class="container-card" shadow="always">
      <el-table v-loading="loading" :data="tableData" border stripe style="width: 100%">
        <el-table-column show-overflow-tooltip sortable prop="prefix" label="路由" />
        <el-table-column show-overflow-tooltip sortable prop="nickname" label="VIA">
          <template v-slot="scope">
            <div v-for="ip in scope.row.node.ip_addresses" :key="ip">{{ ip }}</div>
          </template>
        </el-table-column>
        <el-table-column show-overflow-tooltip sortable prop="node.given_name" label="设备" align="center" />
        <el-table-column show-overflow-tooltip sortable prop="node.name" label="主机名" />
        <el-table-column show-overflow-tooltip sortable prop="node.user.name" label="创建人" />
        <el-table-column show-overflow-tooltip sortable prop="creator" label="状态">
          <template v-slot="scope">
            <el-switch v-model="scope.row.enabled" active-color="#13ce66" inactive-color="#ff4949" @change="switchStatus($event, scope.row.id)" />
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
    </el-card>
  </div>
</template>

<script>
import { getRouteList, enableStatus, deleteRoute } from '@/api/headscale/route'

export default {
  data() {
    return {
      // 查询参数
      params: {
        username: '',
        nickname: '',
        status: '',
        mobile: '',
        pageNum: 1,
        name: 'User',
        pageSize: 10
      },
      // 表格数据
      tableData: [],
      loading: false,

      // dialog对话框
      submitLoading: false,
      dialogFormTitle: '',
      dialogType: '',
      dialogFormVisible: false,
      dialogFormData: {
        username: '',
        password: '',
        nickname: '',
        status: 1,
        mobile: '',
        avatar: '',
        introduction: '',
        roleIds: ''
      },

      // 删除按钮弹出框
      popoverVisible: false
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    // 获取表格数据
    async getTableData() {
      this.loading = true
      try {
        const { data } = await getRouteList()
        this.tableData = data.routes
      } finally {
        this.loading = false
      }
    },

    // 新增
    create() {
      this.dialogFormVisible = true
    },

    // 修改
    update(row) {
      this.dialogFormData.ID = row.ID
      this.dialogFormData.username = row.username
      this.dialogFormData.password = ''
      this.dialogFormData.nickname = row.nickname
      this.dialogFormData.status = row.status
      this.dialogFormData.mobile = row.mobile
      this.dialogFormData.introduction = row.introduction
      this.dialogFormData.roleIds = row.roleIds
      this.dialogFormVisible = true
    },

    async switchStatus(event, routeId) {
      console.log(event, routeId)
      await enableStatus({ route_id: routeId, enable: event }).then(res => {
        this.$message({
          showClose: true,
          message: '变更成功',
          type: 'success'
        })
      })
    },

    resetForm() {
      this.dialogFormVisible = false
      this.$refs['dialogForm'].resetFields()
      this.dialogFormData = {
        username: '',
        password: '',
        nickname: '',
        status: 1,
        mobile: '',
        avatar: '',
        introduction: '',
        roleIds: ''
      }
    },

    // 单个删除
    async singleDelete(id) {
      this.loading = true
      await deleteRoute({ route_id: id }).then(res => {
        this.$message({
          showClose: true,
          message: '删除成功',
          type: 'success'
        })
        this.getTableData()
        this.loading = false
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
