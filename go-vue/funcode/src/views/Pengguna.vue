<template>
<v-app>
<v-layout row wrap>
  <v-flex xs12 sm12>
      <v-card elevation="1">
        <v-card-title class="blue-grey lighten-5 font16">
          Data Pengguna
        </v-card-title>
          <v-list two-line subheader>
          <v-list-tile
            v-for="(user, index) in fetchUsers"
            :key="index"
            avatar
            :to="{name: 'detail-pengguna', params: { id: user.id }}"
          >

            <v-list-tile-avatar>
              <v-icon class="blue white--text">people</v-icon>
            </v-list-tile-avatar>

            <v-list-tile-content>
              <v-list-tile-title>{{ user.name }} - {{user.company.name }}</v-list-tile-title>
              <v-list-tile-sub-title>{{ user.email }} - {{user.phone}}</v-list-tile-sub-title>
            </v-list-tile-content>

          </v-list-tile>
        </v-list>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn flat color="green">Add Data</v-btn>
        </v-card-actions>
      </v-card>
    </v-flex>

    </v-layout>
    </v-app>
</template>

<script>
import {mapGetters} from 'vuex'

export default {
    name: "Pengguna",
    data(){
        return {
            // dataUsers: this.fetchUsers
        }
    },
    computed: {
    ...mapGetters([
            'fetchUsers'
        ])
    }

}
</script>


