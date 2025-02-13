<template>
  <div class="container mt-5">
    <h1 class="mb-4">Notes</h1>
    <div class="add-note-container">
      <button class="btn btn-primary mb-3" @click="showCreateForm = true">Add Note</button>
    </div>
    <div v-if="showCreateForm" class="card mb-3">
      <div class="card-body">
        <h2 class="card-title">Create Note</h2>
        <form @submit.prevent="createNote">
          <div class="mb-3">
            <input v-model="newNote.title" class="form-control" placeholder="Title" required />
          </div>
          <div class="mb-3">
            <textarea v-model="newNote.content" class="form-control" placeholder="Content" required></textarea>
          </div>
          <div class="mb-3">
            <input v-model="newNote.tag" class="form-control" placeholder="Tag" required />
          </div>
          <div class="mb-3">
            <input v-model="newNote.category" class="form-control" placeholder="Category" required />
          </div>
          <button type="submit" class="btn btn-success">Create</button>
          <button type="button" class="btn btn-secondary" @click="showCreateForm = false">Cancel</button>
        </form>
      </div>
    </div>
    <div class="notes-grid">
      <div class="note-card" v-for="note in notes" :key="note.id">
        <div class="card-body">
          <h3 class="card-title">{{ note.title }}</h3>
          <p class="card-text">{{ note.content }}</p>
          <p class="card-text"><small class="text-muted">{{ note.tag }}</small></p>
          <p class="card-text"><small class="text-muted">{{ note.category }}</small></p>
          <button class="btn btn-warning btn-sm" @click="editNote(note)">Edit</button>
          <button class="btn btn-danger btn-sm" @click="deleteNote(note.id)">Delete</button>
        </div>
      </div>
    </div>
    <div v-if="showEditForm" class="card mt-3">
      <div class="card-body">
        <h2 class="card-title">Edit Note</h2>
        <form @submit.prevent="updateNote">
          <div class="mb-3">
            <input v-model="currentNote.title" class="form-control" placeholder="Title" required />
          </div>
          <div class="mb-3">
            <textarea v-model="currentNote.content" class="form-control" placeholder="Content" required></textarea>
          </div>
          <div class="mb-3">
            <input v-model="currentNote.tag" class="form-control" placeholder="Tag" required />
          </div>
          <div class="mb-3">
            <input v-model="currentNote.category" class="form-control" placeholder="Category" required />
          </div>
          <button type="submit" class="btn btn-success">Update</button>
          <button type="button" class="btn btn-secondary" @click="showEditForm = false">Cancel</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      notes: [],
      newNote: {
        title: '',
        content: '',
        tag: '',
        category: ''
      },
      currentNote: null,
      showCreateForm: false,
      showEditForm: false
    };
  },
  created() {
    this.fetchNotes();
  },
  methods: {
    async fetchNotes() {
      try {
        const response = await axios.get('http://localhost:8000/notes');
        this.notes = response.data;
      } catch (error) {
        console.error(error);
      }
    },
    async createNote() {
      try {
        const response = await axios.post('http://localhost:8000/notes', this.newNote);
        this.notes.push(response.data);
        this.newNote = { title: '', content: '', tag: '', category: '' };
        this.showCreateForm = false;
      } catch (error) {
        console.error(error);
      }
    },
    editNote(note) {
      this.currentNote = { ...note };
      this.showEditForm = true;
    },
    async updateNote() {
      try {
        const response = await axios.put(`http://localhost:8000/notes/${this.currentNote.id}`, this.currentNote);
        const index = this.notes.findIndex(note => note.id === this.currentNote.id);
        this.notes.splice(index, 1, response.data);
        this.currentNote = null;
        this.showEditForm = false;
      } catch (error) {
        console.error(error);
      }
    },
    async deleteNote(id) {
      try {
        await axios.delete(`http://localhost:8000/notes/${id}`);
        this.notes = this.notes.filter(note => note.id !== id);
      } catch (error) {
        console.error(error);
      }
    }
  }
};
</script>

<style scoped>
.container {
  max-width: 1200px;
  margin: auto;
}

.add-note-container {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.notes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.note-card {
  background-color: #fff;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 15px;
  transition: box-shadow 0.3s ease;
}

.note-card:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.card-body {
  display: flex;
  flex-direction: column;
}

.card-title {
  font-size: 1.25rem;
  margin-bottom: 10px;
}

.card-text {
  flex-grow: 1;
}

.btn {
  margin-top: 10px;
}
</style>