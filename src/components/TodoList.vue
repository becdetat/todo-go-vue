<template>
  <div>
    <md-list>
      <md-list-item v-for="item in items" v-bind:key="item.id">
        <md-checkbox v-model="item.complete" v-on:change="updateTodo( item )"></md-checkbox>
        <span class="md-list-item-text">
          <span v-if="!item.complete">{{item.title}}</span>
          <strike v-if="item.complete">{{item.title}}</strike>
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
    error: null
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
  },
  methods: {
    updateTodo: ( todo ) => {
      axios
        .put(`https://thawing-bayou-17829.herokuapp.com/todos/${todo.id}`, todo)
        .catch( error => {
          this.error = error
        } )
    }
  }
}
</script>

<style scoped>
</style>
