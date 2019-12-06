<template>
  <ul>
    <li v-for="item in items" v-bind:key="item.id">
      <input type="checkbox"
             v-model="item.complete"
             v-on:change="updateTodo( item )"/>
      <span v-if="!item.editing">
        <span v-if="!item.complete">{{item.title}}</span>
        <strike v-if="item.complete">{{item.title}}</strike>
      </span>
      <button v-if="!item.editing"
              v-on:click="startEditTodo( item )">
        Edit
      </button>
      <div v-if="!!item.editing">
        <input type="text" v-model="item.editedTitle"/>
        <button v-on:click="saveEditTodo( item )">Save</button>
        <button v-on:click="cancelEditTodo( item )">Cancel</button>
      </div>
    </li>
    <li>
      <label>
        Add a new item:
        <input type="text" v-model="newItemTitle"/>
      </label>
      <button v-on:click="createNewTodo()">Create</button>
    </li>
    {{items}}
  </ul>
</template>

<script>
// import draggable from 'vuedraggable'
import axios from 'axios'
import _ from 'lodash'

export default {
  name: 'TodoList',
  data: () => ( {
    items: [],
    error: null,
    newItemTitle: ''
  } ),
  mounted() {
    axios
      .get('https://thawing-bayou-17829.herokuapp.com/todos')
      .then( response => {
        this.items = _.orderBy(response.data, x => x.position).map( x => ( {
            ...x,
            editing: false,
            editedTitle: ''
        } ) )
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
    },
    startEditTodo: ( todo ) => {
      todo.editing = true
      todo.editedTitle = todo.title
    },
    saveEditTodo: ( todo ) => {
      todo.editing = false
      todo.title = todo.editedTitle
      axios
        .put(`https://thawing-bayou-17829.herokuapp.com/todos/${todo.id}`, todo)
        .catch( error => {
          this.error = error
        } )
    },
    cancelEditTodo: ( todo ) => {
      todo.editing = false
    },
    createNewTodo() {
      const todo = {
        id: Math.max( ...this.items.map( x => x.id ) ) + 1,
        title: this.newItemTitle,
        position: Math.max( ...this.items.map( x => x.position ) ) + 1,
        complete: false
      }

      this.items.push( todo )
      this.newItemTitle = ''

      axios
        .post( 'https://thawing-bayou-17829.herokuapp.com/todos', todo )
        .catch ( error => {
          this.error = error
        } )
    }
  }
}
</script>

<style scoped>
  ul {
    list-style: none;
    padding: 0;
  }
  li {
    display: flex;
  }
</style>
