<template>
  <div class="modal" :class="{'is-active': isActive}">
    <div class="modal-background"></div>
    <div class="modal-content">

      <div class="field">
      <h4 class="is-size-4 header-item"> Раздел: &nbsp; </h4>
        <div class="select">
          <select
            v-model="category"
          >
            <option></option>
            <option
              v-for="cat in categories" :key="cat.Name"
            >
              {{ cat.Name }}
            </option>
          </select>
        </div>
      </div>

      <div class="field">
        <div class="control">
          <input class="input is-primary" type="text" placeholder="заголовок" v-model="title">
        </div>
      </div>
      <div class="field">
        <div class="control">
          <textarea class="textarea" placeholder="текст заметки" v-model="body" rows="15"></textarea>
        </div>
      </div>

      <div class="field">
      <h4 class="is-size-4 header-item"> Добавить маркер: &nbsp; </h4>
        <div class="select">
          <select
            v-model="tag"
          >
            <option></option>
            <option
              v-for="tag in tags" :key="tag.Name"
            >
              {{ tag.Name }}
            </option>
          </select>
        </div>
        <a class="button" @click="addTag">Добавить маркер</a>
      </div>

      <div class="field">
        <div class="control"
          v-for="tag in selectedTags" :key="tag.name"
        >
          <div class="tags has-addons">
            <a class="tag">{{tag.Name}}</a>
            <span class="tag is-dark">{{tag.Description}}</span>
          </div>
        </div>
      </div>

      <div class="field is-grouped">
        <p class="control">
          <a class="button" @click="clearAll">Очистить всё</a>
        </p>
        <p class="control">
          <a class="button is-primary" @click="addNote" v-if="addNoteFlag">Добавить заметку</a>
          <a class="button is-primary" @click="updateNote" v-if="!addNoteFlag">Обновить заметку</a>
        </p>
      </div>

    </div>
    <button
      class="modal-close is-large"
      aria-label="close"
      @click="$emit('close')"
    >
    </button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'EditNote',
  props: {
    categories: Array,
    tags: Array,
    isActive: Boolean,
    note: Object
  },
  data: function () {
    return {
      addNoteFlag: false,
      category: '',
      tag: '',
      title: '',
      body: '',
      selectedTags: []
    }
  },
  methods: {
    addTag: function () {
      if (!this.selectedTags) {
        this.selectedTags = []
      }
      for (let i = 0; i < this.tags.length; i++) {
        if (this.tags[i].Name === this.tag) {
          this.selectedTags.push(this.tags[i])
        }
      }
    },
    clearAll: function () {
      this.category = ''
      this.title = ''
      this.body = ''
      this.selectedTags = []
    },
    addNote: function () {
      const postURL = 'http://' + window.location.hostname + '/api/v1/addNote'
      const payload = {
        Title: this.title,
        Body: this.body,
        Tags: this.selectedTags
      }
      for (let i = 0; i < this.categories.length; i++) {
        if (this.categories[i].Name === this.category) {
          payload.CategoryID = this.categories[i].id
        }
      }
      axios.post(postURL, payload, null)
        .then(response => {
          this.$emit('close')
        }).catch(error => {
          console.log(error)
        }) // axios.post
    },
    updateNote: function () {
      const postURL = 'http://' + window.location.hostname + '/api/v1/updateNote'
      const payload = {
        ID: this.note.id,
        Title: this.title,
        Body: this.body,
        Tags: this.selectedTags
      }
      for (let i = 0; i < this.categories.length; i++) {
        if (this.categories[i].Name === this.category) {
          payload.CategoryID = this.categories[i].id
        }
      }
      axios.post(postURL, payload, null)
        .then(response => {
          this.$emit('close')
        }).catch(error => {
          console.log(error)
        }) // axios.post
    }
  },
  created () {
    this.addNoteFlag = true
    if (this.note.Title) {
      this.addNoteFlag = false
      this.title = this.note.Title
      this.body = this.note.Body
      this.selectedTags = this.note.Tags
      for (let i = 0; i < this.categories.length; i++) {
        if (this.note.CategoryID === this.categories[i].id) {
          this.category = this.categories[i].Name
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
