<template>
  <v-container>
      <v-row>
        <v-col cols="12">
          <v-img
            :src="require('../assets/Github.svg')"
            class="my-3"
            contain
            height="200"
          ></v-img>
        </v-col>
      </v-row>
      <v-row :justify="alignment">
            <v-col cols="12">
                <v-text-field
                filled
                v-model="username"
                background-color="white"
                label="Search for a Username to get all repos"
                append-icon="fas fa-search"
                class = "px-2"
                @click:append="search"
                ></v-text-field>
            </v-col>
        </v-row>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    alignment:"center",
    username:null,
  }),
  methods:{
    search() {
      this.$router.push({path:`/repos/${this.username}`})
    }
  }
};
</script>
