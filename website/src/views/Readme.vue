<template>
    <v-container>
        
        <v-row 
        :justify="alignment">
            <v-col cols=12>
                <article class="markdown-body" v-html="text">
                </article>
            </v-col>
        </v-row>
        
    </v-container>

</template>
<script>
const axios = require('axios')
const showdown  = require('showdown')
export default {
    props: [],
    components: {
    },
    data() {
        return {
            alignment:"center",
            readmeError:false,
            text:null,
        }
    },
    async beforeRouteEnter(to,from,next) {
        await axios.post("/api/v1/getRepo",{
            username:to.params.username,
            repository:to.params.repository
        }).then((response) => {
            next(vm => {
                var converter = new showdown.Converter(),
                html = converter.makeHtml(response.data.readme)
                vm.text = `${html}`
            })
        }).catch((error) => {
            console.log(error)
            next(vm => {
                vm.text = "This repository has no readme"
            })
        })
    }
}

</script>
<style>
    .markdown-body {
		box-sizing: border-box;
		min-width: 200px;
		max-width: 980px;
		margin: 0 auto;
		padding: 45px;
	}

	@media (max-width: 767px) {
		.markdown-body {
			padding: 15px;
		}
	}
</style>