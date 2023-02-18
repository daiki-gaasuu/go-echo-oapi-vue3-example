<script lang="ts" setup>
import { ref } from 'vue';
import GreetingApi from '@/services/greeting-api';
// Logo
import logo from '../assets/logo.svg';

const inputName = ref('');
const resName = ref('');
const isError = ref(false);

const submit = async () => {
  try {
    const res = await GreetingApi.getGreeting(inputName.value);
    if (isError) isError.value = false;
    resName.value = res.data;
  } catch (error) {
    console.log(error);
    isError.value = true;
  }
};
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
          <v-text-field
            v-model="inputName"
            label="名前を入力"
            variant="outlined"
          ></v-text-field>
        </v-responsive>
        <v-btn @click="submit" color="blue">Hello</v-btn>
        <p v-if="isError" class="text-red text-center font-weight-bold">
          名前がありません
        </p>
        <p v-if="resName" class="text-center font-weight-bold">
          Hello {{ resName }}!!
        </p>
      </v-col>
    </v-row>
  </v-container>
</template>
