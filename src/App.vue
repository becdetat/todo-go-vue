<template>
  <div id="app">
    <md-app>
      <md-app-toolbar class="md-primary md-large">
        <div class="md-toolbar-row">
          <span class="md-title">Things to do</span>
        </div>
      </md-app-toolbar>

      <md-app-content>
        <TodoList/>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
import TodoList from './components/TodoList.vue'

export default {
  name: 'app',
  components: {
    TodoList
  },
  data: () => ( {
  } )
}
</script>

<style>
#app {
  display: flex;
}
.md-app {
  width: 100%;
}
</style>
