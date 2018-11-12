<template>
  <nav class="navbar gonotes-navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <router-link to="/">
      <h3 class="is-size-3 header-item" style="color: #375EAB;"> Заметки о Go </h3>
    </router-link>

    <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div id="navbarBasicExample" class="navbar-menu">
    <div class="navbar-start">
      <div class="navbar-item">
      <h4 class="is-size-4 header-item"> Разделы: &nbsp; </h4>
      <div class="select">
        <select
          v-model="category"
          @change="$emit('categoryChanged', category)"
        >
          <option>Все разделы</option>
          <option
            v-for="cat in categories" :key="cat.name"
          >
            {{ cat.Name }}
          </option>
        </select>
      </div>
      </div>
      <div class="navbar-item">
      <div class="buttons">
        <a
          class="button"
          @click="$emit('addNote')"
        >
          Добавить заметку
        </a>
        <a
          class="button"
          @click="$emit('addCategory')"
        >
          Добавить раздел
        </a>
        <a
          class="button"
          @click="$emit('addTag')"
        >
          Добавить маркер
        </a>
      </div>
      </div>
    </div>
    <div class="navbar-end">
      <div class="navbar-item">
        <div class="buttons">
          <a class="button" disabled>
            Зарегистроваться
          </a>
          <a class="button" disabled>
            Войти
          </a>
        </div>
      </div>
    </div>
  </div>
</nav>
</template>

<script>
export default {
  name: 'Navbar',
  components: {
  },
  props: {
    categories: Array
  },
  data: function () {
    return {
      category: 'Все разделы'
    }
  }
}
</script>

<style scoped>
.gonotes-navbar {
  background: #E0EBF5;
}
</style>
