<template>
  <el-card style="margin: 10px">
    <div slot="header">
      <span>ACLs</span>
    </div>
    <YamlEditor v-model="acl" theme="rubyblue" />
    <el-row type="flex" justify="end">
      <el-col :span="1">
        <el-button type="success" style="margin-top: 20px" @click="submit">保存</el-button>
      </el-col>
    </el-row>
  </el-card>
</template>

<script>
import YamlEditor from '@/components/YamlEditor/index.vue'
import { getACL, setACL } from '@/api/headscale/acl'

export default {
  name: 'ACL',
  components: { YamlEditor },
  data() {
    return {
      autoReload: false,
      acl: ''
    }
  },
  created() {
    this.getData()
  },
  methods: {
    submit() {
      if (!this.autoReload) {
        this.confirm(this.postData)
      } else {
        this.postData()
      }
    },
    confirm(func) {
      this.$confirm('此操作将修改Tailscale访权限问, 是否继续?', '注意', {
        type: 'warning'
      }).then(() => {
        func()
      }).catch(() => {
      })
    },
    async postData() {
      await setACL({ content: this.acl }).then(res => {
        this.$message.success('保存成功')
      }).finally(
        this.getData()
      )
    },
    async getData() {
      const { code, data, message } = await getACL()
      if (code !== 200) {
        this.$message.error(message)
        return
      }
      this.acl = data.data
    }
  }
}
</script>

<style scoped>

</style>
