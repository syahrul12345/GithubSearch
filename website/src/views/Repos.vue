<template>
  <v-container>
    <v-row :justify="alignment">
        <v-col cols=12>
            <h1 style="text-align:center"> List of all repositories </h1>
        </v-col>
    </v-row>
    <v-row :justify="alignment">
        <v-col cols="4"
        v-for="repo in repos"
        v-bind:key="repo.Name"
        >
            <v-card>
                <v-card-title> {{repo.Name}} </v-card-title>
            </v-card>
        </v-col>
    </v-row>
  </v-container>
</template>

<script>
const axios = require('axios')
export default {
    props:[],
    components: {
    },
    data: () => ({
        alignment:"center",
        username:null,
        repos: null,
    }),

    async beforeRouteEnter(to,from,next){
        const params = to.params.username
        const post = await axios.post("http://127.0.0.1:9999/api/v1/getUser",{
                    username:params
                }).then((response) => {
                    //Valid response with all Repos, set it to the data object
                    next(vm => {
                        vm.repos = response.data.results
                    })
                }).catch((error) => {
                    console.log(error)
                    // next('/error')
                })
        
    }
    
};
</script>