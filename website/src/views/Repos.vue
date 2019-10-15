<template>
  <v-container>
    <v-row :justify="alignment">
        <v-col cols="12">
            <v-card>
                <v-card-title> This is a card </v-card-title>
            </v-card>
        </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
    props:[],
    components: {
    },
    data: () => ({
        alignment:"center",
        username:null,
    }),
    
};
</script>
