<template>
  <div class="modal" :class="{'is-active': isActive}">
    <div class="modal-background"></div>
    <div class="modal-content">
      <div class="field">
        <div class="control">
          <input class="input is-primary" type="text" placeholder="название" v-model="name">
        </div>
      </div>
      <div class="field">
        <div class="control">
          <input class="input is-info" type="text" placeholder="описание" v-model="description">
        </div>
      </div>
      <div class="field">
        <a class="button" @click="addTag">Добавить</a>
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
  name: 'AddTag',
  props: {
    isActive: Boolean
  },
  data: function () {
    return {
      name: '',
      description: ''
    }
  },
  methods: {
    addTag: function () {
      const postURL = 'http://' + window.location.hostname + '/api/v1/addTag'
      const payload = {
        Name: this.name,
        Description: this.description
      }
      axios.post(postURL, payload, null)
        .then(response => {
          this.$emit('close')
        }).catch(error => {
          console.log(error)
        }) // axios.post
    }
  }
}
</script>

<style scoped>

</style>
