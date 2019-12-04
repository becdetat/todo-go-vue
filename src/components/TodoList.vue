<template>
  <div>
    <md-list>
      <md-list-item v-for="item in items" v-bind:key="item.id">
        <span v-if="item.complete">☑</span>
        <span v-if="!item.complete">⍻</span>
        <span class="md-list-item-text">
          {{item.title}}
        </span>
      </md-list-item>
    </md-list>
  </div>
</template>

<script>
// import draggable from 'vuedraggable'
import axios from 'axios'
import _ from 'lodash'

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
        this.items = _.orderBy(response.data, x => x.position)
      } )
      .catch( error => {
        this.error = error
      } )
  }
}
</script>

<style scoped>
</style>
