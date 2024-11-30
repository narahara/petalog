<template>
  <div class="home">
    <div>test</div>
    <button @click="ping">Ping</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { ApiClient } from '@/axios';
import { API_BASE_URL } from '@/constatnt';
import {petalog} from '@/dto/dto';


export default defineComponent({
  name: 'HomeView',
  components: {
  },
  setup() {
    const apiClient = new ApiClient(API_BASE_URL);
    return {
      apiClient,
    };
  },
  methods: {
    async ping() {
      try {
        var x = new petalog.service.console.PingRequest({
          message: 'ping22',
        });
        console.log(x);

        var res = await this.apiClient.get('/ping', {
          params: x,
        });
        var q = new petalog.service.console.PingResponse(res.data);
        

        console.log(res);
        console.log(q.message);
      } catch (e) {
        console.error(e);
      }
    },
  },
});
</script>
