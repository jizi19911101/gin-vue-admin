<template>
  <!-- <iframe style="width: 100%;" :src="url" class="iframe" frameborder="0" scrolling="no"></iframe> -->
  <iframe style="width: 100%;" :src="url" width="400" height="600" frameborder="0" scrolling="no"></iframe>
</template>

<script>
import { store } from '@/store'
export default {
  name: 'Devices',
  data () {
    return {
      url: "http://127.0.0.1:4000/login",
    }
  },
  methods: {
    adds:function(userInfo) {
       this.url = this.url+"?name="+userInfo.userName; 
  }
},
  mounted() { 
    const userInfo = store.getters['user/userInfo']
    this.adds(userInfo);
  } 
}
</script>
