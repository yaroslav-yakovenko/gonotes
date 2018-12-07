<template>
  <div>

    <Navbar
      :categories="categories"
      @categoryChanged="categoryChanged"
      @addCategory="showAddCategory = true"
      @addTag="showAddTag = true"
      @addNote="addNote"
    >
    </Navbar>

    <div class="columns" style="margin-left: 5px; margin-top: 5px;">

      <Tags
        :tags="tags"
        @tagClicked="tagClicked"
        @clearTag="clearTag"
      >
      </Tags>

      <Notes
        :notes="filteredNotes"
        :categories="categories"
        :categoryDescription="categoryDescription"
        @editNote="editNote"
      >
      </Notes>

    </div>

    <AddCategory
      :isActive="showAddCategory"
      @close="closeCategory"
    >
    </AddCategory>
    <AddTag
      :isActive="showAddTag"
      @close="closeTag"
    >
    </AddTag>
    <EditNote
      v-if="showNoteEditor"
      :isActive="showNoteEditor"
      :categories="categories"
      :tags="tags"
      :note="note"
      @close="closeNote"
    >
    </EditNote>

    <Footer>
    </Footer>

  </div>
</template>

<script>
import Navbar from './Navbar'
import Tags from './Tags'
import Notes from './Notes'
import EditNote from './EditNote'
import AddCategory from './AddCategory'
import AddTag from './AddTag'

import axios from 'axios'

export default {
  name: 'Home',
  components: {
    Navbar,
    Tags,
    Notes,
    EditNote,
    AddCategory,
    AddTag
  },
  data: function () {
    return {
      showAddTag: false,
      showAddCategory: false,
      showNoteEditor: false,
      selectedCategory: 'Все разделы',
      selectedTag: '',
      categoryDescription: 'All Notes',
      categories: [],
      tags: [],
      notes: [],
      filteredNotes: [],
      note: {}
    }
  },
  methods: {
    fetchCategories: function () {
      const fetchURL = 'http://' + window.location.hostname + '/api/v1/getCategories'
      axios.get(fetchURL)
        .then(response => {
          this.categories = response.data
        }).catch(error => {
          console.log(error)
        }) // axios.get
    },
    fetchTags: function () {
      const fetchURL = 'http://' + window.location.hostname + '/api/v1/getTags'
      axios.get(fetchURL)
        .then(response => {
          this.tags = response.data
        }).catch(error => {
          console.log(error)
        }) // axios.get
    },
    fetchNotes: function () {
      const fetchURL = 'http://' + window.location.hostname + '/api/v1/getNotes'
      axios.get(fetchURL)
        .then(response => {
          this.notes = response.data
          if (this.selectedCategory === 'Все разделы' && this.selectedTag === '') {
            this.filteredNotes = this.notes
          }
        }).catch(error => {
          console.log(error)
        }) // axios.get
    },
    categoryChanged: function (category) {
      this.selectedCategory = category
    },
    tagClicked: function (tag) {
      this.selectedTag = tag
    },
    clearTag: function () {
      this.selectedTag = ''
    },
    getDescription: function () {
      this.categoryDescription = 'Все разделы'
      for (let i = 0; i < this.categories.length; i++) {
        if (this.categories[i].Name === this.selectedCategory) {
          this.categoryDescription = this.categories[i].Description
        }
      }
    },
    closeCategory: function () {
      this.showAddCategory = false
      this.fetchCategories()
    },
    closeTag: function () {
      this.showAddTag = false
      this.fetchTags()
    },
    closeNote: function () {
      this.showNoteEditor = false
      this.fetchNotes()
    },
    filterNotesByCategory: function () {
      this.filteredNotes = []
      if (this.selectedCategory === 'Все разделы') {
        this.filteredNotes = this.notes
      }
      let catID = ''
      for (let i = 0; i < this.categories.length; i++) {
        if (this.selectedCategory === this.categories[i].Name) {
          catID = this.categories[i].id
        }
      }
      if (catID !== '') {
        for (let i = 0; i < this.notes.length; i++) {
          if (this.notes[i].CategoryID === catID) {
            this.filteredNotes.push(this.notes[i])
          }
        }
      }
    },
    filterNotesByTag: function () {
      this.filteredNotes = []
      if (this.selectedTag === '') {
        this.filteredNotes = this.notes
        this.filterNotesByCategory()
        return
      }
      let tagID = ''
      for (let i = 0; i < this.tags.length; i++) {
        if (this.selectedTag === this.tags[i].Name) {
          tagID = this.tags[i].id
        }
      }
      if (tagID !== '') {
        for (let i = 0; i < this.notes.length; i++) {
          if (!this.notes[i].Tags) {
            continue
          }
          for (let j = 0; j < this.notes[i].Tags.length; j++) {
            if (this.notes[i].Tags[j].id === tagID) {
              this.filteredNotes.push(this.notes[i])
            }
          }
        }
      }
    },
    addNote: function () {
      this.note = {}
      this.showNoteEditor = true
    },
    editNote: function (note) {
      this.note = note
      this.showNoteEditor = true
    }
  },
  computed: {
  },
  watch: {
    selectedCategory: function () {
      this.getDescription()
      this.filterNotesByCategory()
    },
    selectedTag: function () {
      this.filterNotesByTag()
    }
  },
  created () {
    this.fetchCategories()
    this.fetchTags()
    this.fetchNotes()
  }
}
</script>

<style scoped>
</style>
