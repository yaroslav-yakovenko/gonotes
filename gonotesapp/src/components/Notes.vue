<template>
  <div class="column notes">
      <div class="container">
        <div
          v-for="cat in catsFiltered" :key="cat.id"
        >
        <h4 class="is-size-4 note-title">
        {{cat.Name}} - <span style="color: black;">{{cat.Description}}</span>
        </h4>
        <div
          v-for="note in notes" :key="note.name"
        >
          <h5
            v-if="note.CategoryID === cat.id"
            class="is-size-5 note-title"
          >
            <a @click="titleClicked(note)">{{ note.Title }}</a>
            <a class="button is-white" v-if="loggedIn" @click="editClicked(note)">Редактировать</a>
          </h5>
          <div
            class="note-body"
            v-if="note.CategoryID === cat.id"
            v-show="selectedTitle === note.Title"
          >
            <div class="control">
              <textarea class="textarea" readonly placeholder="Normal textarea" v-model="note.Body" rows="20"></textarea>
            </div>
          </div>
        </div>
      </div>
      </div>
  </div>
</template>

<script>
export default {
  name: 'Notes',
  props: {
    loggedIn: Boolean,
    notes: Array,
    categoryDescription: String,
    categories: Array
  },
  data: function () {
    return {
      selectedTitle: ''
    }
  },
  methods: {
    titleClicked: function (note) {
      if (this.selectedTitle === note.Title) {
        this.selectedTitle = ''
        return
      }
      this.selectedTitle = note.Title
    },
    editClicked: function (note) {
      this.$emit('editNote', note)
    }
  },
  computed: {
    catsFiltered: function () {
      let cats = []
      for (let i = 0; i < this.categories.length; i++) {
        let found = false
        for (let j = 0; j < this.notes.length; j++) {
          if (this.notes[j].CategoryID === this.categories[i].id) {
            found = true
          }
        }
        if (found) {
          cats.push(this.categories[i])
        }
      }
      return cats
    }
  }
}
</script>

<style scoped>
.note-title {
  color: #375EAB;
  padding-top: 10px;
}
.note-body {
  color: black;
}
</style>
