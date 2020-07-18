<template>

    <v-dialog v-model="Dialog" persistent max-width="600px">
      <v-card>
        <v-card-title class="blue-grey lighten-5 font16">
          Data Perusahaan IT
        </v-card-title>

        <v-card-text>
          <v-container  wrap grid-list-md>
            <v-layout wrap>
              <v-flex xs12 sm6 md6>
                <v-text-field v-model="names" label="First Company Name *" required></v-text-field>
              </v-flex>
              <v-flex xs12 sm6 md6>
                <v-text-field v-model="urls" label="Company Url" hint="example www.google.com"></v-text-field>
              </v-flex>
              <v-flex xs12 sm6 md6>
                <v-text-field v-model="icons" label="Icon Name*" required></v-text-field>
              </v-flex>
              <v-flex xs12 sm6 md6>
                <v-select v-model="iconStyle"
                  :items="['blue white--text', 'green white--text', 'red white--text', 'yellow white--text']"
                  label="Age*"
                  required
                ></v-select>
              </v-flex>

            </v-layout>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>



        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" flat @click="dialogState">Close</v-btn>
          <v-btn color="blue darken-1" flat @click="addDataPerusahaan">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
</template>

<script>

import { mapState, mapMutations } from "vuex";

export default {
    name: "AddDialog",
    data () {
        return {
          names: '',
          urls: '',
          icons: '',
          iconStyle: '',
          dataPer: {}
        }
    },
    computed: {
      ...mapState([
        'Dialog'
      ])
    },
    methods: {
      ...mapMutations([
      'CHANGE_DIALOG_STATE',
      'ADD_PERUSAHAAN'
      ]),
        dialogState: function() {
        this.CHANGE_DIALOG_STATE();
      },
      addDataPerusahaan: function(){
        this.dataPer = {
          name: this.names,
          url: this.urls,
          icon: this.icons,
          iconClass: this.iconStyle
        }
        this.ADD_PERUSAHAAN(this.dataPer);
        this.CHANGE_DIALOG_STATE();

        this.names='';
        this.urls='';
        this.icons='';
        this.iconStyle = '';
      },
    }
}
</script>

