<template>
  <nav class="navbar gonotes-navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <router-link to="/">
      <h3 class="is-size-3 header-item" style="color: #375EAB; margin-left: 5px;"> Заметки о Go </h3>
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
          :disabled="!loggedIn"
          class="button"
          @click="addNote"
        >
          Добавить заметку
        </a>
        <a
          :disabled="!loggedIn"
          class="button"
          @click="addCategory"
        >
          Добавить раздел
        </a>
        <a
          :disabled="!loggedIn"
          class="button"
          @click="addTag"
        >
          Добавить маркер
        </a>
      </div>
      </div>
    </div>
    <div class="navbar-end">
      <div class="navbar-item" v-if="!loggedIn">
        <input class="input" type="email" placeholder="почта" v-model="email" @keyup.enter="login">
      </div>
      <div class="navbar-item" v-if="!loggedIn">
        <input class="input" type="password" placeholder="пароль" v-model="password" @keyup.enter="login">
      </div>
      <div class="navbar-item">
        <div class="buttons">
          <a class="button" disabled>
            Зарегистроваться
          </a>
          <a
            class="button"
            @click="login"
            v-if="!loggedIn"
          >
            Войти
          </a>
          <a
            class="button"
            @click="logout"
            v-if="loggedIn"
          >
            Выйти
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
    loggedIn: Boolean,    
    categories: Array
  },
  data: function () {
    return {
      email: '',
      password: '',
      category: 'Все разделы'
    }
  },
  methods: {
    login: function () {
      this.$emit('login', this.email, this.password)
    },
    logout: function () {
      this.$emit('logout')
    },
    addNote: function () {
      if (this.loggedIn) {
        this.$emit('addNote')
      }
    },
    addCategory: function () {
      if (this.loggedIn) {
        this.$emit('addCategory')
      }
    },
    addTag: function () {
      if (this.loggedIn) {
        this.$emit('addTag')
      }
    },
  }
}
</script>

<style scoped>
.gonotes-navbar {
  background: #E0EBF5;
}
</style>
