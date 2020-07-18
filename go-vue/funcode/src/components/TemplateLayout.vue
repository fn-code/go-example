<template>
<v-app>
  <v-layout wrap style="height: 200px;">
    <v-navigation-drawer v-model="drawer" fixed clipped app>
      <v-list class="pa-1">
        <v-list-tile avatar>
          <v-list-tile-avatar>
            <img src="https://secure.gravatar.com/avatar/0bd6cc64aafc39299b5b46c31956f4af.jpg?s=80&r=g&d=mm">
          </v-list-tile-avatar>

          <v-list-tile-content>
            <v-list-tile-title> Ludin Nento</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>

      <v-list class="pt-0" dense>
        <v-divider></v-divider>

        <v-list-tile
          v-for="item in items"
          :key="item.title"
          :to="{name: item.action}"
        >
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>

          <v-list-tile-content>
            <v-list-tile-title>{{ item.title }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>



    <v-toolbar dark color="primary" dense height="54" app clipped-left>
    <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
    <v-toolbar-title class="white--text"> {{APPNAME}}</v-toolbar-title>
    <v-spacer></v-spacer>

    <v-btn icon>
      <v-icon>search</v-icon>
    </v-btn>

    <v-btn icon>
      <v-icon>apps</v-icon>
    </v-btn>

    <v-btn icon>
      <v-icon>refresh</v-icon>
    </v-btn>

    <v-btn icon>
      <v-icon>more_vert</v-icon>
    </v-btn>
  </v-toolbar>


    <v-content class="grey lighten-4">
            <v-container grid-list-md text-xs-center>
                    <router-view/>
            </v-container>
    </v-content>


  </v-layout>

</v-app>
</template>

<script>

import {  mapState } from "vuex";

export default {
    name: "TemplateLayout",
    data: function() {
        return {
            drawer: null,
            items: [
                { title: 'Home', icon: 'home', action: "home" },
                { title: 'Pengguna', icon: 'people', action: "pengguna" },
                { title: 'About', icon: 'dashboard', action: "about" }
            ]
        }
    },
    computed: {
      ...mapState([
        'APPNAME'
      ])
    },
     methods: {

    }

}
</script>

<style>
</style>