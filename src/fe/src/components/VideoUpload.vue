<template>
    <div class="container mx-auto py-8">
      <div class="mx-auto">
        <div class="bg-white shadow rounded p-6">
          <h1 class="text-2xl font-bold mb-6  text-black">Video Upload</h1>
          <div class="flex items-center justify-between mb-6">
            <label for="fileInput" class="px-4 py-2 bg-blue-500 w-full flex text-white rounded cursor-pointer hover:bg-blue-600">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
              Select Video
            </label>
            <input id="fileInput" type="file" @change="handleFileSelect" accept="video/*" class="hidden">
            <span v-if="selectedFile" class="ml-4 text-gray-500 text-center">{{ selectedFile.name }}</span>
            <span v-else class="ml-4 text-gray-500 text-center" >no file selected</span>
          </div>
          <div v-if="selectedFile" class="mb-6 flex flex-row">
            <div class="text-lg font-bold mb-2 flex justify-between text-black">
                <h2>Preview</h2>
                <button @click="dismissFile" class="hover:text-red-400">X</button>
            </div>
          <div class="flex justify-center">
            <video :src="selectedVideoUrl" controls class=""></video>
          </div>
          </div>
          <button @click="uploadVideo" :disabled="!selectedFile" class="w-full bg-blue-500 text-white py-2 rounded disabled:bg-gray-400 disabled:cursor-not-allowed">
            Upload
          </button>
        </div>
        <div v-if="uploadedVideoUrl" class="bg-white shadow rounded mt-6">
          <h2 class="text-xl font-bold p-4">Uploaded Video</h2>
          <video :src="uploadedVideoUrl" controls class="w-full"></video>
        </div>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { ref, defineComponent } from 'vue';
  
  export default defineComponent({
    name: 'VideoUpload',
    data() {
      return {
        selectedFile: null as File | null,
        selectedVideoUrl: null as string | null,
        uploadedVideoUrl: null as string | null,
      };
    },
    methods: {
      handleFileSelect(event: Event) {
        const inputElement = event.target as HTMLInputElement;
        if (inputElement.files && inputElement.files.length > 0) {
          this.selectedFile = inputElement.files[0];
          this.selectedVideoUrl = URL.createObjectURL(this.selectedFile);
        }
      },
      dismissFile() {
        this.selectedFile = null
        this.selectedVideoUrl = null
      },
      uploadVideo() {
        if (this.selectedFile) {
          // TODO: Upload video to storage
          this.uploadedVideoUrl = URL.createObjectURL(this.selectedFile);
        }
      },
    },
  });
  </script>
  
  <style>

  </style>
  