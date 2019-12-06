<template>
  <div id="app">
    <h1>Things to do</h1>
    <TodoList/>
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
</style>
