<template>
  <div>
    <ol>
      <draggable v-model="items"
               draggable=".item"
               :move="onMove">
        <li v-for="item in items"
             v-bind:key="item.id"
             class="item">
          <input type="checkbox"
                 v-model="item.complete"
                 v-on:change="updateTodo( item )"/>
          <span v-if="!item.editing"
                class="itemTitle">
            <span v-if="!item.complete">{{item.title}}</span>
            <strike v-if="item.complete">{{item.title}}</strike>
          </span>
          <button v-if="!item.editing"
                  v-on:click="startEditTodo( item )">
            Edit
          </button>
          <button v-if="!item.editing"
                  v-on:click="deleteTodo( item )">
            Delete
          </button>
          <div v-if="!!item.editing"
               class="editTools">
            <input type="text"
                   v-model="item.editedTitle"
                   class="editedTitle"/>
            <button v-on:click="saveEditTodo( item )">Save</button>
            <button v-on:click="cancelEditTodo( item )">Cancel</button>
          </div>
        </li>
      </draggable>
    </ol>
    <div slot="footer">
      <label>
        Add a new item:
        <input type="text" v-model="newItemTitle"/>
      </label>
      <button v-on:click="createNewTodo()">Create</button>
    </div>
  </div>
</template>

<script>
import draggable from 'vuedraggable'
import axios from 'axios'
import _ from 'lodash'

function putTodo(todo) {
  return axios
    .put(`/api/v1/todos/${todo.id}`, todo)
}

export default {
  name: 'TodoList',
  data: () => ( {
    items: [],
    error: null,
    newItemTitle: ''
  } ),
  components: {
    draggable
  },
  mounted() {
    axios
      .get('/api/v1/todos')
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
      putTodo( todo )
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
      putTodo( todo )
        .catch( error => {
          this.error = error
        } )
    },
    cancelEditTodo: ( todo ) => {
      todo.editing = false
    },
    createNewTodo() {
      const todo = {
        id: Math.max( ...this.items.map( x => x.id ), -1 ) + 1,
        title: this.newItemTitle,
        position: Math.max( ...this.items.map( x => x.position ), -1 ) + 1,
        complete: false
      }

      this.items.push( todo )
      this.newItemTitle = ''

      axios
        .post( '/api/v1/todos', todo )
        .catch ( error => {
          this.error = error
        } )
    },
    deleteTodo( todo ) {
      this.items = this.items.filter( x => x.id !== todo.id )
      axios
        .delete(`/api/v1/todos/${todo.id}`)
        .catch( error => {
          this.error = error
        } )
    },
    onMove(evt) {
      const index = evt.draggedContext.index
      const futureIndex = evt.draggedContext.futureIndex
      const item = this.items[index]
      const futureItem = this.items[futureIndex]
      const futurePosition = futureItem.position
      futureItem.position = item.position
      item.position = futurePosition

      putTodo( item )
        .catch( error => {
          this.error = error
        } )
      putTodo( futureItem )
        .catch( error => {
          this.error = error
        } )
    }
  }
}
</script>

<style scoped>
  .itemTitle {
    display: inline-block;
    width: 150px;
  }
  .editedTitle {
    width: 145px;
  }
  .editTools {
    display: inline;
  }
</style>
