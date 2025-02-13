<template>
  <div id="app">
    <NoteList />
  </div>
</template>

<script>
import NoteList from './components/NoteList.vue';

export default {
  name: 'App',
  components: {
    NoteList
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
