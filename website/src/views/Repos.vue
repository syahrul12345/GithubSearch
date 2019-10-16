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
            <v-card
            :hover ="true"
            :elevated="true"
            @click="goToReadme($route.params.username,repo.Name)">
                <v-card-title> {{repo.Name}} </v-card-title>
            </v-card>
        </v-col>
    </v-row>
        <v-dialog 
        max-width="500" 
        v-model="readmeError">
            <v-card>
                <v-card-title> Oops something went wrong </v-card-title>
                <v-card-text> This repository has no readme </v-card-text>
            </v-card>
        </v-dialog> 
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
        readmeError:false,
        repos: null,
    }),
    async beforeRouteEnter(to,from,next){
        const params = to.params.username
        await axios.post("/api/v1/getUser",{
                    username:params
                }).then((response) => {
                    //Valid response with all Repos, set it to the data object
                    next(vm => {
                        vm.repos = response.data.results
                    })
                }).catch((error) => {
                    next('/error')
                })
    },
    async beforeRouteLeave(to,from,next){
        //Only route if the 'to' field is to the readme route, otherwise, ignore
        if(to.name == "readme") {
            const username = to.params.username
            const repository = to.params.repository
            await axios.post("/api/v1/getRepo",{
                username,
                repository
            }).then((response) => {
                next()
            }).catch((error) => {
                console.log(error)
                this.readmeError = true
            })
        }else{
            //If going back
            next()
        }
    },
    methods: {
        goToReadme(username,repository) {
            this.$router.push({path:`/readme/${username}/${repository}`})
        },
    }
    
};
</script>