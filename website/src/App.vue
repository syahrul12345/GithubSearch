<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>GITHUB SEARCH</span>
      </v-toolbar-title>
    </v-app-bar>

    <v-content>
      <Router-View/>
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App',
  components: {
    
  },
  data: () => ({
    //
  }),
};
</script>
