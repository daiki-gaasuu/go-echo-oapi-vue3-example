<script lang="ts" setup>
import { defineComponent } from 'vue';

// Logo
import logo from '../assets/logo.svg';
</script>

<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img :src="logo" class="my-3" contain height="200" />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          名前を入力してHelloを押すと挨拶してくれます
        </h1>
        <v-responsive class="mx-auto" max-width="300">
          <v-text-field label="名前を入力" variant="outlined"></v-text-field>
        </v-responsive>
        <v-btn color="blue">Hello</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>
