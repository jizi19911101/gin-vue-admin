<template>
  <div id="app">
    <router-view />
            <!-- <component
            v-for="item in hasOpenComponentsArr"
            :key="item.name"
            :is="item.name"
            v-show="$route.path === item.path"
        ></component> -->


  </div>
</template>

<script>
export default {
  name: 'App',
  
  created() {
        console.log(this.$route.path)
        // // 设置iframe页的数组对象
        // const componentsArr = this.getComponentsArr();
        // componentsArr.forEach((item) => {
        //     Vue.component(item.name, item.component);
        // });
        // this.componentsArr = componentsArr;
        // // 判断当前路由是否iframe页
        // this.isOpenIframePage();
    },
  //   data() {
  //       return {
  //           componentsArr: [] // 含有iframe的页面
  //       }
  //   },
  //   watch: {
  //       $route() {
  //           // 判断当前路由是否iframe页
  //           this.isOpenIframePage();
  //       }
  //   },
  //   computed: {
  //       // 实现懒加载，只渲染已经打开过（hasOpen:true）的iframe页
  //       hasOpenComponentsArr() {
  //           return this.componentsArr.filter(item => item.hasOpen);
  //       }
  //   },
  //   methods: {
  //       // 根据当前路由设置hasOpen
  //       isOpenIframePage() {
  //           const target = this.componentsArr.find(item => {
  //               return item.path === this.$route.path
  //           });
  //           if (target && !target.hasOpen) {
  //               target.hasOpen = true;
  //           }
  //       },
  //       // 遍历路由的所有页面，把含有iframeComponent标识的收集起来
  //       getComponentsArr() {
  //           const router = this.$router;
  //           const routes = router.options.routes;
  //           const iframeArr = routes.filter(item => item.iframeComponent);
            
  //           return iframeArr.map((item) => {
  //               const name = item.name || item.path.replace('/', '');
  //               return {
  //                   name: name,
  //                   path: item.path,
  //                   hasOpen: false, // 是否打开过，默认false
  //                   component: item.iframeComponent // 组件文件的引用
  //               };
  //           });
  //       }
  //   }
}
</script>

<style lang="scss">
// 引入初始化样式
@import '@/style/main.scss';
@import '@/style/base.scss';
@import '@/style/mobile.scss';
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}
.el-button{
  font-weight: 400 !important;
}
</style>
