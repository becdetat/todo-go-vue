<template>
  <div>
    <md-list>
      <md-list-item>
        ☑
        <span class="md-list-item-text">
          todo
        </span>
      </md-list-item>
    </md-list>


    <ul>
      <li>items...</li>
      <li>new item form</li>
    </ul>
  </div>
</template>

<script>
// import draggable from 'vuedraggable'
import axios from 'axios'

export default {
  name: 'TodoList',
  data: () => ( {
    items: [],
    error: false
  } ),
  mounted() {
    axios
      .get('https://thawing-bayou-17829.herokuapp.com/todos')
      .then( response => {
        this.items = response.data
      } )
      .catch( error => {
      this.error = error
      } )
  }
}
</script>

<style scoped>
</style>
