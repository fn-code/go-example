<template>
<v-layout row wrap>
  <InfoComponent/>
  <v-flex xs12 sm12>
      <v-card elevation="1">
        <v-card-title class="blue-grey lighten-5 font16">
          Data Perusahaan IT
        </v-card-title>
          <v-list two-line subheader>
          <v-list-tile
            v-for="(link, index) in links"
            :key="index"
            avatar
            hover
          >
            <v-list-tile-avatar>
              <v-icon :class="[link.iconClass]">{{ link.icon }}</v-icon>
            </v-list-tile-avatar>

            <v-list-tile-content>
              <v-list-tile-title>{{ link.name }}</v-list-tile-title>
              <v-list-tile-sub-title>{{ link.url }}</v-list-tile-sub-title>
            </v-list-tile-content>

            <v-list-tile-action>
              <v-btn icon ripple class="grey lighten-1" @click="removeLinks(index)">
                <v-icon color="grey lighten-4">close</v-icon>
              </v-btn>
            </v-list-tile-action>
          </v-list-tile>
        </v-list>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn flat color="green" @click="dialogState">Add Data</v-btn>
        </v-card-actions>
      </v-card>
    </v-flex>

    <AddDialog/>
    </v-layout>
</template>

<script>
import { mapState, mapMutations, mapActions} from 'vuex'
import InfoComponent from "./../components/about/Info"
import AddDialog from "./../components/about/AddDialog"
;
export default {
  name: "About",
  components: {
    InfoComponent,
    AddDialog
  },
  data (){
    return {

    }
  },
  computed: {
    ...mapState([
      'links',
    ])
  },
  methods: {
    ...mapMutations([
      'CHANGE_DIALOG_STATE'
    ]),
    ...mapActions([
      'removeLink'
    ]),
    dialogState: function() {
      this.CHANGE_DIALOG_STATE();
    },
    removeLinks: function(data) {
      this.removeLink(data);
    }
  }
}
</script>

