<template>
  <div class="column tags is-one-fifth">
    <div class="level-item has-text-centered tags-title">
      <h3 class="is-size-3 note-title">Маркеры</h3>
    </div>
    <div class="level-item has-text-centered">
      <div class="field is-grouped is-grouped-multiline">
        <div class="control"
          v-for="tag in tags" :key="tag.Name"
        >
          <div class="tags has-addons">
            <a
             class="tag"
             @click="tagClicked(tag)"
           >
             {{tag.Name}}
           </a>
            <span class="tag is-dark">{{tag.Description}}</span>
          </div>
        </div>
      </div>
    </div>
    <br />
    <div class="level-item has-text-centered tags-title" v-if="currentTag.Name">
        <div class="tags has-addons">
          <span class="tag is-danger">{{currentTag.Name}}</span>
          <a class="tag is-delete" @click="clearTag"></a>
      </div>
    </div>
    <br />

    <Footer>
    </Footer>

  </div>
</template>

<script>
import Footer from './Footer'

export default {
  name: 'Tags',
  components: {
    // eslint-disable-next-line
    Footer
  },
  props: {
    tags: Array
  },
  data: function () {
    return {
      currentTag: {}
    }
  },
  methods: {
    tagClicked: function (tag) {
      this.currentTag = tag
      this.$emit('tagClicked', tag.Name)
    },
    clearTag: function () {
      this.currentTag = {}
      this.$emit('clearTag')
    }
  }
}
</script>

<style scoped>
.tags-title {
  color: #375EAB;
  padding-top: 10px;
}
</style>
